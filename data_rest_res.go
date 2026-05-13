package mypolymarketapi

// DataGetLiveVolumeForEventRes 是 Get live volume for an event 接口的响应体。
type DataGetLiveVolumeForEventRes []DataLiveVolume

// DataLiveVolume 表示事件实时成交量聚合。
type DataLiveVolume struct {
	Total   float64           `json:"total"`   // number 实时总成交量
	Markets []DataMarketVolume `json:"markets"` // MarketVolume[] 各市场实时成交量
}

// DataMarketVolume 表示单个市场实时成交量。
type DataMarketVolume struct {
	Market string   `json:"market"` // string 市场标识（0x 开头 64 位十六进制哈希）
	Value  float64 `json:"value"`  // number 该市场实时成交量
}

// DataGetOpenInterestRes 是 Get open interest 接口的响应体。
type DataGetOpenInterestRes []DataOpenInterest

// DataOpenInterest 表示单个市场未平仓量。
type DataOpenInterest struct {
	Market string   `json:"market"` // string 市场标识（0x 开头 64 位十六进制哈希）
	Value  float64 `json:"value"`  // number 该市场未平仓量
}

// DataGetCurrentPositionsRes 是 Get current positions for a user 接口的响应体。
type DataGetCurrentPositionsRes []DataPosition

// DataPosition 表示用户当前持仓条目（OpenAPI `Position`）。
type DataPosition struct {
	ProxyWallet        string  `json:"proxyWallet"`        // string 用户 proxy 钱包地址
	Asset              string  `json:"asset"`              // string outcome 代币 asset id
	ConditionID        string  `json:"conditionId"`        // string 市场 condition id（0x + 64 hex）
	Size               float64 `json:"size"`               // number 持仓份额数量
	AvgPrice           float64 `json:"avgPrice"`           // number 持仓均价
	InitialValue       float64 `json:"initialValue"`       // number 初始价值
	CurrentValue       float64 `json:"currentValue"`       // number 当前价值
	CashPnl            float64 `json:"cashPnl"`            // number 现金盈亏
	PercentPnl         float64 `json:"percentPnl"`         // number 盈亏百分比
	TotalBought        float64 `json:"totalBought"`        // number 累计买入量
	RealizedPnl        float64 `json:"realizedPnl"`        // number 已实现盈亏
	PercentRealizedPnl float64 `json:"percentRealizedPnl"` // number 已实现盈亏百分比
	CurPrice           float64 `json:"curPrice"`           // number 当前价格
	Redeemable         bool    `json:"redeemable"`         // boolean 是否可赎回
	Mergeable          bool    `json:"mergeable"`          // boolean 是否可合并
	Title              string  `json:"title"`              // string 市场标题
	Slug               string  `json:"slug"`               // string 市场 slug
	Icon               string  `json:"icon"`               // string 图标 URL
	EventSlug          string  `json:"eventSlug"`          // string 事件 slug
	Outcome            string  `json:"outcome"`            // string 持仓方向文案
	OutcomeIndex       int64   `json:"outcomeIndex"`       // integer outcome 索引
	OppositeOutcome    string  `json:"oppositeOutcome"`    // string 对手方 outcome 文案
	OppositeAsset      string  `json:"oppositeAsset"`      // string 对手方 asset id
	EndDate            string  `json:"endDate"`            // string 结束/到期时间
	NegativeRisk       bool    `json:"negativeRisk"`       // boolean 是否为 negative risk 市场
}

// DataGetClosedPositionsRes 是 Get closed positions for a user 接口的响应体。
type DataGetClosedPositionsRes []DataClosedPosition

// DataClosedPosition 表示用户已平仓持仓条目（OpenAPI `ClosedPosition`）。
type DataClosedPosition struct {
	ProxyWallet     string  `json:"proxyWallet"`     // string 用户 proxy 钱包地址
	Asset           string  `json:"asset"`           // string outcome 代币 asset id
	ConditionID     string  `json:"conditionId"`     // string 市场 condition id（0x + 64 hex）
	AvgPrice        float64 `json:"avgPrice"`        // number 均价
	TotalBought     float64 `json:"totalBought"`     // number 累计买入量
	RealizedPnl     float64 `json:"realizedPnl"`     // number 已实现盈亏
	CurPrice        float64 `json:"curPrice"`        // number 当前价格
	Timestamp       int64   `json:"timestamp"`       // integer(int64) 时间戳（秒或毫秒以服务端为准）
	Title           string  `json:"title"`           // string 市场标题
	Slug            string  `json:"slug"`            // string 市场 slug
	Icon            string  `json:"icon"`            // string 图标 URL
	EventSlug       string  `json:"eventSlug"`       // string 事件 slug
	Outcome         string  `json:"outcome"`         // string 持仓方向文案
	OutcomeIndex    int64   `json:"outcomeIndex"`    // integer outcome 索引
	OppositeOutcome string  `json:"oppositeOutcome"` // string 对手方 outcome 文案
	OppositeAsset   string  `json:"oppositeAsset"`   // string 对手方 asset id
	EndDate         string  `json:"endDate"`         // string 结束/到期时间
}

// DataDownloadAccountingSnapshotRes 是 Download an accounting snapshot（ZIP of CSVs）接口的响应体。
// 结构拆分为 positions 与 equities 两类强类型数组。
type DataDownloadAccountingSnapshotRes struct {
	Positions []DataAccountingSnapshotPositionRow `json:"positions"` // [] positions.csv 解析结果
	Equities  []DataAccountingSnapshotEquityRow   `json:"equities"`  // [] equity.csv 解析结果
}

// DataAccountingSnapshotEquityRow 表示 equity.csv 的一行。
type DataAccountingSnapshotEquityRow struct {
	CashBalance        float64 `json:"cashBalance,omitempty"`        // number 现金余额
	PositionsValue     float64 `json:"positionsValue,omitempty"`     // number 持仓市值
	Equity             float64 `json:"equity,omitempty"`             // number 总权益
	ValuationTime      string  `json:"valuationTime,omitempty"`      // string 估值时间（RFC3339）
	ValuationTimestamp int64   `json:"valuationTimestamp,omitempty"` // integer 秒级时间戳
}

// DataAccountingSnapshotPositionRow 表示 positions.csv 的一行。
type DataAccountingSnapshotPositionRow struct {
	ConditionID        string  `json:"conditionId,omitempty"`        // string 市场 condition id
	Asset              string  `json:"asset,omitempty"`              // string outcome token asset id
	Size               float64 `json:"size,omitempty"`               // number 持仓数量
	CurPrice           float64 `json:"curPrice,omitempty"`           // number 当前价格
	ValuationTime      string  `json:"valuationTime,omitempty"`      // string 估值时间（RFC3339）
	ValuationTimestamp int64   `json:"valuationTimestamp,omitempty"` // integer 秒级时间戳
}
