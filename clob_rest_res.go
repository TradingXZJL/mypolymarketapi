package mypolymarketapi

// CLOBGetOrderBookRes 是 Get order book 接口的响应体。
type CLOBGetOrderBookRes CLOBOrderBookSummary

// CLOBGetOrderBooksRes 是 Get order books（request body）接口的响应体。
type CLOBGetOrderBooksRes []CLOBOrderBookSummary

// CLOBOrderBookSummary 表示订单簿快照摘要。
type CLOBOrderBookSummary struct {
	Market         string  `json:"market"`           // string 市场 condition ID
	AssetID        string  `json:"asset_id"`         // string Token ID（asset ID）
	Timestamp      string  `json:"timestamp"`        // string 订单簿快照时间戳
	Hash           string  `json:"hash"`             // string 订单簿摘要哈希
	Bids           []Books `json:"bids"`             // Books[] 买单列表（按价格降序）
	Asks           []Books `json:"asks"`             // Books[] 卖单列表（按价格升序）
	MinOrderSize   string  `json:"min_order_size"`   // string 最小下单数量
	TickSize       string  `json:"tick_size"`        // string 最小价格变动单位
	NegRisk        bool   `json:"neg_risk"`         // boolean 是否启用 negative risk
	LastTradePrice string  `json:"last_trade_price"` // string 最近成交价
}

// Books 表示订单簿中的订单。
type Books struct {
	Price string `json:"price"` // string 订单价格
	Size  string `json:"size"`  // string 订单数量
}

// CLOBGetMarketPriceRes 是 Get market price 接口的响应体。
type CLOBGetMarketPriceRes struct {
	Price float64 `json:"price"` // number 市场价格
}

// CLOBMarketPricesMap 表示 token_id 到 side 到价格的映射。
type CLOBMarketPricesMap map[string]map[string]float64

// CLOBGetMarketPricesGetRes 是 Get market prices（query parameters）接口的响应体。
type CLOBGetMarketPricesGetRes CLOBMarketPricesMap

// CLOBGetMarketPricesPostRes 是 Get market prices（request body）接口的响应体。
type CLOBGetMarketPricesPostRes CLOBMarketPricesMap

// CLOBGetMidpointPriceRes 是 Get midpoint price 接口的响应体。
type CLOBGetMidpointPriceRes struct {
	MidPrice string `json:"mid_price"` // string midpoint 价格
}

// CLOBMidpointPricesMap 表示 token_id 到 midpoint 价格的映射。
type CLOBMidpointPricesMap map[string]string

// CLOBGetMidpointPricesGetRes 是 Get midpoint prices（query parameters）接口的响应体。
type CLOBGetMidpointPricesGetRes CLOBMidpointPricesMap

// CLOBGetMidpointPricesPostRes 是 Get midpoint prices（request body）接口的响应体。
type CLOBGetMidpointPricesPostRes CLOBMidpointPricesMap

// CLOBGetSpreadRes 是 Get spread 接口的响应体。
type CLOBGetSpreadRes struct {
	Spread string `json:"spread"` // string spread 值
}

// CLOBSpreadsMap 表示 token_id 到 spread 的映射。
type CLOBSpreadsMap map[string]string

// CLOBGetSpreadsRes 是 Get spreads 接口的响应体。
type CLOBGetSpreadsRes CLOBSpreadsMap

// CLOBGetLastTradePriceRes 是 Get last trade price 接口的响应体。
// 无成交时 API 可能返回 price "0.5"、side 为空字符串。
type CLOBGetLastTradePriceRes struct {
	Price string `json:"price"` // string 最近成交价
	Side  string `json:"side"`  // string 最近成交方向：BUY、SELL 或空
}

// CLOBLastTradePriceItem 表示单个 token 的最近成交价与方向。
type CLOBLastTradePriceItem struct {
	TokenID string `json:"token_id"` // string Token ID（asset ID）
	Price   string `json:"price"`    // string 最近成交价
	Side    string `json:"side"`     // string 最近成交方向：BUY 或 SELL
}

// CLOBGetLastTradesPricesGetRes 是 Get last trade prices（query parameters）接口的响应体。
type CLOBGetLastTradesPricesGetRes []CLOBLastTradePriceItem

// CLOBGetLastTradesPricesPostRes 是 Get last trade prices（request body）接口的响应体。
type CLOBGetLastTradesPricesPostRes []CLOBLastTradePriceItem

// CLOBMarketPricePoint 表示价格历史曲线上的单点（时间戳 + 价格）。
type CLOBMarketPricePoint struct {
	T int64   `json:"t"` // integer Unix 时间戳（秒）
	P float64 `json:"p"` // number 价格
}

// CLOBGetPricesHistoryRes 是 Get prices history 接口的响应体。
type CLOBGetPricesHistoryRes struct {
	History []CLOBMarketPricePoint `json:"history"`
}

// CLOBGetBatchPricesHistoryRes 是 Get batch prices history 接口的响应体。
type CLOBGetBatchPricesHistoryRes struct {
	History map[string][]CLOBMarketPricePoint `json:"history"` // market asset id -> 价格点序列
}

// CLOBGetFeeRateRes 是 Get fee rate 接口的响应体（query 与 path 两种调用返回一致）。
type CLOBGetFeeRateRes struct {
	BaseFee int64 `json:"base_fee"` // integer 基础费率（basis points）
}

// CLOBGetTickSizeRes 是 Get tick size 接口的响应体（query 与 path 两种调用返回一致）。
type CLOBGetTickSizeRes struct {
	MinimumTickSize float64 `json:"minimum_tick_size"` // number 最小价格变动单位
}

// CLOBClobMarketInfoToken 表示 CLOB 市场信息中的单条 outcome 代币（JSON 短键 t / o）。
type CLOBClobMarketInfoToken struct {
	TokenID string `json:"t"` // string outcome 代币 ID
	Outcome string `json:"o"` // string 方向文案（如 Yes、No）
}

// CLOBGetClobMarketInfoRes 是 GET /clob-markets/{condition_id} 的响应体。字段以线上 API 短键名为主，未出现的费项等可能缺省。
type CLOBGetClobMarketInfoRes struct {
	GST   string                   `json:"gst,omitempty"`   // 体育类市场的比赛开始时间（ISO 8601），可空
	R     any                       `json:"r,omitempty"`     // 奖励/激励配置，结构因市场可能不同
	T     []CLOBClobMarketInfoToken `json:"t,omitempty"`     // 市场 outcome 与对应 token
	C     string                    `json:"c,omitempty"`     // 市场 condition ID
	MOS   float64                   `json:"mos"`             // 最小下单数量
	MTS   float64                   `json:"mts"`             // 最小价格变动（tick）
	MBF   int64                    `json:"mbf,omitempty"`   // maker 基础费（bps）
	TBF   int64                    `json:"tbf,omitempty"`   // taker 基础费（bps）
	RFQE  bool                     `json:"rfqe,omitempty"`  // 是否开启 RFQ
	Itode bool                     `json:"itode,omitempty"` // 是否开启 taker 订单延迟
	IBCE  bool                     `json:"ibce,omitempty"`  // 是否开启 Blockaid 校验
	FD    any                       `json:"fd,omitempty"`    // 手续费曲线等详情
	OAS   int64                    `json:"oas,omitempty"`   // 最小订单存活时间（秒）等
	AO    bool                     `json:"ao,omitempty"`    // 是否正在接受订单
	CBOS  bool                     `json:"cbos,omitempty"`  // 是否在开始时清空订单簿
	AOT   string                   `json:"aot,omitempty"`   // 开始接受订单的时间（ISO 8601）
}

// CLOBGetServerTimeRes 表示 GET /time 返回的 Unix 时间戳（秒）；HTTP 体为裸 JSON number。
type CLOBGetServerTimeRes int64

// CLOBTradeMakerOrder 表示单个成交里参与撮合的 maker 订单明细。
type CLOBTradeMakerOrder struct {
	OrderID       string `json:"order_id"`       // string maker 订单 ID
	Owner         string `json:"owner"`          // string maker 所属 API key
	MakerAddress  string `json:"maker_address"`  // string maker 链上地址
	MatchedAmount string `json:"matched_amount"` // string 成交数量
	Price         string `json:"price"`          // string 成交价格
	FeeRateBps    string `json:"fee_rate_bps"`   // string 费率（bps）
	AssetID       string `json:"asset_id"`       // string token ID
	Outcome       string `json:"outcome"`        // string outcome 文案（如 Yes/No）
	Side          string `json:"side"`           // string BUY / SELL
}

// CLOBTradeItem 表示单条成交记录。
type CLOBTradeItem struct {
	ID              string                `json:"id"`               // string 交易 ID
	TakerOrderID    string                `json:"taker_order_id"`   // string taker 订单 ID
	Market          string                `json:"market"`           // string 市场 condition ID
	AssetID         string                `json:"asset_id"`         // string token ID
	Side            string                `json:"side"`             // string BUY / SELL
	Size            string                `json:"size"`             // string 成交数量
	FeeRateBps      string                `json:"fee_rate_bps"`     // string 费率（bps）
	Price           string                `json:"price"`            // string 成交价格
	Status          string                `json:"status"`           // string 状态
	MatchTime       string                `json:"match_time"`       // string 成交时间
	LastUpdate      string                `json:"last_update"`      // string 更新时间
	Outcome         string                `json:"outcome"`          // string outcome 文案（如 Yes/No）
	BucketIndex     int64                `json:"bucket_index"`     // integer 分桶索引
	Owner           string                `json:"owner"`            // string 交易所属 API key
	MakerAddress    string                `json:"maker_address"`    // string maker 链上地址
	MakerOrders     []CLOBTradeMakerOrder `json:"maker_orders"`     // MakerOrder[] 参与撮合的 maker 订单
	TransactionHash string                `json:"transaction_hash"` // string 交易哈希
}

// CLOBGetTradesRes 是 Get trades 接口的响应体（分页对象）。
type CLOBGetTradesRes struct {
	Limit      int64           `json:"limit"`        // integer 每页返回上限
	NextCursor string          `json:"next_cursor"`  // string 分页游标（空字符串表示无下一页）
	Count      int64           `json:"count"`        // integer 当前页返回数量
	Data       []CLOBTradeItem `json:"data"`        // Trade[] 当前页成交记录
}

// CLOBGetBalanceAllowanceRes 是 GET /balance-allowance 的响应体。
type CLOBGetBalanceAllowanceRes struct {
	Balance    string            `json:"balance"`              // string 当前可用余额
	Allowances map[string]string `json:"allowances,omitempty"` // map[spender]amount 各 spender 授权额度
	Allowance  string            `json:"allowance,omitempty"`  // string 兼容旧字段（若服务端返回单值授权）
}

// CLOBApiCredentialsRes 是 L1 Create / Derive API key 的响应体（JSON 字段 `apiKey`、`secret`、`passphrase`，用于后续 L2 POLY_* 鉴权）。
type CLOBApiCredentialsRes struct {
	APIKey     string `json:"apiKey"`     // string API Key（UUID）
	Secret     string `json:"secret"`     // string Secret（Base64 等编码，以服务端返回为准）
	Passphrase string `json:"passphrase"` // string Passphrase
}

// CLOBPostOrderRes 是 Post a new order 接口的响应体。
type CLOBPostOrderRes struct {
	Success            bool     `json:"success"`                      // boolean 是否下单成功
	OrderID            string   `json:"orderID"`                      // string 订单 ID（order hash）
	Status             string   `json:"status"`                       // string 订单状态：live、matched、delayed
	MakingAmount       string  `json:"makingAmount,omitempty"`       // string 成功时返回 maker 数量
	TakingAmount       string  `json:"takingAmount,omitempty"`       // string 成功时返回 taker 数量
	TransactionsHashes []string `json:"transactionsHashes,omitempty"` // string[] matched 时返回链上交易哈希
	TradeIDs           []string `json:"tradeIDs,omitempty"`           // string[] matched 时返回成交 ID
	ErrorMsg           string  `json:"errorMsg,omitempty"`           // string 失败或部分失败时返回错误信息
}

// CLOBPostMultipleOrdersRes 是 Post multiple orders 接口的响应体。
type CLOBPostMultipleOrdersRes []CLOBPostOrderRes

// CLOBCancelSingleOrderRes 是 Cancel single order 接口的响应体。
type CLOBCancelSingleOrderRes struct {
	Canceled    []string          `json:"canceled"`     // string[] 已成功撤销的订单 ID 列表
	NotCanceled map[string]string `json:"not_canceled"` // map[order_id]error 未成功撤销的订单及原因
}

// CLOBCancelMultipleOrdersRes 是 Cancel multiple orders 接口的响应体。
type CLOBCancelMultipleOrdersRes CLOBCancelSingleOrderRes

// CLOBCancelAllOrdersRes 是 Cancel all orders 接口的响应体。
type CLOBCancelAllOrdersRes CLOBCancelSingleOrderRes

// CLOBGetSingleOrderByIDRes 是 Get single order by ID 接口的响应体。
type CLOBGetSingleOrderByIDRes struct {
	ID             string   `json:"id"`               // string 订单 ID（order hash）
	Status         string   `json:"status"`           // string 订单状态
	Owner          string   `json:"owner"`            // string 订单所属 API key owner（UUID）
	MakerAddress   string   `json:"maker_address"`    // string maker 链上地址
	Market         string   `json:"market"`           // string 市场 condition ID
	AssetID        string   `json:"asset_id"`         // string Token ID（asset ID）
	Side           string   `json:"side"`             // string 方向：BUY 或 SELL
	OriginalSize   string   `json:"original_size"`    // string 初始下单数量（6 位定点）
	SizeMatched    string   `json:"size_matched"`     // string 已成交数量（6 位定点）
	Price          string   `json:"price"`            // string 下单价格
	Outcome        string   `json:"outcome"`          // string outcome 文案（如 YES/NO）
	Expiration     string   `json:"expiration"`       // string 过期时间（Unix 秒）
	OrderType      string   `json:"order_type"`       // string 时效类型：GTC、FOK、GTD、FAK
	AssociateTrade []string `json:"associate_trades"` // string[] 关联成交 ID 列表
	CreatedAt      int64    `json:"created_at"`       // integer 创建时间（Unix 秒）
}

// CLOBGetUserOrdersRes 是 Get user orders 接口的响应体。
type CLOBGetUserOrdersRes struct {
	Limit      int64                       `json:"limit"`       // integer 每页返回上限
	NextCursor string                      `json:"next_cursor"` // string 分页游标（空字符串表示无下一页）
	Count      int64                       `json:"count"`       // integer 当前页订单数量
	Data       []CLOBGetSingleOrderByIDRes `json:"data"`        // Order[] 当前页订单列表
}

// CLOBGetOrderScoringRes 是 Get order scoring status 接口的响应体。
type CLOBGetOrderScoringRes struct {
	Scoring bool `json:"scoring"` // boolean 是否正在参与奖励计分
}
