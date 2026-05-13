package mypolymarketapi

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    RequestType = "GET"
	POST   RequestType = "POST"
	DELETE RequestType = "DELETE"
	PUT    RequestType = "PUT"
)

func (r RequestType) String() string {
	return string(r)
}

var NIL_REQBODY = []byte{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

var httpTimeout = 100 * time.Second

func SetHttpTimeout(timeout time.Duration) {
	httpTimeout = timeout
}

func GetPointer[T any](v T) *T {
	return &v
}

type MyPolymarket struct{}

const (
	PM_GAMMA_API_HTTP = "gamma-api.polymarket.com" // 市场、事件、标签、系列、评论、体育、搜索和公开个人资料。这是发现和浏览市场数据的主要 API。
	PM_DATA_API_HTTP  = "data-api.polymarket.com"  // 用户持仓、交易、活动、持有者数据、未平仓合约、排行榜和 Builder 分析。
	PM_CLOB_API_HTTP  = "clob.polymarket.com"      // 订单簿数据、价格、中间价、价差和价格历史。同时处理下单、撤单和其他交易操作。交易端点需要身份验证。

	IS_GZIP = true
)

type NetType int

const (
	MAIN_NET NetType = iota
	TEST_NET
)

type ChainType int64

const (
	POLYGON_MAINNET_CHAIN_ID ChainType = 137
)

func (c ChainType) Int64() int64 {
	return int64(c)
}

func (c ChainType) String() string {
	return strconv.FormatInt(int64(c), 10)
}

type APIType int

const (
	REST APIType = iota
	WEBSOCKET
)

type RestApiType string

const (
	GAMMA_REST RestApiType = "GAMMA"
	DATA_REST  RestApiType = "DATA"
	CLOB_REST  RestApiType = "CLOB"
)

type Client struct {
	Wallet       *Wallet
	ProxyAddress string
	ApiKeyCreds  *ApiKeyCreds
}

type RestClient struct {
	c *Client
}

type GammaRestClient RestClient
type CLOBRestClient RestClient
type DataRestClient RestClient

func (*MyPolymarket) NewGammaRestClient() *GammaRestClient {
	return &GammaRestClient{
		c: &Client{},
	}
}

func (*MyPolymarket) NewDataRestClient() *DataRestClient {
	client := &DataRestClient{
		c: &Client{},
	}
	return client
}

func (*MyPolymarket) NewCLOBRestClient(wallet *Wallet, proxyAddress string, nonce uint64) (*CLOBRestClient, error) {
	apiKeyCreds, err := wallet.DeriveApiKey(nonce)
	if err != nil {
		return nil, err
	}
	client := &CLOBRestClient{
		c: &Client{
			Wallet:       wallet,
			ProxyAddress: proxyAddress,
			ApiKeyCreds:  apiKeyCreds,
		},
	}
	return client, nil
}

// ApiKeyCreds 返回 CLOB 客户端已保存的 L2 凭证（CreateApiKey 成功后可用）。
func (c *CLOBRestClient) ApiKeyCreds() *ApiKeyCreds {
	if c == nil {
		return nil
	}
	rc := (*RestClient)(c)
	if rc.c == nil {
		return nil
	}
	return rc.c.ApiKeyCreds
}

// 通用接口调用
func pmCallAPI[T any](client *Client, url url.URL, reqBody []byte, method RequestType) (*PolyMarketRestRes[T], error) {
	body, err := Request(url.String(), reqBody, method, IS_GZIP)
	if err != nil {
		return nil, err
	}

	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}

	return res, res.handlerError()
}

// 通用 L2 鉴权接口调用
func pmCallAPIWithSecret[T any](client *Client, url url.URL, reqBody []byte, method RequestType) (*PolyMarketRestRes[T], error) {
	if client == nil || client.Wallet == nil || client.ApiKeyCreds == nil {
		return nil, fmt.Errorf("client, wallet or ApiKeyCreds is nil")
	}
	ts := time.Now().Unix()
	tsStr := strconv.FormatInt(ts, 10)

	sigL2, err := signL2(client.ApiKeyCreds.Secret, method.String(), tsStr, url, reqBody)
	if err != nil {
		return nil, err
	}

	body, err := RequestWithHeader(url.String(), reqBody, method,
		map[string]string{
			"POLY_ADDRESS":    client.Wallet.Signer.Hex(),
			"POLY_SIGNATURE":  sigL2,
			"POLY_TIMESTAMP":  tsStr,
			"POLY_API_KEY":    client.ApiKeyCreds.APIKey,
			"POLY_PASSPHRASE": client.ApiKeyCreds.Passphrase,
		}, IS_GZIP)

	if err != nil {
		return nil, err
	}

	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}

	return res, res.handlerError()
}

// pmCallAPIWithHeaders 使用自定义 HTTP Header 调用 REST（用于 CLOB L1 等需 POLY_* 头但不走默认 JSON 客户端体的场景）。
func pmCallAPIWithHeaders[T any](u url.URL, reqBody []byte, method RequestType, headers map[string]string) (*PolyMarketRestRes[T], error) {
	body, err := RequestWithHeader(u.String(), reqBody, method, headers, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

func PmGetRestHostByAPIType(apiType APIType, restApiType RestApiType) string {
	switch apiType {
	case REST:
		switch restApiType {
		case GAMMA_REST:
			return PM_GAMMA_API_HTTP
		case DATA_REST:
			return PM_DATA_API_HTTP
		case CLOB_REST:
			return PM_CLOB_API_HTTP
		}
	}
	return ""
}

// URL标准封装 不带路径参数
func pmHandlerRequestAPIWithoutPathQueryParam(apiType APIType, restApiType RestApiType, path string) url.URL {
	u := url.URL{
		Scheme:   "https",
		Host:     PmGetRestHostByAPIType(apiType, restApiType),
		Path:     path,
		RawQuery: "",
	}
	return u
}

// URL标准封装 带路径参数
func pmHandlerRequestAPIWithPathQueryParam[T any](apiType APIType, restApiType RestApiType, request *T, name string) url.URL {
	query := pmHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     PmGetRestHostByAPIType(apiType, restApiType),
		Path:     name,
		RawQuery: query,
	}
	return u
}

func pmHandlerReq[T any](req *T) string {
	var argBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			// 对字符串值做 URL 编码，防止含特殊字符时破坏 query string 结构。
			argBuffer.WriteString(argName + "=" + url.QueryEscape(v.Field(i).Elem().String()) + "&")
		case reflect.Int, reflect.Int64:
			argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			toStr := sv.MethodByName("String")
			if !toStr.IsValid() {
				log.Errorf("req struct field %s has no String() method, skipped", argName)
				continue
			}
			result := toStr.Call(nil)
			argBuffer.WriteString(argName + "=" + url.QueryEscape(result[0].String()) + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(argBuffer.String(), "&")
}
