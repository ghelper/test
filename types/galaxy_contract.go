package types

type Assets struct {
	Assets []BalanceRes `json:"assets"`
}

type BalanceRes struct {
	Currency string `json:"currency"` // 交易币种
	Balance  string `json:"balance"`  // 余额
}

type BalanceReq struct {
	Currency  string `json:"currency"`   // 交易币种
	AccountId uint64 `json:"account_id"` // 账户id
}

type TokenReq struct {
	AssetCode string `json:"AssetCode"`
	AccountId uint64 `json:"AccountId"`
	Amount    uint64 `json:"Amount"`
	Type      uint32 `json:"type"`
	OrderId   uint64 `json:"OrderId"`
}

type DepositTokenReq struct {
	AssetCode      string `json:"AssetCode"`
	AccountId      uint64 `json:"AccountId"`
	Amount         string `json:"Amount"`
	CloseType      uint8  `json:"CloseType"`
	OrderId        uint64 `json:"OrderId"`
	ClosePrice     string `json:"ClosePrice"`    // 平仓价格
	EarningAmount  string `json:"EarningAmount"` // 收益
	UserEarning    string // 用户返回金额,本金 + 收益
	FeeAmount      string `json:"FeeAmount"`      // 手续费
	BlastingAmount string `json:"BlastingAmount"` // 爆仓费
	CloseTime      int64  `json:"CloseTime"`      // 平仓时间
	FundCharge     string `json:"FundCharge"`     // 资金费用
}

type WithdrawTokenReq struct {
	AccountId     uint64 `json:"AccountId"`
	Symbol        string `json:"Symbol"` // 交易对
	Level         uint64 `json:"Level"`  // 杠杆倍数
	OrderId       uint64 `json:"OrderId"`
	AssetCode     string `json:"AssetCode"`
	Amount        string `json:"Amount"`
	OrderType     string `json:"OrderType"` //买卖类型(买涨买跌)
	OpenPrice     string `json:"OpenPrice"`
	BlastingPrice string `json:"BlastingPrice"`
	OpenTime      int64  `json:"OpenTime"`
	Mode          string `json:"mode"` // 合约模式(逐仓/全仓)
}

type UserInfoRes struct {
	AccountId uint64 `json:"account_id"`
}
