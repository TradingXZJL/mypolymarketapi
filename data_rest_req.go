package mypolymarketapi

import (
	"strconv"
	"strings"
)

type DataGetLiveVolumeForEventAPI struct {
	client *DataRestClient
	req    *DataGetLiveVolumeForEventReq
}

type DataGetLiveVolumeForEventReq struct {
	Id *int64 `json:"id"` // integer true 事件 ID，建议传 >= 1
}

// int64 true 事件 ID，建议传 >= 1
func (api *DataGetLiveVolumeForEventAPI) Id(id int64) *DataGetLiveVolumeForEventAPI {
	api.req.Id = GetPointer(id)
	return api
}

type DataGetOpenInterestAPI struct {
	client *DataRestClient
	req    *DataGetOpenInterestReq
}

type DataGetOpenInterestReq struct {
	Market *[]string `json:"market"` // string[] true 市场标识列表（0x 开头 64 位十六进制哈希）
}

// []string true 市场标识列表（0x 开头 64 位十六进制哈希）
func (api *DataGetOpenInterestAPI) Market(market []string) *DataGetOpenInterestAPI {
	api.req.Market = GetPointer(market)
	return api
}

type DataGetCurrentPositionsAPI struct {
	client *DataRestClient
	req    *DataGetCurrentPositionsReq
}

// DataGetCurrentPositionsReq 对应 GET /positions 的 query（参数名与 OpenAPI 一致；`market` / `eventId` 为逗号分隔字符串）。
type DataGetCurrentPositionsReq struct {
	User          *string  `json:"user"`          // string 必填，用户地址（0x + 40 位 hex）
	Market        *string  `json:"market"`        // string 可选，condition id 列表，逗号分隔；与 eventId 互斥
	EventID       *string  `json:"eventId"`       // string 可选，事件 ID 列表，逗号分隔；与 market 互斥
	SizeThreshold *float64 `json:"sizeThreshold"` // number 可选，默认 1，最小 0
	Redeemable    *bool    `json:"redeemable"`    // boolean 可选，默认 false
	Mergeable     *bool    `json:"mergeable"`     // boolean 可选，默认 false
	Limit         *int64   `json:"limit"`         // integer 可选，默认 100，范围 0–500
	Offset        *int64   `json:"offset"`        // integer 可选，默认 0，范围 0–10000
	SortBy        *string  `json:"sortBy"`        // string 可选，默认 TOKENS：CURRENT、INITIAL、TOKENS、CASHPNL、PERCENTPNL、TITLE、RESOLVING、PRICE、AVGPRICE
	SortDirection *string  `json:"sortDirection"` // string 可选，默认 DESC：ASC、DESC
	Title         *string  `json:"title"`         // string 可选，标题筛选，最大长度 100
}

// string 必填，用户链上地址（proxy wallet 或 user address）
func (api *DataGetCurrentPositionsAPI) User(user string) *DataGetCurrentPositionsAPI {
	api.req.User = GetPointer(user)
	return api
}

// []string 可选，按 condition id 过滤（逗号拼接）；与 EventIDs 互斥
func (api *DataGetCurrentPositionsAPI) Markets(markets []string) *DataGetCurrentPositionsAPI {
	if len(markets) == 0 {
		api.req.Market = nil
		return api
	}
	api.req.Market = GetPointer(strings.Join(markets, ","))
	return api
}

// []int64 可选，按事件 ID 过滤（逗号拼接）；与 Markets 互斥
func (api *DataGetCurrentPositionsAPI) EventIDs(eventIDs []int64) *DataGetCurrentPositionsAPI {
	if len(eventIDs) == 0 {
		api.req.EventID = nil
		return api
	}
	parts := make([]string, len(eventIDs))
	for i, id := range eventIDs {
		parts[i] = strconv.FormatInt(id, BIT_BASE_10)
	}
	api.req.EventID = GetPointer(strings.Join(parts, ","))
	return api
}

// float64 可选，持仓数量阈值，默认 1，最小 0
func (api *DataGetCurrentPositionsAPI) SizeThreshold(sizeThreshold float64) *DataGetCurrentPositionsAPI {
	api.req.SizeThreshold = GetPointer(sizeThreshold)
	return api
}

// bool 可选，是否只返回可赎回持仓，默认 false
func (api *DataGetCurrentPositionsAPI) Redeemable(redeemable bool) *DataGetCurrentPositionsAPI {
	api.req.Redeemable = GetPointer(redeemable)
	return api
}

// bool 可选，是否只返回可合并持仓，默认 false
func (api *DataGetCurrentPositionsAPI) Mergeable(mergeable bool) *DataGetCurrentPositionsAPI {
	api.req.Mergeable = GetPointer(mergeable)
	return api
}

// int64 可选，每页条数，默认 100，最大 500
func (api *DataGetCurrentPositionsAPI) Limit(limit int64) *DataGetCurrentPositionsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// int64 可选，分页偏移，默认 0，最大 10000
func (api *DataGetCurrentPositionsAPI) Offset(offset int64) *DataGetCurrentPositionsAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// string 可选，排序字段：CURRENT、INITIAL、TOKENS、CASHPNL、PERCENTPNL、TITLE、RESOLVING、PRICE、AVGPRICE
func (api *DataGetCurrentPositionsAPI) SortBy(sortBy string) *DataGetCurrentPositionsAPI {
	api.req.SortBy = GetPointer(sortBy)
	return api
}

// string 可选，排序方向：ASC、DESC
func (api *DataGetCurrentPositionsAPI) SortDirection(sortDirection string) *DataGetCurrentPositionsAPI {
	api.req.SortDirection = GetPointer(sortDirection)
	return api
}

// string 可选，按标题子串筛选（最大 100 字符）
func (api *DataGetCurrentPositionsAPI) Title(title string) *DataGetCurrentPositionsAPI {
	api.req.Title = GetPointer(title)
	return api
}

type DataGetClosedPositionsAPI struct {
	client *DataRestClient
	req    *DataGetClosedPositionsReq
}

// DataGetClosedPositionsReq 对应 GET /closed-positions 的 query（参数名与 OpenAPI 一致；`market` / `eventId` 为逗号分隔字符串）。
type DataGetClosedPositionsReq struct {
	User          *string `json:"user"`          // string 必填，用户地址（0x + 40 位 hex）
	Market        *string `json:"market"`        // string 可选，condition id 列表，逗号分隔；与 eventId 互斥
	Title         *string `json:"title"`         // string 可选，按市场标题筛选，最大长度 100
	EventID       *string `json:"eventId"`       // string 可选，事件 ID 列表，逗号分隔；与 market 互斥
	Limit         *int64  `json:"limit"`         // integer 可选，默认 10，范围 0–50
	Offset        *int64  `json:"offset"`        // integer 可选，默认 0，范围 0–100000
	SortBy        *string `json:"sortBy"`        // string 可选，默认 REALIZEDPNL：REALIZEDPNL、TITLE、PRICE、AVGPRICE、TIMESTAMP
	SortDirection *string `json:"sortDirection"` // string 可选，默认 DESC：ASC、DESC
}

// string 必填，用户链上地址（proxy wallet 或 user address）
func (api *DataGetClosedPositionsAPI) User(user string) *DataGetClosedPositionsAPI {
	api.req.User = GetPointer(user)
	return api
}

// []string 可选，按 condition id 过滤（逗号拼接）；与 EventIDs 互斥
func (api *DataGetClosedPositionsAPI) Markets(markets []string) *DataGetClosedPositionsAPI {
	if len(markets) == 0 {
		api.req.Market = nil
		return api
	}
	api.req.Market = GetPointer(strings.Join(markets, ","))
	return api
}

// string 可选，按市场标题子串筛选（最大 100 字符）
func (api *DataGetClosedPositionsAPI) Title(title string) *DataGetClosedPositionsAPI {
	api.req.Title = GetPointer(title)
	return api
}

// []int64 可选，按事件 ID 过滤（逗号拼接）；与 Markets 互斥
func (api *DataGetClosedPositionsAPI) EventIDs(eventIDs []int64) *DataGetClosedPositionsAPI {
	if len(eventIDs) == 0 {
		api.req.EventID = nil
		return api
	}
	parts := make([]string, len(eventIDs))
	for i, id := range eventIDs {
		parts[i] = strconv.FormatInt(id, BIT_BASE_10)
	}
	api.req.EventID = GetPointer(strings.Join(parts, ","))
	return api
}

// int64 可选，返回条数上限，默认 10，最大 50
func (api *DataGetClosedPositionsAPI) Limit(limit int64) *DataGetClosedPositionsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// int64 可选，分页起始索引，默认 0，最大 100000
func (api *DataGetClosedPositionsAPI) Offset(offset int64) *DataGetClosedPositionsAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// string 可选，排序字段：REALIZEDPNL、TITLE、PRICE、AVGPRICE、TIMESTAMP
func (api *DataGetClosedPositionsAPI) SortBy(sortBy string) *DataGetClosedPositionsAPI {
	api.req.SortBy = GetPointer(sortBy)
	return api
}

// string 可选，排序方向：ASC、DESC
func (api *DataGetClosedPositionsAPI) SortDirection(sortDirection string) *DataGetClosedPositionsAPI {
	api.req.SortDirection = GetPointer(sortDirection)
	return api
}

type DataDownloadAccountingSnapshotAPI struct {
	client *DataRestClient
	req    *DataDownloadAccountingSnapshotReq
}

// DataDownloadAccountingSnapshotReq 对应 GET /v1/accounting/snapshot 的 query。
type DataDownloadAccountingSnapshotReq struct {
	User *string `json:"user"` // string 必填，用户地址（0x + 40 位 hex）
}

// string 必填，用户地址（0x + 40 位 hex） Polymarket Proxy Wallet Address
func (api *DataDownloadAccountingSnapshotAPI) User(user string) *DataDownloadAccountingSnapshotAPI {
	api.req.User = GetPointer(user)
	return api
}
