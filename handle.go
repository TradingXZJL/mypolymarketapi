package mypolymarketapi

import "fmt"

type PolyMarketErrorRes struct {
	Type  string `json:"type,omitempty"`
	Error string `json:"error,omitempty"`
}

func (err *PolyMarketErrorRes) handlerError() error {
	if err == nil || err.Error == "" {
		return nil
	}
	errStr := ""
	if err.Type != "" {
		errStr += "[" + err.Type + "] "
	}
	if err.Error != "" {
		errStr += err.Error
	}
	return fmt.Errorf("request error: %s", errStr)
}

type PolyMarketRestRes[T any] struct {
	PolyMarketErrorRes
	Data T `json:"data,omitempty"`
}

func handlerCommonRest[T any](body []byte) (*PolyMarketRestRes[T], error) {
	res := &PolyMarketRestRes[T]{}
	_ = json.Unmarshal(body, &res.PolyMarketErrorRes)
	if res.Error != "" {
		return res, nil
	}

	data := new(T)
	err := json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	res.Data = *data
	return res, nil
}
