package mypolymarketapi

import (
	"errors"
)

type EventType string

const (
	EventTypeOrderbook      EventType = "orderbook"
	EventTypePriceChange    EventType = "price_change"
	EventTypeLastTradePrice EventType = "last_trade_price"
	EventTypeTickSizeChange EventType = "tick_size_change"
	EventTypeBestBidAsk     EventType = "best_bid_ask"
	EventTypeNewMarket      EventType = "new_market"
	EventTypeMarketResolved EventType = "market_resolved"
)

func (e EventType) String() string {
	return string(e)
}

// UpdateMarket 动态订阅或取消资产（operation: subscribe / unsubscribe），不占用 waitSubResult。
func (ws *MarketWsStreamClient) SubscribeMarket(assetIDs []string, customFeature bool, level int) (*Subscription, error) {
	if len(assetIDs) == 0 {
		return nil, errors.New("assets_ids is empty")
	}
	payload := WsSubscribeReq{
		AssetsIDs:            assetIDs,
		Operation:            SUBSCRIBE,
		Level:                level,
		CustomFeatureEnabled: customFeature,
	}
	doSub, err := subscribe(&ws.WsStreamClient, payload)
	if err != nil {
		return nil, err
	}
	sub := &Subscription{
		SubId: doSub.SubId,
		Ws:    &ws.WsStreamClient,
		Args:  doSub.Args,

		resultChan: make(chan any, 1024),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
	}
	for _, assetID := range assetIDs {
		ws.commonSubMap.Store(assetID, sub)
	}
	if err := ws.catchSubscribeResult(sub); err != nil {
		return nil, err
	}
	log.Infof("SubscribeMarket Success: args:%v", doSub.Args)

	return sub, nil
}

func (ws *MarketWsStreamClient) UnsubscribeMarket(assetIDs []string) error {
	if len(assetIDs) == 0 {
		return errors.New("assets_ids is empty")
	}
	payload := WsSubscribeReq{
		AssetsIDs: assetIDs,
		Operation: UNSUBSCRIBE,
	}
	doSub, err := subscribe(&ws.WsStreamClient, payload)
	if err != nil {
		return err
	}
	for _, assetID := range assetIDs {
		ws.commonSubMap.Delete(assetID)
	}
	log.Infof("UnsubscribeMarket Success: args:%v", doSub.Args)
	return nil
}
