# mypolymarketapi

An unofficial Go SDK for the [Polymarket](https://polymarket.com) exchange, covering the CLOB REST API, Gamma REST API, Data REST API, and real-time Market WebSocket subscriptions.

---

## Installation

```
module: github.com/Hongssd/mypolymarketapi
go:     1.21+
```

```go
import "github.com/Hongssd/mypolymarketapi"
```

---

## API Endpoints

| Client | Host | Description |
|--------|------|-------------|
| `CLOBRestClient` | `clob.polymarket.com` | Order book, pricing, order placement & cancellation (auth required for trading) |
| `GammaRestClient` | `gamma-api.polymarket.com` | Market / event / tag discovery (public) |
| `DataRestClient` | `data-api.polymarket.com` | Positions, trades, leaderboard (public) |
| `MarketWsStreamClient` | `ws-subscriptions-clob.polymarket.com` | Real-time market data (WebSocket) |

---

## Global Configuration

```go
// Override the default logger (optional; defaults to logrus.New())
mypolymarketapi.SetLogger(myLogger)

// Override the HTTP request timeout (default: 100s)
mypolymarketapi.SetHttpTimeout(30 * time.Second)

// HTTP proxy for REST calls (must be called before any REST request)
mypolymarketapi.SetUseProxy(true, "http://your-proxy:port")

// HTTP proxy for WebSocket connections (must be called before OpenConn)
mypolymarketapi.SetWsUseProxy(true)
```

Multiple proxies are supported with weighted round-robin selection:

```go
mypolymarketapi.SetUseProxy(true,
    "http://proxy1:port",
    "http://proxy2:port",
    "http://proxy3:port",
)
```

---

## REST API

### Public Endpoints (no authentication required)

#### Get order book for a single asset

```go
p := mypolymarketapi.MyPolymarket{}

// A wallet is only required for authenticated (trading) endpoints.
// For public market data, pass a wallet obtained from NewWallet or leave
// authentication fields empty; the SDK will still send the request unsigned.
wallet, _ := mypolymarketapi.NewWallet(os.Getenv("PRIVATE_KEY"))
clobClient, err := p.NewCLOBRestClient(wallet, wallet.Signer.Hex(), 1)
if err != nil {
    log.Fatal(err)
}

res, err := clobClient.
    NewCLOBGetOrderBook().
    TokenID("<clob-token-id>").
    Do()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("bids: %d  asks: %d\n", len(res.Data.Bids), len(res.Data.Asks))
```

#### Get midpoint price

```go
res, err := clobClient.
    NewCLOBGetMidpointPrice().
    TokenID("<clob-token-id>").
    Do()
fmt.Println("midpoint:", res.Data.Mid)
```

#### Get last trade price

```go
res, err := clobClient.
    NewCLOBGetLastTradePrice().
    TokenID("<clob-token-id>").
    Do()
fmt.Println("last trade:", res.Data.Price)
```

#### Get spread

```go
res, err := clobClient.
    NewCLOBGetSpread().
    TokenID("<clob-token-id>").
    Do()
fmt.Println("spread:", res.Data.Spread)
```

#### Get price history

```go
res, err := clobClient.
    NewCLOBGetPricesHistory().
    Market("<condition-id>").
    Do()
for _, pt := range res.Data.History {
    fmt.Printf("t=%d p=%s\n", pt.T, pt.P)
}
```

#### Get multiple order books (batch)

```go
res, err := clobClient.
    NewCLOBGetOrderBooks().
    AddOrderBookReq(mypolymarketapi.CLOBGetOrderBookReq{
        TokenID: mypolymarketapi.GetPointer("<token-id-1>"),
    }).
    AddOrderBookReq(mypolymarketapi.CLOBGetOrderBookReq{
        TokenID: mypolymarketapi.GetPointer("<token-id-2>"),
    }).
    Do()
for _, ob := range res.Data {
    fmt.Printf("asset=%s bids=%d asks=%d\n", ob.AssetID, len(ob.Bids), len(ob.Asks))
}
```

---

### Authenticated Client Setup

Trading endpoints (order placement, cancellation, balance queries) require a two-layer signature: **L1 EIP-712** for API key derivation and **L2 HMAC-SHA256** per request.

```go
// Load wallet from private key (keep private keys in environment variables,
// never hard-code them in source files).
wallet, err := mypolymarketapi.NewWallet(os.Getenv("PRIVATE_KEY"))
if err != nil {
    log.Fatal(err)
}

// proxyAddress: the address used as the "maker" in orders.
//   EOA wallets  → use wallet.Signer.Hex()
//   POLY_PROXY   → use the proxy contract address
// nonce: deterministic seed for API key derivation; the same nonce always
//        produces the same API key credentials.
proxyAddress := wallet.Signer.Hex()
clobClient, err := p.NewCLOBRestClient(wallet, proxyAddress, 1)
if err != nil {
    log.Fatal(err)
}
```

On first use, register an API key with `CreateApiKey`. For subsequent sessions, `DeriveApiKey` (called internally by `NewCLOBRestClient`) restores the same credentials without a network round-trip to the registration endpoint.

```go
// One-time registration — store the returned credentials securely.
creds, err := wallet.CreateApiKey(1)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("APIKey=%s\nPassphrase=%s\n", creds.APIKey, creds.Passphrase)
```

---

### Place an Order

```go
import (
    "strconv"
    "time"
)

creds := clobClient.ApiKeyCreds()

order := mypolymarketapi.CLOBOrderReq{
    Maker:         mypolymarketapi.GetPointer(proxyAddress),
    Signer:        mypolymarketapi.GetPointer(wallet.Signer.Hex()),
    TokenID:       mypolymarketapi.GetPointer("<clob-token-id>"),
    MakerAmount:   mypolymarketapi.GetPointer("1000000"),  // USDC in 6-decimal units
    TakerAmount:   mypolymarketapi.GetPointer("10000000"), // outcome token units
    Side:          mypolymarketapi.GetPointer("BUY"),      // "BUY" or "SELL"
    Expiration:    mypolymarketapi.GetPointer("0"),        // 0 = no expiry
    Timestamp:     mypolymarketapi.GetPointer(strconv.FormatInt(time.Now().UnixMilli(), 10)),
    Salt:          mypolymarketapi.GetPointer(int64(1000000001)),
    SignatureType: mypolymarketapi.GetPointer(int64(0)), // 0=EOA, 1=POLY_PROXY, 2=GNOSIS_SAFE
}

res, err := clobClient.
    NewCLOBPostOrder().
    Order(order).
    Owner(creds.APIKey).
    OrderType("GTC"). // "GTC" | "FOK" | "GTD"
    DeferExec(false).
    Do()
if err != nil {
    log.Fatal(err)
}
fmt.Println("order id:", res.Data.OrderID)
```

---

### Cancel Orders

```go
// Cancel a single order
res, err := clobClient.
    NewCLOBCancelSingleOrder().
    OrderID("<order-id>").
    Do()

// Cancel multiple orders (up to 3,000 per request)
cancelReq := mypolymarketapi.CLOBCancelMultipleOrdersReq{
    {OrderID: mypolymarketapi.GetPointer("<order-id-1>")},
    {OrderID: mypolymarketapi.GetPointer("<order-id-2>")},
}
res, err = clobClient.NewCLOBCancelMultipleOrders().
    Req(cancelReq).
    Do()

// Cancel all open orders
res, err = clobClient.NewCLOBCancelAllOrders().Do()
```

---

### Query Orders & Balance

```go
// Get a specific order
res, err := clobClient.
    NewCLOBGetSingleOrderByID().
    OrderID("<order-id>").
    Do()

// Get all open orders for the authenticated user
res, err := clobClient.NewCLOBGetUserOrders().Do()

// Get USDC balance and allowance
res, err := clobClient.
    NewCLOBGetBalanceAllowance().
    AssetType("collateral").
    Do()
fmt.Printf("balance=%s  allowance=%s\n", res.Data.Balance, res.Data.Allowance)
```

---

### Gamma REST (market discovery, public)

```go
gammaClient := p.NewGammaRestClient()

// Paginated market list
res, err := gammaClient.
    NewGammaGetMarkets().
    Limit(20).
    Offset(0).
    Do()
for _, m := range res.Data {
    fmt.Printf("slug=%s  question=%.60s\n", m.Slug, m.Question)
}

// Fetch a specific market by condition ID
res, err = gammaClient.
    NewGammaGetMarkets().
    ConditionID("<condition-id>").
    Do()
```

---

### Data REST (positions & trade history, public)

```go
dataClient := p.NewDataRestClient()

// Positions for an address
res, err := dataClient.
    NewDataGetPositions().
    User("<wallet-address>").
    Do()
for _, pos := range res.Data {
    fmt.Printf("market=%s  size=%s\n", pos.Market, pos.Size)
}

// Trade history for an address
res, err = dataClient.
    NewDataGetTrades().
    User("<wallet-address>").
    Do()
```

---

## WebSocket Market Subscriptions

The `MarketWsStreamClient` implements a **generic broadcast-subscribe model** over a single WebSocket connection:

- Each subscriber receives its own **type-safe channel** — no `interface{}` casts required.
- Multiple subscribers share the same underlying connection and protocol subscription transparently.
- **Slow consumers drop messages** without affecting other subscribers or the dispatch goroutine.
- On disconnect, the client automatically reconnects with **exponential back-off** and re-sends all active protocol subscriptions.

### Connect

```go
p := mypolymarketapi.MyPolymarket{}
ws := p.NewMarketWsStreamClient()

// Optional — call before OpenConn
ws.SetLevel(2)                   // order book depth: 1 | 2 | 3  (default: 2)
ws.SetCustomFeatureEnabled(true) // required for best_bid_ask, new_market, market_resolved

if err := ws.OpenConn(); err != nil {
    log.Fatal(err)
}
defer ws.Close()
```

### Subscribe: OrderBook

```go
assetIDs := []string{"<token-id-1>", "<token-id-2>"}

obCh, err := ws.SubscribeOrderBook("strategy1", assetIDs, 512)
if err != nil {
    log.Fatal(err)
}

go func() {
    for ob := range obCh {
        fmt.Printf("book  asset=%s  bids=%d  asks=%d\n",
            ob.AssetID, len(ob.Bids), len(ob.Asks))
    }
    // Loop exits automatically when Unsubscribe is called.
}()
```

### Subscribe: Price Change

```go
pcCh, err := ws.SubscribePriceChange("strategy1", assetIDs, 1024)
if err != nil {
    log.Fatal(err)
}

go func() {
    for pc := range pcCh {
        fmt.Printf("price  asset=%s  side=%s  price=%s  size=%s\n",
            pc.AssetID, pc.Side, pc.Price, pc.Size)
    }
}()
```

### Subscribe: Last Trade Price

```go
ltCh, err := ws.SubscribeLastTradePrice("strategy2", assetIDs, 256)
if err != nil {
    log.Fatal(err)
}

go func() {
    for ltp := range ltCh {
        fmt.Printf("trade  asset=%s  price=%s\n", ltp.AssetID, ltp.Price)
    }
}()
```

### Subscribe: Best Bid/Ask

```go
// Requires: ws.SetCustomFeatureEnabled(true)
bbaCh, err := ws.SubscribeBestBidAsk("strategy1", assetIDs, 256)
if err != nil {
    log.Fatal(err)
}

go func() {
    for bba := range bbaCh {
        fmt.Printf("bba  asset=%s  bid=%s  ask=%s  spread=%s\n",
            bba.AssetID, bba.BestBid, bba.BestAsk, bba.Spread)
    }
}()
```

### Subscribe: Global Events

```go
// Requires: ws.SetCustomFeatureEnabled(true)

// Fired when a new market is created
nmCh, _ := ws.SubscribeNewMarket("monitor", 64)
go func() {
    for nm := range nmCh {
        fmt.Printf("new_market  id=%s  question=%s\n", nm.ID, nm.Question)
    }
}()

// Fired when a market is resolved
mrCh, _ := ws.SubscribeMarketResolved("monitor", 64)
go func() {
    for mr := range mrCh {
        fmt.Printf("resolved  id=%s  winner=%s\n", mr.ID, mr.WinningAssetID)
    }
}()
```

### Multiple Independent Subscribers

Multiple logical consumers can share the same WebSocket connection. Each holds its own channel and unsubscribes independently:

```go
// strategyA subscribes to order books
obCh, _ := ws.SubscribeOrderBook("strategyA", assetIDs, 512)

// strategyB subscribes to price changes — same connection, independent channel
pcCh, _ := ws.SubscribePriceChange("strategyB", assetIDs, 1024)

// Cancelling strategyA closes obCh but leaves strategyB unaffected.
ws.Unsubscribe("strategyA")
```

### Unsubscribe

```go
// Cancels all subscriptions registered under the given subscriber ID,
// closes all associated channels, and sends protocol-level unsubscribe
// messages for any assets with no remaining subscribers.
ws.Unsubscribe("strategy1")

// Consumer goroutines using `for ob := range obCh` exit automatically.
```

---

## Supported WebSocket Event Types

| Method | Event | Notes |
|--------|-------|-------|
| `SubscribeOrderBook` | `book` | Full snapshot on connect, incremental updates thereafter |
| `SubscribePriceChange` | `price_change` | Emitted on every quote update |
| `SubscribeLastTradePrice` | `last_trade_price` | Emitted on every matched trade |
| `SubscribeTickSizeChange` | `tick_size_change` | Low-frequency |
| `SubscribeBestBidAsk` | `best_bid_ask` | Requires `custom_feature_enabled` |
| `SubscribeNewMarket` | `new_market` | Global; low-frequency; requires `custom_feature_enabled` |
| `SubscribeMarketResolved` | `market_resolved` | Global; low-frequency; requires `custom_feature_enabled` |

---

## Error Handling

All methods return a `(result, error)` pair following standard Go conventions:

```go
res, err := client.NewXxx().Param("value").Do()
if err != nil {
    // Covers network errors, signing failures, non-2xx responses, etc.
    log.Error(err)
    return
}
// res.Data holds the typed response payload.
```

WebSocket channels are closed automatically when `Unsubscribe` is called, so consumer goroutines using `range` exit without additional bookkeeping:

```go
for ob := range obCh {
    // process ob
}
// Reaches here after Unsubscribe("strategy1") is called.
```

---

## Test Programs

| Path | Description |
|------|-------------|
| `tests/ws/market/main.go` | Full integration test for Market WebSocket (11 test modes) |
| `tests/clob/main.go` | CLOB REST smoke test: order placement & cancellation |
| `tests/data/main.go` | Data REST interface smoke test |
| `tests/gamma/main.go` | Gamma REST interface smoke test |

Run the WebSocket integration test (book mode, exits after 5 messages):

```sh
go run ./tests/ws/market/main.go -mode book -count 5 -timeout 30s
```

Full test parameter reference: [`tests/ws/market/GUIDE.md`](tests/ws/market/GUIDE.md)
