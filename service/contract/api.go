package contract

import (
	"Project/service/contract/controller/core"
	contractError "Project/service/contract/errors"
	"Project/types"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	defaultTokenKey = "defaultTokenKey"
)

type backend interface {
	//wallet.Backend
	core.Backend
}

type Server struct {
	//contract *wallet.Contract_Wallet
	contract *core.Contract
	ginGroup *gin.RouterGroup
	log      log.Logger
	backend  backend
}

func NewServer(group *gin.RouterGroup, backend backend) (*Server, error) {
	//contract, err := wallet.NewContract_Wallet(backend)
	contract, err := core.NewContract(backend)
	if err != nil {
		return nil, err
	}
	s := &Server{
		contract: contract,
		ginGroup: group,
		log:      log.New("method", "contract server"),
		backend:  backend,
	}
	s.route()
	return s, nil
}

func (s *Server) route() {
	appGalaxyContract := s.ginGroup.Group("/") //middleware.SetGetAuthorization, s.backend.Middleware().CheckContractLogin
	{
		appGalaxyContract.GET("/balance", s.balance)
	}
}


type Response struct {
	Code int         `json:"errNo"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (s *Server) handleResponse(ctx *gin.Context, err error, obj interface{}) {
	var res Response
	res.Code, res.Msg = contractError.Error(err)
	if err != nil {
		s.log.Error("response", "url", ctx.Request.URL, "err", err)
		res.Data = obj
		ctx.Abort()
	} else {
		s.log.Info("response", "data", fmt.Sprintf("%+v", obj))
		res.Data = obj
	}
	ctx.JSON(http.StatusOK, &res)
}

func (s *Server) handleRequest(ctx *gin.Context, obj interface{}) error {
	var err error
	switch ctx.Request.Method {
	case http.MethodGet:
		err = ctx.BindQuery(obj)
	case http.MethodPost:
		err = ctx.ShouldBindBodyWith(obj, binding.JSON)
	}
	if err == nil {
		s.log.Debug("request", "obj", fmt.Sprintf("%+v", obj))
	}
	return err
}

func (s *Server) balance(ctx *gin.Context) {
	var req types.BalanceReq
	//err := s.handleRequest(ctx, &req)
	//if err != nil {
	//	s.handleResponse(ctx, err, nil)
	//	return
	//}
	assets, err := s.contract.GetBalance(ctx, &req)
	s.handleResponse(ctx, err, assets)
}
