package mypolymarketapi

type WsMarketType struct {
	EventType string `json:"event_type"`
}

// WsMarketOrderSummary 聚合订单簿某一档位的价格与数量。
type WsMarketOrderSummary struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

type WsMarketOrderBook struct {
	EventType string                 `json:"event_type"`
	AssetID   string                 `json:"asset_id"`
	Market    string                 `json:"market"`
	Bids      []WsMarketOrderSummary `json:"bids"`
	Asks      []WsMarketOrderSummary `json:"asks"`
	Timestamp string                 `json:"timestamp"`
	Hash      string                 `json:"hash"`
}

func handleWsMarketOrderBook(data []byte) []WsMarketOrderBook {
	// 服务端可能返回数组（initial_dump 常见）或单个对象
	for len(data) > 0 && (data[0] == ' ' || data[0] == '\n' || data[0] == '\r' || data[0] == '\t') {
		data = data[1:]
	}
	if len(data) == 0 {
		return nil
	}
	if data[0] == '{' {
		var ob WsMarketOrderBook
		if err := json.Unmarshal(data, &ob); err != nil {
			log.Error("unmarshal orderBook error: ", err)
			return nil
		}
		return []WsMarketOrderBook{ob}
	}

	var orderbooks []WsMarketOrderBook
	if err := json.Unmarshal(data, &orderbooks); err != nil {
		log.Error("unmarshal orderBook error: ", err)
		return nil
	}
	return orderbooks
}

type WsMarketPriceChange struct {
	AssetID string `json:"asset_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Side    string `json:"side"` // BUY / SELL
	Hash    string `json:"hash"`
	BestBid string `json:"best_bid,omitempty"`
	BestAsk string `json:"best_ask,omitempty"`
}

type WsMarketPriceChangeEvent struct {
	EventType    string                `json:"event_type"` // price_change
	Market       string                `json:"market"`
	PriceChanges []WsMarketPriceChange `json:"price_changes"`
	Timestamp    string                `json:"timestamp"`
}

func handleWsMarketPriceChange(data []byte) []WsMarketPriceChange {
	var evt WsMarketPriceChangeEvent
	err := json.Unmarshal(data, &evt)
	if err != nil {
		log.Error("unmarshal priceChangeEvent error: ", err)
		return nil
	}
	return evt.PriceChanges
}

type WsMarketLastTradePrice struct {
	EventType       string `json:"event_type"` // last_trade_price
	AssetID         string `json:"asset_id"`
	Market          string `json:"market"`
	Price           string `json:"price"`
	Size            string `json:"size"`
	FeeRateBps      string `json:"fee_rate_bps,omitempty"`
	Side            string `json:"side"` // BUY / SELL (from taker's perspective)
	Timestamp       string `json:"timestamp"`
	TransactionHash string `json:"transaction_hash,omitempty"`
}

func handleWsMarketLastTradePrice(data []byte) *WsMarketLastTradePrice {
	var lastTradePrice WsMarketLastTradePrice
	err := json.Unmarshal(data, &lastTradePrice)
	if err != nil {
		log.Error("unmarshal lastTradePrice error: ", err)
		return nil
	}
	return &lastTradePrice
}

type WsMarketTickSizeChange struct {
	EventType   string `json:"event_type"` // tick_size_change
	AssetID     string `json:"asset_id"`
	Market      string `json:"market"`
	OldTickSize string `json:"old_tick_size"`
	NewTickSize string `json:"new_tick_size"`
	Timestamp   string `json:"timestamp"`
}

func handleWsMarketTickSizeChange(data []byte) *WsMarketTickSizeChange {
	var tickSizeChange WsMarketTickSizeChange
	err := json.Unmarshal(data, &tickSizeChange)
	if err != nil {
		log.Error("unmarshal tickSizeChange error: ", err)
		return nil
	}
	return &tickSizeChange
}

type WsMarketBestBidAsk struct {
	EventType string `json:"event_type"` // best_bid_ask
	AssetID   string `json:"asset_id"`
	Market    string `json:"market"`
	BestBid   string `json:"best_bid"`
	BestAsk   string `json:"best_ask"`
	Spread    string `json:"spread"`
	Timestamp string `json:"timestamp"`
}

func handleWsMarketBestBidAsk(data []byte) *WsMarketBestBidAsk {
	var bestBidAsk WsMarketBestBidAsk
	err := json.Unmarshal(data, &bestBidAsk)
	if err != nil {
		log.Error("unmarshal bestBidAsk error: ", err)
		return nil
	}
	return &bestBidAsk
}

type WsMarketEventMessage struct {
	ID          string `json:"id,omitempty"`
	Ticker      string `json:"ticker,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type WsMarketNewMarket struct {
	EventType string `json:"event_type"` // new_market

	// required (per AsyncAPI)
	ID        string   `json:"id"`
	Question  string   `json:"question"`
	Market    string   `json:"market"`
	Slug      string   `json:"slug"`
	AssetsIDs []string `json:"assets_ids"`
	Outcomes  []string `json:"outcomes"`
	Timestamp string   `json:"timestamp"`

	// optional
	Description           string                `json:"description,omitempty"`
	EventMessage          *WsMarketEventMessage `json:"event_message,omitempty"`
	Tags                  []string              `json:"tags,omitempty"`
	ConditionID           string                `json:"condition_id,omitempty"`
	Active                *bool                 `json:"active,omitempty"`
	ClobTokenIDs          []string              `json:"clob_token_ids,omitempty"`
	SportsMarketType      string                `json:"sports_market_type,omitempty"`
	Line                  string                `json:"line,omitempty"`
	GameStartTime         string                `json:"game_start_time,omitempty"`
	OrderPriceMinTickSize string                `json:"order_price_min_tick_size,omitempty"`
	GroupItemTitle        string                `json:"group_item_title,omitempty"`
}

type WsMarketMarketResolved struct {
	EventType      string   `json:"event_type"` // market_resolved
	ID             string   `json:"id"`
	Market         string   `json:"market"`
	AssetsIDs      []string `json:"assets_ids"`
	WinningAssetID string   `json:"winning_asset_id"`
	WinningOutcome string   `json:"winning_outcome"`
	Timestamp      string   `json:"timestamp"`
	Tags           []string `json:"tags,omitempty"`
}

func handleWsMarketNewMarket(data []byte) *WsMarketNewMarket {
	var nm WsMarketNewMarket
	if err := json.Unmarshal(data, &nm); err != nil {
		log.Error("unmarshal newMarket error: ", err)
		return nil
	}
	return &nm
}

func handleWsMarketMarketResolved(data []byte) *WsMarketMarketResolved {
	var mr WsMarketMarketResolved
	if err := json.Unmarshal(data, &mr); err != nil {
		log.Error("unmarshal marketResolved error: ", err)
		return nil
	}
	return &mr
}
