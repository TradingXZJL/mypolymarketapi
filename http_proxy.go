package mypolymarketapi

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

// ProxyWeight 记录单个代理 IP 的剩余权重与限流状态，所有字段通过内置互斥锁保护。
type ProxyWeight struct {
	mu           sync.Mutex
	RemainWeight int64
	IsLimited    bool
}

func (w *ProxyWeight) restore() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.RemainWeight = 1200
	w.IsLimited = false
}

// consume 原子地扣减一次权重，返回 false 表示权重已耗尽。
func (w *ProxyWeight) consume() bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.IsLimited || w.RemainWeight <= 0 {
		w.IsLimited = true
		return false
	}
	w.RemainWeight--
	if w.RemainWeight <= 0 {
		w.IsLimited = true
	}
	return true
}

func (w *ProxyWeight) setLimited() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.IsLimited = true
}

// available 原子地返回该代理是否可用及当前剩余权重。
func (w *ProxyWeight) available() (bool, int64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return !w.IsLimited && w.RemainWeight > 0, w.RemainWeight
}

type RestProxy struct {
	ProxyUrl string
	Weight   ProxyWeight
}

var (
	proxyListMu sync.RWMutex
	proxyList   []*RestProxy
)

func GetCurrentProxyList() []*RestProxy {
	proxyListMu.RLock()
	defer proxyListMu.RUnlock()
	result := make([]*RestProxy, len(proxyList))
	copy(result, proxyList)
	return result
}

var UseProxy = false
var WsUseProxy = false

func SetUseProxy(useProxy bool, proxyUrls ...string) {
	UseProxy = useProxy
	newList := make([]*RestProxy, 0, len(proxyUrls))
	for _, pu := range proxyUrls {
		newList = append(newList, &RestProxy{
			ProxyUrl: pu,
			Weight: ProxyWeight{
				RemainWeight: 1200,
			},
		})
	}
	proxyListMu.Lock()
	proxyList = newList
	proxyListMu.Unlock()
}

func SetWsUseProxy(useProxy bool) error {
	if !UseProxy {
		return errors.New("please set UseProxy first")
	}
	WsUseProxy = useProxy
	return nil
}

func isUseProxy() bool {
	return UseProxy
}

func init() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("0 */1 * * * *", func() {
		proxyListMu.RLock()
		defer proxyListMu.RUnlock()
		for _, proxy := range proxyList {
			proxy.Weight.restore()
		}
	})
	if err != nil {
		log.Error(err)
		return
	}
	c.Start()
}

func getBestProxyAndWeight() (*RestProxy, *ProxyWeight) {
	proxyListMu.RLock()
	defer proxyListMu.RUnlock()

	var best *RestProxy
	var bestWeight *ProxyWeight
	var bestRemain int64 = -1

	for _, proxy := range proxyList {
		avail, remain := proxy.Weight.available()
		if !avail {
			continue
		}
		if best == nil || remain > bestRemain {
			best = proxy
			bestWeight = &proxy.Weight
			bestRemain = remain
		}
	}
	return best, bestWeight
}

func getRandomProxy() (*RestProxy, error) {
	proxyListMu.RLock()
	defer proxyListMu.RUnlock()
	n := len(proxyList)
	if n == 0 {
		return nil, errors.New("proxyList is empty")
	}
	return proxyList[rand.Intn(n)], nil
}

// sensitiveHeaders 是 Debug 日志中需脱敏的请求头集合。
var sensitiveHeaders = map[string]bool{
	"POLY_SIGNATURE":  true,
	"POLY_API_KEY":    true,
	"POLY_PASSPHRASE": true,
	"POLY_SECRET":     true,
}

// sanitizeHeaders 返回脱敏后的 Header 副本（仅用于日志输出）。
func sanitizeHeaders(h http.Header) http.Header {
	out := make(http.Header, len(h))
	for k, v := range h {
		if sensitiveHeaders[k] {
			out[k] = []string{"[REDACTED]"}
		} else {
			out[k] = v
		}
	}
	return out
}

func Request(rawURL string, reqBody []byte, method RequestType, isGzip bool) ([]byte, error) {
	return RequestWithHeader(rawURL, reqBody, method, map[string]string{}, isGzip)
}

func RequestWithHeader(urlStr string, reqBody []byte, method RequestType, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest(method.String(), urlStr, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headerMap {
		req.Header[k] = []string{v}
	}
	req.Header.Set("Content-Type", "application/json")

	// 只请求服务端返回 gzip 压缩响应；请求体本身不压缩，不设 Content-Encoding。
	if isGzip {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	if log.IsLevelEnabled(logrus.DebugLevel) {
		log.Debug("reqHeader: ", sanitizeHeaders(req.Header))
	}
	log.Debug("reqURL: ", req.URL.String())
	if len(reqBody) > 0 {
		log.Debug("reqBody: ", string(reqBody))
		req.Body = io.NopCloser(bytes.NewReader(reqBody))
	}

	// 每次请求使用独立的 http.Client，并注入全局超时。
	client := &http.Client{Timeout: httpTimeout}

	if UseProxy {
		proxy, proxyWeight := getBestProxyAndWeight()
		if proxy == nil || proxyWeight == nil {
			return nil, errors.New("all proxy ip weight limit reached")
		}

		proxyURL, err := url.Parse(proxy.ProxyUrl)
		if err != nil {
			return nil, err
		}
		// 不禁用 TLS 证书校验，保持完整的传输层安全。
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == http.StatusTooEarly {
			resp.Body.Close()
			return nil, errors.New("[425] The matching engine is restarting.")
		}
		defer resp.Body.Close()

		data, readErr := readResponseBody(resp)
		log.Debug("respHeader: ", resp.Header)
		log.Debug("respBody: ", string(data))

		if resp.StatusCode == http.StatusTooManyRequests {
			proxyWeight.setLimited()
			return data, errors.New("proxy ip weight limit reached")
		}
		if !proxyWeight.consume() {
			return data, errors.New("proxy ip weight limit reached")
		}
		return data, readErr
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusTooEarly {
		resp.Body.Close()
		return nil, errors.New("[425] The matching engine is restarting.")
	}
	defer resp.Body.Close()

	data, readErr := readResponseBody(resp)
	log.Debug("respHeader: ", resp.Header)
	log.Debug("respBody: ", string(data))
	return data, readErr
}

// readResponseBody 读取响应体，按需解压 gzip。
func readResponseBody(resp *http.Response) ([]byte, error) {
	var bodyReader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = gr
	}
	return io.ReadAll(bodyReader)
}
