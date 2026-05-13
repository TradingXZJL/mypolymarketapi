package mypolymarketapi

type CLOBAPIType int

const (
	// Orderbook & Pricing
	CLOBGetOrderBook            CLOBAPIType = iota // GET Get order book
	CLOBGetOrderBooks                              // POST Get order books (request body)
	CLOBGetMarketPrice                             // GET Get market price
	CLOBGetMarketPricesGet                         // GET Get market prices (query parameters)
	CLOBGetMarketPricesPost                        // POST Get market prices (request body)
	CLOBGetMidpointPrice                           // GET Get midpoint price
	CLOBGetMidpointPricesGet                       // GET Get midpoint prices (query parameters)
	CLOBGetMidpointPricesPost                      // POST Get midpoint prices (request body)
	CLOBGetSpread                                  // GET Get spread
	CLOBGetSpreads                                 // POST Get spreads
	CLOBGetLastTradePrice                          // GET Get last trade price
	CLOBGetLastTradesPricesGet                     // GET Get last trade prices (query parameters)
	CLOBGetLastTradesPricesPost                    // POST Get last trade prices (request body)
	CLOBGetPricesHistory                           // GET Get prices history
	CLOBGetBatchPricesHistory                      // POST Get batch prices history
	CLOBGetFeeRate                                 // GET Get fee rate (query parameter)
	CLOBGetFeeRateByPath                           // GET Get fee rate (path parameter)
	CLOBGetTickSize                                // GET Get tick size (query parameter)
	CLOBGetTickSizeByPath                          // GET Get tick size (path parameter)
	CLOBGetClobMarketInfo                          // GET Get CLOB market info (path condition_id)
	CLOBGetServerTime                              // GET Get server time

	// Trades
	CLOBGetTrades // GET Get trades

	// Balance & Allowance
	CLOBGetBalanceAllowance // GET Get balance and allowance

	// API Credentials（L1 EIP-712，见 https://docs.polymarket.com/cn/api-reference/authentication ）
	CLOBCreateApiKey // POST Create API key
	CLOBDeriveApiKey // GET Derive API key

	// Orders
	CLOBPostOrder          // POST Post a new order
	CLOBPostMultipleOrders // POST Post multiple orders
	CLOBCancelSingleOrder  // DELETE Cancel single order
	CLOBCancelMultiple     // DELETE Cancel multiple orders
	CLOBCancelAll          // DELETE Cancel all orders
	CLOBGetSingleOrderByID // GET Get single order by ID
	CLOBGetUserOrders      // GET Get user orders
	CLOBGetOrderScoring    // GET Get order scoring status
)

var CLOBAPITypeMap = map[CLOBAPIType]string{
	// Orderbook & Pricing
	CLOBGetOrderBook:            "/book",                        // GET Get order book
	CLOBGetOrderBooks:           "/books",                       // POST Get order books (request body)
	CLOBGetMarketPrice:          "/price",                       // GET Get market price
	CLOBGetMarketPricesGet:      "/prices",                      // GET Get market prices (query parameters)
	CLOBGetMarketPricesPost:     "/prices",                      // POST Get market prices (request body)
	CLOBGetMidpointPrice:        "/midpoint",                    // GET Get midpoint price
	CLOBGetMidpointPricesGet:    "/midpoints",                   // GET Get midpoint prices (query parameters)
	CLOBGetMidpointPricesPost:   "/midpoints",                   // POST Get midpoint prices (request body)
	CLOBGetSpread:               "/spread",                      // GET Get spread
	CLOBGetSpreads:              "/spreads",                     // POST Get spreads
	CLOBGetLastTradePrice:       "/last-trade-price",            // GET Get last trade price
	CLOBGetLastTradesPricesGet:  "/last-trades-prices",          // GET Get last trade prices (query parameters)
	CLOBGetLastTradesPricesPost: "/last-trades-prices",          // POST Get last trade prices (request body)
	CLOBGetPricesHistory:        "/prices-history",              // GET Get prices history
	CLOBGetBatchPricesHistory:   "/batch-prices-history",        // POST Get batch prices history
	CLOBGetFeeRate:              "/fee-rate",                    // GET Get fee rate (query parameter)
	CLOBGetFeeRateByPath:        "/fee-rate/{token_id}",         // GET Get fee rate (path parameter)
	CLOBGetTickSize:             "/tick-size",                   // GET Get tick size (query parameter)
	CLOBGetTickSizeByPath:       "/tick-size/{token_id}",        // GET Get tick size (path parameter)
	CLOBGetClobMarketInfo:       "/clob-markets/{condition_id}", // GET Get CLOB market info (path parameter)
	CLOBGetServerTime:           "/time",                        // GET Get server time

	// Trades
	CLOBGetTrades: "/trades", // GET Get trades

	// Balance & Allowance
	CLOBGetBalanceAllowance: "/balance-allowance", // GET Get balance and allowance

	// API Credentials（L1）
	CLOBCreateApiKey: "/auth/api-key",        // POST Create API key
	CLOBDeriveApiKey: "/auth/derive-api-key", // GET Derive API key

	// Orders
	CLOBPostOrder:          "/order",           // POST Post a new order
	CLOBPostMultipleOrders: "/orders",          // POST Post multiple orders
	CLOBCancelSingleOrder:  "/order",           // DELETE Cancel single order
	CLOBCancelMultiple:     "/orders",          // DELETE Cancel multiple orders
	CLOBCancelAll:          "/cancel-all",      // DELETE Cancel all orders
	CLOBGetSingleOrderByID: "/order/{orderID}", // GET Get single order by ID
	CLOBGetUserOrders:      "/data/orders",     // GET Get user orders
	CLOBGetOrderScoring:    "/order-scoring",   // GET Get order scoring status
}
