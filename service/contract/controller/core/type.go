package core

import (
	"encoding/json"
)

type WithdrawReq struct {
	AccountId     uint64                  // 用户id
	Symbol        string                  // 交易对
	Level         uint64                  // 杠杆倍数
	OrderId       uint64                  // 订单号
	HistoryId     uint64                  // 记录号
	OpenPrice     string                  // 开仓单价
	BlastingPrice string                  // 预计爆仓价
	OpenTime      int64                   // 开仓时间
	AssetCode     string                  // 资产名称
}

type withdrawReqJSON struct {
	AccountId     uint64 `json:"account_id"`     // 用户id
	Symbol        string `json:"symbol"`         // 交易对
	Level         uint64 `json:"level"`          // 杠杆倍数
	OrderId       uint64 `json:"order_id"`       // 订单号
	HistoryId     uint64 `json:"history_id"`     // 记录号
	OpenPrice     string `json:"open_price"`     // 开仓单价
	BlastingPrice string `json:"blasting_price"` // 预计爆仓价
	OpenTime      int64  `json:"open_time"`      // 开仓时间
	AssetCode     string `json:"asset_code"`     // 资产名称
	OrderType     string `json:"order_type"`     // 买卖类型(买涨买跌)
	Amount        string `json:"amount"`         // 资产金额
	Mode          string `json:"mode"`           // 合约模式(逐仓/全仓)
}

func (w *WithdrawReq) UnmarshalJSON(b []byte) error {
	var req withdrawReqJSON
	if err := json.Unmarshal(b, &req); err != nil {
		return err
	}
	*w = WithdrawReq{
		AccountId:     req.AccountId,
		Symbol:        req.Symbol,
		Level:         req.Level,
		OrderId:       req.OrderId,
		HistoryId:     req.HistoryId,
		OpenPrice:     req.OpenPrice,
		BlastingPrice: req.BlastingPrice,
		OpenTime:      req.OpenTime,
		AssetCode:     req.AssetCode,
	}
	return nil
}

