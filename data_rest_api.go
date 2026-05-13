package mypolymarketapi

import (
	"errors"
)

// GET Get live volume for an event
func (c *DataRestClient) NewDataGetLiveVolumeForEvent() *DataGetLiveVolumeForEventAPI {
	return &DataGetLiveVolumeForEventAPI{
		client: c,
		req:    &DataGetLiveVolumeForEventReq{},
	}
}

func (api *DataGetLiveVolumeForEventAPI) Do() (*PolyMarketRestRes[DataGetLiveVolumeForEventRes], error) {
	if api.req.Id == nil {
		return nil, errors.New("id is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, DATA_REST, api.req, DataAPITypeMap[DataGetLiveVolumeForEvent])
	return pmCallAPI[DataGetLiveVolumeForEventRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get open interest
func (c *DataRestClient) NewDataGetOpenInterest() *DataGetOpenInterestAPI {
	return &DataGetOpenInterestAPI{
		client: c,
		req:    &DataGetOpenInterestReq{},
	}
}

func (api *DataGetOpenInterestAPI) Do() (*PolyMarketRestRes[DataGetOpenInterestRes], error) {
	url := pmHandlerRequestAPIWithPathQueryParam(REST, DATA_REST, api.req, DataAPITypeMap[DataGetOpenInterest])
	return pmCallAPI[DataGetOpenInterestRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get current positions for a user（Data API：`GET /positions`；必填 query `user` 为用户地址，须符合 `^0x[a-fA-F0-9]{40}$`。`market` 与 `eventId` 互斥，均为逗号分隔列表。）
func (c *DataRestClient) NewDataGetCurrentPositions() *DataGetCurrentPositionsAPI {
	return &DataGetCurrentPositionsAPI{
		client: c,
		req:    &DataGetCurrentPositionsReq{},
	}
}

func (api *DataGetCurrentPositionsAPI) Do() (*PolyMarketRestRes[DataGetCurrentPositionsRes], error) {
	if api.req.User == nil {
		return nil, errors.New("user is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, DATA_REST, api.req, DataAPITypeMap[DataGetCurrentPositions])
	return pmCallAPI[DataGetCurrentPositionsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get closed positions for a user（Data API：`GET /closed-positions`；必填 query `user`；`market` 与 `eventId` 互斥，均为逗号分隔。）
func (c *DataRestClient) NewDataGetClosedPositions() *DataGetClosedPositionsAPI {
	return &DataGetClosedPositionsAPI{
		client: c,
		req:    &DataGetClosedPositionsReq{},
	}
}

func (api *DataGetClosedPositionsAPI) Do() (*PolyMarketRestRes[DataGetClosedPositionsRes], error) {
	if api.req.User == nil {
		return nil, errors.New("user is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, DATA_REST, api.req, DataAPITypeMap[DataGetClosedPositions])
	return pmCallAPI[DataGetClosedPositionsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Download an accounting snapshot (ZIP of CSVs)
// 文档来源：Polymarket Data API `GET /v1/accounting/snapshot`，query 必填 `user`，返回 application/zip（包含 positions.csv 与 equity.csv）。
func (c *DataRestClient) NewDataDownloadAccountingSnapshot() *DataDownloadAccountingSnapshotAPI {
	return &DataDownloadAccountingSnapshotAPI{
		client: c,
		req:    &DataDownloadAccountingSnapshotReq{},
	}
}

func (api *DataDownloadAccountingSnapshotAPI) Do() (*PolyMarketRestRes[DataDownloadAccountingSnapshotRes], error) {
	if api.req.User == nil {
		return nil, errors.New("user is required")
	}

	url := pmHandlerRequestAPIWithPathQueryParam(REST, DATA_REST, api.req, DataAPITypeMap[DataDownloadAccountingSnapshot])
	body, err := Request(url.String(), NIL_REQBODY, GET, IS_GZIP)
	if err != nil {
		return nil, err
	}

	data, err := parseAccountingSnapshotZip(body)
	if err != nil {
		return nil, err
	}

	res := &PolyMarketRestRes[DataDownloadAccountingSnapshotRes]{
		Data: *data,
	}
	return res, nil
}
