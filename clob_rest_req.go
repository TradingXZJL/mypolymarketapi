package mypolymarketapi

import "strings"

// CLOBTokenIDReq 表示包含 token_id 的公共请求字段。
type CLOBTokenIDReq struct {
	TokenID *string `json:"token_id"` // string Token ID（asset ID）
}

type CLOBGetOrderBookAPI struct {
	client *CLOBRestClient
	req    *CLOBGetOrderBookReq
}

type CLOBGetOrderBookReq CLOBTokenIDReq

// string Token ID（asset ID）
func (api *CLOBGetOrderBookAPI) TokenID(tokenID string) *CLOBGetOrderBookAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

type CLOBGetOrderBooksAPI struct {
	client *CLOBRestClient
	req    *CLOBGetOrderBooksReq
}

type CLOBGetOrderBooksReq []CLOBGetOrderBookReq

func (api *CLOBGetOrderBooksAPI) AddOrderBookReq(req CLOBGetOrderBookReq) *CLOBGetOrderBooksAPI {
	if api.req == nil {
		api.req = &CLOBGetOrderBooksReq{}
	}
	*api.req = append(*api.req, req)
	return api
}

type CLOBGetMarketPriceAPI struct {
	client *CLOBRestClient
	req    *CLOBGetMarketPriceReq
}

type CLOBGetMarketPriceReq struct {
	TokenID *string `json:"token_id"` // string Token ID（asset ID）
	Side    *string `json:"side"`     // string 方向，BUY 或 SELL
}

// string Token ID（asset ID）
func (api *CLOBGetMarketPriceAPI) TokenID(tokenID string) *CLOBGetMarketPriceAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

// string 方向，BUY 或 SELL
func (api *CLOBGetMarketPriceAPI) Side(side string) *CLOBGetMarketPriceAPI {
	api.req.Side = GetPointer(side)
	return api
}

type CLOBGetMarketPricesGetAPI struct {
	client *CLOBRestClient
	req    *CLOBGetMarketPricesGetReq
}

type CLOBGetMarketPricesGetReq struct {
	TokenIDs *string `json:"token_ids"` // string 逗号分隔的 token_id 列表
	Sides    *string `json:"sides"`     // string 逗号分隔的 side 列表（与 token_ids 一一对应）
}

// []string token_id 列表，会自动拼接为逗号分隔字符串
func (api *CLOBGetMarketPricesGetAPI) TokenIDs(tokenIDs []string) *CLOBGetMarketPricesGetAPI {
	api.req.TokenIDs = GetPointer(strings.Join(tokenIDs, ","))
	return api
}

// []string side 列表，会自动拼接为逗号分隔字符串
func (api *CLOBGetMarketPricesGetAPI) Sides(sides []string) *CLOBGetMarketPricesGetAPI {
	api.req.Sides = GetPointer(strings.Join(sides, ","))
	return api
}

type CLOBGetMarketPricesPostAPI struct {
	client *CLOBRestClient
	req    *CLOBGetMarketPricesPostReq
}

type CLOBGetMarketPricesPostReq []CLOBGetMarketPricesPostItem

// CLOBGetMarketPricesPostItem 表示批量价格请求项。
type CLOBGetMarketPricesPostItem struct {
	TokenID *string `json:"token_id"` // string Token ID（asset ID）
	Side    *string `json:"side"`     // string 方向，BUY 或 SELL
}

// string,string 新增一条价格请求项（token_id + side）
func (api *CLOBGetMarketPricesPostAPI) AddPriceReq(tokenID, side string) *CLOBGetMarketPricesPostAPI {
	if api.req == nil {
		api.req = &CLOBGetMarketPricesPostReq{}
	}
	*api.req = append(*api.req, CLOBGetMarketPricesPostItem{
		TokenID: GetPointer(tokenID),
		Side:    GetPointer(side),
	})
	return api
}

type CLOBGetMidpointPriceAPI struct {
	client *CLOBRestClient
	req    *CLOBGetMidpointPriceReq
}

type CLOBGetMidpointPriceReq struct {
	TokenID *string `json:"token_id"` // string Token ID（asset ID）
}

// string Token ID（asset ID）
func (api *CLOBGetMidpointPriceAPI) TokenID(tokenID string) *CLOBGetMidpointPriceAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

type CLOBGetMidpointPricesGetAPI struct {
	client *CLOBRestClient
	req    *CLOBGetMidpointPricesGetReq
}

type CLOBGetMidpointPricesGetReq struct {
	TokenIDs *string `json:"token_ids"` // string 逗号分隔的 token_id 列表
}

// []string token_id 列表，会自动拼接为逗号分隔字符串
func (api *CLOBGetMidpointPricesGetAPI) TokenIDs(tokenIDs []string) *CLOBGetMidpointPricesGetAPI {
	api.req.TokenIDs = GetPointer(strings.Join(tokenIDs, ","))
	return api
}

type CLOBGetMidpointPricesPostAPI struct {
	client *CLOBRestClient
	req    *CLOBGetMidpointPricesPostReq
}

type CLOBGetMidpointPricesPostReq []CLOBGetMidpointPricesPostItem

// CLOBGetMidpointPricesPostItem 表示批量 midpoint 请求项。
type CLOBGetMidpointPricesPostItem struct {
	TokenID *string `json:"token_id"`       // string Token ID（asset ID）
	Side    *string `json:"side,omitempty"` // string 可选，BUY 或 SELL（不用于 midpoint 计算）
}

// string 新增一条 midpoint 请求项（仅 token_id）
func (api *CLOBGetMidpointPricesPostAPI) AddMidpointReq(tokenID string) *CLOBGetMidpointPricesPostAPI {
	if api.req == nil {
		api.req = &CLOBGetMidpointPricesPostReq{}
	}
	*api.req = append(*api.req, CLOBGetMidpointPricesPostItem{
		TokenID: GetPointer(tokenID),
	})
	return api
}

type CLOBGetSpreadAPI struct {
	client *CLOBRestClient
	req    *CLOBGetSpreadReq
}

type CLOBGetSpreadReq struct {
	TokenID *string `json:"token_id"` // string Token ID（asset ID）
}

// string Token ID（asset ID）
func (api *CLOBGetSpreadAPI) TokenID(tokenID string) *CLOBGetSpreadAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

type CLOBGetSpreadsAPI struct {
	client *CLOBRestClient
	req    *CLOBGetSpreadsReq
}

type CLOBGetSpreadsReq []CLOBGetSpreadsItem

// CLOBGetSpreadsItem 表示批量 spread 请求项。
type CLOBGetSpreadsItem struct {
	TokenID *string `json:"token_id"`       // string Token ID（asset ID）
	Side    *string `json:"side,omitempty"` // string 可选，BUY 或 SELL（不用于 spread 计算）
}

// string 新增一条 spread 请求项（仅 token_id）
func (api *CLOBGetSpreadsAPI) AddSpreadReq(tokenID string) *CLOBGetSpreadsAPI {
	if api.req == nil {
		api.req = &CLOBGetSpreadsReq{}
	}
	*api.req = append(*api.req, CLOBGetSpreadsItem{
		TokenID: GetPointer(tokenID),
	})
	return api
}

type CLOBGetLastTradePriceAPI struct {
	client *CLOBRestClient
	req    *CLOBGetLastTradePriceReq
}

type CLOBGetLastTradePriceReq struct {
	TokenID *string `json:"token_id"` // string Token ID（asset ID）
}

// string Token ID（asset ID）
func (api *CLOBGetLastTradePriceAPI) TokenID(tokenID string) *CLOBGetLastTradePriceAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

type CLOBGetLastTradesPricesGetAPI struct {
	client *CLOBRestClient
	req    *CLOBGetLastTradesPricesGetReq
}

type CLOBGetLastTradesPricesGetReq struct {
	TokenIDs *string `json:"token_ids"` // string 逗号分隔的 token_id 列表（单次最多 500）
}

// []string token_id 列表，会自动拼接为逗号分隔字符串
func (api *CLOBGetLastTradesPricesGetAPI) TokenIDs(tokenIDs []string) *CLOBGetLastTradesPricesGetAPI {
	api.req.TokenIDs = GetPointer(strings.Join(tokenIDs, ","))
	return api
}

type CLOBGetLastTradesPricesPostAPI struct {
	client *CLOBRestClient
	req    *CLOBGetLastTradesPricesPostReq
}

type CLOBGetLastTradesPricesPostReq []CLOBGetLastTradesPricesPostReqItem

// CLOBGetLastTradesPricesPostReqItem 表示批量最近成交价请求项（与 BookRequest 一致）。
type CLOBGetLastTradesPricesPostReqItem struct {
	TokenID *string `json:"token_id"`       // string Token ID（asset ID）
	Side    *string `json:"side,omitempty"` // string 可选，BUY 或 SELL
}

// string 新增一条最近成交价请求项（仅 token_id）
func (api *CLOBGetLastTradesPricesPostAPI) AddLastTradesPriceReq(tokenID string) *CLOBGetLastTradesPricesPostAPI {
	if api.req == nil {
		api.req = &CLOBGetLastTradesPricesPostReq{}
	}
	*api.req = append(*api.req, CLOBGetLastTradesPricesPostReqItem{
		TokenID: GetPointer(tokenID),
	})
	return api
}

type CLOBGetPricesHistoryAPI struct {
	client *CLOBRestClient
	req    *CLOBGetPricesHistoryReq
}

// CLOBGetPricesHistoryReq 对应 GET /prices-history 的 query（字段名与 OpenAPI 一致）。
type CLOBGetPricesHistoryReq struct {
	Market   *string  `json:"market"`   // string 必填，市场 asset id
	StartTs  *float64 `json:"startTs"`  // number 仅包含该 Unix 时间戳之后的点
	EndTs    *float64 `json:"endTs"`    // number 仅包含该 Unix 时间戳之前的点
	Interval *string  `json:"interval"` // string 聚合间隔：max、all、1m、1w、1d、6h、1h
	Fidelity *int64   `json:"fidelity"` // integer 精度（分钟），默认 1
}

// string 市场 asset id（必填）
func (api *CLOBGetPricesHistoryAPI) Market(market string) *CLOBGetPricesHistoryAPI {
	api.req.Market = GetPointer(market)
	return api
}

// float64 起始 Unix 时间戳（秒）
func (api *CLOBGetPricesHistoryAPI) StartTs(ts float64) *CLOBGetPricesHistoryAPI {
	api.req.StartTs = GetPointer(ts)
	return api
}

// float64 结束 Unix 时间戳（秒）
func (api *CLOBGetPricesHistoryAPI) EndTs(ts float64) *CLOBGetPricesHistoryAPI {
	api.req.EndTs = GetPointer(ts)
	return api
}

// string 聚合间隔：max、all、1m、1w、1d、6h、1h
func (api *CLOBGetPricesHistoryAPI) Interval(interval string) *CLOBGetPricesHistoryAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

// int64 精度（分钟）
func (api *CLOBGetPricesHistoryAPI) Fidelity(minutes int64) *CLOBGetPricesHistoryAPI {
	api.req.Fidelity = GetPointer(minutes)
	return api
}

type CLOBGetBatchPricesHistoryAPI struct {
	client *CLOBRestClient
	req    *CLOBGetBatchPricesHistoryReq
}

// CLOBGetBatchPricesHistoryReq 对应 POST /batch-prices-history 的 JSON body（最多 20 个 market）。
type CLOBGetBatchPricesHistoryReq struct {
	Markets  []string `json:"markets"`            // string[] 必填，market asset id 列表
	StartTs  *float64 `json:"start_ts,omitempty"` // number 起始 Unix 时间戳（秒）
	EndTs    *float64 `json:"end_ts,omitempty"`   // number 结束 Unix 时间戳（秒）
	Interval *string  `json:"interval,omitempty"` // string 聚合间隔
	Fidelity *int64   `json:"fidelity,omitempty"` // integer 精度（分钟）
}

// []string market asset id 列表（必填，单次最多 20）
func (api *CLOBGetBatchPricesHistoryAPI) Markets(markets []string) *CLOBGetBatchPricesHistoryAPI {
	api.req.Markets = markets
	return api
}

func (api *CLOBGetBatchPricesHistoryAPI) StartTs(ts float64) *CLOBGetBatchPricesHistoryAPI {
	api.req.StartTs = GetPointer(ts)
	return api
}

func (api *CLOBGetBatchPricesHistoryAPI) EndTs(ts float64) *CLOBGetBatchPricesHistoryAPI {
	api.req.EndTs = GetPointer(ts)
	return api
}

func (api *CLOBGetBatchPricesHistoryAPI) Interval(interval string) *CLOBGetBatchPricesHistoryAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

func (api *CLOBGetBatchPricesHistoryAPI) Fidelity(minutes int64) *CLOBGetBatchPricesHistoryAPI {
	api.req.Fidelity = GetPointer(minutes)
	return api
}

type CLOBGetFeeRateAPI struct {
	client *CLOBRestClient
	req    *CLOBGetFeeRateReq
}

// CLOBGetFeeRateReq 对应 GET /fee-rate 的可选 query。
type CLOBGetFeeRateReq struct {
	TokenID *string `json:"token_id"` // string Token ID（asset ID），可选
}

// string Token ID（asset ID）；不传则仅依赖服务端默认行为
func (api *CLOBGetFeeRateAPI) TokenID(tokenID string) *CLOBGetFeeRateAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

type CLOBGetFeeRateByPathAPI struct {
	client *CLOBRestClient
	req    *CLOBGetFeeRateByPathReq
}

type CLOBGetFeeRateByPathReq struct {
	TokenID *string `json:"token_id"` // 仅用于 path 替换，不参与 query
}

// string Token ID（asset ID），作为路径参数
func (api *CLOBGetFeeRateByPathAPI) TokenID(tokenID string) *CLOBGetFeeRateByPathAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

type CLOBGetTickSizeAPI struct {
	client *CLOBRestClient
	req    *CLOBGetTickSizeReq
}

// CLOBGetTickSizeReq 对应 GET /tick-size 的 query。
// token_id 可选（不传则由服务端默认行为决定）。
type CLOBGetTickSizeReq struct {
	TokenID *string `json:"token_id"` // string Token ID（asset ID），可选
}

// string Token ID（asset ID）
func (api *CLOBGetTickSizeAPI) TokenID(tokenID string) *CLOBGetTickSizeAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

type CLOBGetTickSizeByPathAPI struct {
	client *CLOBRestClient
	req    *CLOBGetTickSizeByPathReq
}

type CLOBGetTickSizeByPathReq struct {
	TokenID *string `json:"token_id"` // 仅用于 path 替换，不参与 query
}

// string Token ID（asset ID），作为路径参数
func (api *CLOBGetTickSizeByPathAPI) TokenID(tokenID string) *CLOBGetTickSizeByPathAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

type CLOBGetClobMarketInfoAPI struct {
	client *CLOBRestClient
	req    *CLOBGetClobMarketInfoReq
}

type CLOBGetClobMarketInfoReq struct {
	ConditionID *string `json:"condition_id"` // 仅用于 path 替换，不参与 query
}

// string 市场 condition ID（0x…），作为路径参数
func (api *CLOBGetClobMarketInfoAPI) ConditionID(conditionID string) *CLOBGetClobMarketInfoAPI {
	api.req.ConditionID = GetPointer(conditionID)
	return api
}

type CLOBGetServerTimeAPI struct {
	client *CLOBRestClient
}

type CLOBGetTradesAPI struct {
	client *CLOBRestClient
	req    *CLOBGetTradesReq
}

// CLOBGetTradesReq 对应 GET /trades 的筛选 query：
// maker_address 必填，其它字段均可选。
type CLOBGetTradesReq struct {
	ID           *string `json:"id"`            // string 交易 ID
	MakerAddress *string `json:"maker_address"` // string maker 链上地址
	Market       *string `json:"market"`        // string 市场 condition ID
	AssetID      *string `json:"asset_id"`      // string outcome token ID
	Before       *string `json:"before"`        // string 指定时间之前（字符串时间戳）
	After        *string `json:"after"`         // string 指定时间之后（字符串时间戳）
	NextCursor   *string `json:"next_cursor"`   // string 分页 cursor（base64）
}

func (api *CLOBGetTradesAPI) ID(id string) *CLOBGetTradesAPI {
	api.req.ID = GetPointer(id)
	return api
}

func (api *CLOBGetTradesAPI) MakerAddress(addr string) *CLOBGetTradesAPI {
	api.req.MakerAddress = GetPointer(addr)
	return api
}

func (api *CLOBGetTradesAPI) Market(conditionID string) *CLOBGetTradesAPI {
	api.req.Market = GetPointer(conditionID)
	return api
}

func (api *CLOBGetTradesAPI) AssetID(assetID string) *CLOBGetTradesAPI {
	api.req.AssetID = GetPointer(assetID)
	return api
}

func (api *CLOBGetTradesAPI) Before(before string) *CLOBGetTradesAPI {
	api.req.Before = GetPointer(before)
	return api
}

func (api *CLOBGetTradesAPI) After(after string) *CLOBGetTradesAPI {
	api.req.After = GetPointer(after)
	return api
}

func (api *CLOBGetTradesAPI) NextCursor(nextCursor string) *CLOBGetTradesAPI {
	api.req.NextCursor = GetPointer(nextCursor)
	return api
}

type CLOBGetBalanceAllowanceAPI struct {
	client *CLOBRestClient
	req    *CLOBGetBalanceAllowanceReq
}

// CLOBGetBalanceAllowanceReq 对应 GET /balance-allowance 的 query。
type CLOBGetBalanceAllowanceReq struct {
	AssetType     *string `json:"asset_type"`     // string 资产类型：COLLATERAL 或 CONDITIONAL
	TokenID       *string `json:"token_id"`       // string 条件代币 ID；asset_type=CONDITIONAL 时必填
	SignatureType *int64  `json:"signature_type"` // integer 可选：0=EOA、1=POLY_PROXY、2=POLY_GNOSIS_SAFE
}

// string 资产类型：COLLATERAL 或 CONDITIONAL
func (api *CLOBGetBalanceAllowanceAPI) AssetType(assetType string) *CLOBGetBalanceAllowanceAPI {
	api.req.AssetType = GetPointer(assetType)
	return api
}

// string 条件代币 ID；查询 CONDITIONAL 时必填
func (api *CLOBGetBalanceAllowanceAPI) TokenID(tokenID string) *CLOBGetBalanceAllowanceAPI {
	api.req.TokenID = GetPointer(tokenID)
	return api
}

// int64 可选签名类型：0=EOA、1=POLY_PROXY、2=POLY_GNOSIS_SAFE
func (api *CLOBGetBalanceAllowanceAPI) SignatureType(signatureType int64) *CLOBGetBalanceAllowanceAPI {
	api.req.SignatureType = GetPointer(signatureType)
	return api
}

type CLOBPostOrderAPI struct {
	client *CLOBRestClient
	req    *CLOBPostOrderReq
}

// CLOBPostOrderReq 对应 POST /order 的请求体。
type CLOBPostOrderReq struct {
	Order     *CLOBOrderReq `json:"order"`               // object 订单签名载荷
	Owner     *string       `json:"owner"`               // string API key 所有者 UUID
	OrderType *string       `json:"orderType,omitempty"` // string 可选：GTC、FOK、GTD、FAK
	PostOnly  *bool         `json:"postOnly"`            // boolean 与 clob-client-v2 一致，须参与 JSON（勿 omitempty）
	DeferExec *bool         `json:"deferExec"`           // boolean omit false 会导致与官方载荷不一致（*bool + omitempty 会省略 deferExec:false）
}

// CLOBOrderReq 表示订单签名明细（字段命名与上游 API 一致）。
type CLOBOrderReq struct {
	Maker         *string `json:"maker"`              // string maker 地址（通常为代理地址）
	Signer        *string `json:"signer"`             // string 签名地址
	TokenID       *string `json:"tokenId"`            // string Token ID（asset ID）
	MakerAmount   *string `json:"makerAmount"`        // string maker 数量（6 位精度定点）
	TakerAmount   *string `json:"takerAmount"`        // string taker 数量（6 位精度定点）
	Side          *string `json:"side"`               // string BUY 或 SELL
	Expiration    *string `json:"expiration"`         // string 过期时间（Unix 秒）
	Timestamp     *string `json:"timestamp"`          // string 下单时间（Unix 毫秒）
	Metadata      *string `json:"metadata,omitempty"` // string 保留字段
	Builder       *string `json:"builder"`            // string builder code（bytes32 hex）
	Signature     *string `json:"signature"`          // string EIP-712 签名
	Salt          *int64  `json:"salt"`               // integer 随机盐
	SignatureType *int64  `json:"signatureType"`      // integer 签名类型：0/1/2
}

// CLOBOrderReqBuilder 提供构建 order 字段的链式写法。
type CLOBOrderReqBuilder struct {
	order *CLOBOrderReq
}

// NewCLOBOrderReqBuilder 创建订单签名载荷构建器。
func NewCLOBOrderReqBuilder() *CLOBOrderReqBuilder {
	return &CLOBOrderReqBuilder{order: &CLOBOrderReq{}}
}

// Build 返回构建完成的订单签名载荷。
func (b *CLOBOrderReqBuilder) Build() CLOBOrderReq {
	return *b.order
}

func (b *CLOBOrderReqBuilder) Maker(v string) *CLOBOrderReqBuilder {
	b.order.Maker = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) Signer(v string) *CLOBOrderReqBuilder {
	b.order.Signer = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) TokenID(v string) *CLOBOrderReqBuilder {
	b.order.TokenID = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) MakerAmount(v string) *CLOBOrderReqBuilder {
	b.order.MakerAmount = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) TakerAmount(v string) *CLOBOrderReqBuilder {
	b.order.TakerAmount = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) Side(v string) *CLOBOrderReqBuilder {
	b.order.Side = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) Expiration(v string) *CLOBOrderReqBuilder {
	b.order.Expiration = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) Timestamp(v string) *CLOBOrderReqBuilder {
	b.order.Timestamp = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) Metadata(v string) *CLOBOrderReqBuilder {
	b.order.Metadata = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) Builder(v string) *CLOBOrderReqBuilder {
	b.order.Builder = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) Signature(v string) *CLOBOrderReqBuilder {
	b.order.Signature = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) Salt(v int64) *CLOBOrderReqBuilder {
	b.order.Salt = GetPointer(v)
	return b
}

func (b *CLOBOrderReqBuilder) SignatureType(v int64) *CLOBOrderReqBuilder {
	b.order.SignatureType = GetPointer(v)
	return b
}

// CLOBOrderReq 订单签名载荷
func (api *CLOBPostOrderAPI) Order(order CLOBOrderReq) *CLOBPostOrderAPI {
	api.req.Order = &order
	return api
}

// string API key 所有者 UUID
func (api *CLOBPostOrderAPI) Owner(owner string) *CLOBPostOrderAPI {
	api.req.Owner = GetPointer(owner)
	return api
}

// string 可选：GTC、FOK、GTD、FAK
func (api *CLOBPostOrderAPI) OrderType(orderType string) *CLOBPostOrderAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}

// bool post-only（仅做 maker，不匹配），默认 false。
func (api *CLOBPostOrderAPI) PostOnly(postOnly bool) *CLOBPostOrderAPI {
	api.req.PostOnly = GetPointer(postOnly)
	return api
}

// bool 可选，是否延迟执行
func (api *CLOBPostOrderAPI) DeferExec(deferExec bool) *CLOBPostOrderAPI {
	api.req.DeferExec = GetPointer(deferExec)
	return api
}

type CLOBPostMultipleOrdersAPI struct {
	client *CLOBRestClient
	req    *CLOBPostMultipleOrdersReq
}

// CLOBPostMultipleOrdersReq 对应 POST /orders 的请求体（单次最多 15 笔）。
type CLOBPostMultipleOrdersReq []CLOBPostOrderReq

// CLOBPostOrderReq 新增一条批量下单请求项。
func (api *CLOBPostMultipleOrdersAPI) AddOrderReq(req CLOBPostOrderReq) *CLOBPostMultipleOrdersAPI {
	if api.req == nil {
		api.req = &CLOBPostMultipleOrdersReq{}
	}
	*api.req = append(*api.req, req)
	return api
}

type CLOBCancelSingleOrderAPI struct {
	client *CLOBRestClient
	req    *CLOBCancelSingleOrderReq
}

// CLOBCancelSingleOrderReq 对应 DELETE /order 的请求体。
type CLOBCancelSingleOrderReq struct {
	OrderID *string `json:"orderID"` // string 订单 ID（order hash）
}

// string 订单 ID（order hash）
func (api *CLOBCancelSingleOrderAPI) OrderID(orderID string) *CLOBCancelSingleOrderAPI {
	api.req.OrderID = GetPointer(orderID)
	return api
}

type CLOBCancelMultipleOrdersAPI struct {
	client *CLOBRestClient
	req    *CLOBCancelMultipleOrdersReq
}

// CLOBCancelMultipleOrdersReq 对应 DELETE /orders 的请求体（订单 ID 列表，单次最多 3000）。
type CLOBCancelMultipleOrdersReq []string

// string 新增一个待撤销订单 ID
func (api *CLOBCancelMultipleOrdersAPI) AddOrderID(orderID string) *CLOBCancelMultipleOrdersAPI {
	if api.req == nil {
		api.req = &CLOBCancelMultipleOrdersReq{}
	}
	*api.req = append(*api.req, orderID)
	return api
}

type CLOBCancelAllOrdersAPI struct {
	client *CLOBRestClient
}

type CLOBGetSingleOrderByIDAPI struct {
	client *CLOBRestClient
	req    *CLOBGetSingleOrderByIDReq
}

// CLOBGetSingleOrderByIDReq 对应 GET /order/{orderID} 的 path 参数。
type CLOBGetSingleOrderByIDReq struct {
	OrderID *string `json:"order_id"` // 仅用于 path 替换，不参与 query
}

// string 订单 ID（order hash），作为路径参数
func (api *CLOBGetSingleOrderByIDAPI) OrderID(orderID string) *CLOBGetSingleOrderByIDAPI {
	api.req.OrderID = GetPointer(orderID)
	return api
}

type CLOBGetUserOrdersAPI struct {
	client *CLOBRestClient
	req    *CLOBGetUserOrdersReq
}

// CLOBGetUserOrdersReq 对应 GET /data/orders 的筛选 query（均为可选）。
type CLOBGetUserOrdersReq struct {
	ID         *string `json:"id"`          // string 订单 ID（order hash）
	Market     *string `json:"market"`      // string 市场 condition ID
	AssetID    *string `json:"asset_id"`    // string outcome token ID
	NextCursor *string `json:"next_cursor"` // string 分页 cursor（base64）
}

func (api *CLOBGetUserOrdersAPI) ID(id string) *CLOBGetUserOrdersAPI {
	api.req.ID = GetPointer(id)
	return api
}

func (api *CLOBGetUserOrdersAPI) Market(conditionID string) *CLOBGetUserOrdersAPI {
	api.req.Market = GetPointer(conditionID)
	return api
}

func (api *CLOBGetUserOrdersAPI) AssetID(assetID string) *CLOBGetUserOrdersAPI {
	api.req.AssetID = GetPointer(assetID)
	return api
}

func (api *CLOBGetUserOrdersAPI) NextCursor(nextCursor string) *CLOBGetUserOrdersAPI {
	api.req.NextCursor = GetPointer(nextCursor)
	return api
}

type CLOBGetOrderScoringAPI struct {
	client *CLOBRestClient
	req    *CLOBGetOrderScoringReq
}

// CLOBGetOrderScoringReq 对应 GET /order-scoring 的 query。
type CLOBGetOrderScoringReq struct {
	OrderID *string `json:"order_id"` // string 订单 ID（order hash），必填
}

// string 订单 ID（order hash）
func (api *CLOBGetOrderScoringAPI) OrderID(orderID string) *CLOBGetOrderScoringAPI {
	api.req.OrderID = GetPointer(orderID)
	return api
}
