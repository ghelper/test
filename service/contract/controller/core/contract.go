package core

import (
	globalTypes "Project/types"
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/shopspring/decimal"
	"sync"
)

var (
	decimal0 = decimal.New(0, 0)
)

type Backend interface {
	//DB() *model.GalaxyDB
	//PrivateWalletServer() privatePB.PrivateWalletServiceClient
	//Cache() database.Cache
	//BaseWalletServer() basePB.BaseServiceClient
}
type Contract struct {
	backend Backend
	log     log.Logger
	mux     sync.RWMutex
}

func NewContract(backend Backend) (*Contract, error) {
	c := &Contract{
		backend: backend,
		log:     log.New("method", "contract"),
	}
	return c, nil
}

// 获取代币的余额
func (c *Contract) GetBalance(ctx context.Context, req *globalTypes.BalanceReq) (*globalTypes.BalanceRes, error) {
	return &globalTypes.BalanceRes{Balance: "156.69", Currency: "USDT"}, nil
}
