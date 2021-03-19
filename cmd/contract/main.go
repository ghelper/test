package main

import (
	"Project/common/config"
	"fmt"
	//"galaxy-wallet/common/database"
	//"galaxy-wallet/common/database/redis"
	//"galaxy-wallet/common/model"
	//globalRedis "galaxy-wallet/common/model/redis"
	//privatePB "galaxy-wallet/common/proto/wallet/private_wallet"
	//"galaxy-wallet/services/middleware"
	//"galaxy-wallet/services/user"
	//"galaxy-wallet/services/user/server"
	//userServer "galaxy-wallet/services/user_service/controller/user"
	"Project/service/contract"
	"io"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type galaxyWallet struct {
	//db                  *model.GalaxyDB
	//cache               database.Cache
	//privateWalletServer privatePB.PrivateWalletServiceClient
	//middleware          *middleware.Middleware
	//userServer          user.Userer
}

//func (g *galaxyWallet) DB() *model.GalaxyDB {
//	return g.db
//}
//
//func (g *galaxyWallet) Cache() database.Cache {
//	return g.cache
//}
//
//func (g *galaxyWallet) PrivateWalletServer() privatePB.PrivateWalletServiceClient {
//	return g.privateWalletServer
//}
//
//func (g *galaxyWallet) Middleware() *middleware.Middleware {
//	return g.middleware
//}
//
//func (g *galaxyWallet) UserServer() user.Userer {
//	return g.userServer
//}

func main() {
	logFile, err := os.OpenFile(
		fmt.Sprintf("./logs/%d.log", time.Now().Unix()),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}
	log.Root().SetHandler(log.LvlFilterHandler(
		log.LvlDebug,
		log.StreamHandler(io.MultiWriter(os.Stdout, logFile), log.LogfmtFormat()),
	))
	// 通用配置
	if err = config.DefaultConfig(); err != nil {
		log.Error("default config", "err", err)
		return
	}
	// 初始化数据库
	//db, err := model.NewGalaxyDB(model.Config{
	//	Dialect:     viper.GetString("db.dialect"),
	//	Dsn:         viper.GetString("account_service.dsn"),
	//	Mode:        viper.GetString("server.dev"),
	//	MaxIdle:     viper.GetInt("db.max_idle"),
	//	MaxOpen:     viper.GetInt("db.max_open"),
	//	MaxLifetime: viper.GetInt("db.max_lifetime"),
	//})
	//if err != nil {
	//	log.Error("new db", "err", err)
	//	return
	//}
	//if err != nil {
	//	log.Error("dial private")
	//	return
	//}
	//conn, err := grpc.Dial(
	//	fmt.Sprintf(
	//		"%s:%d",
	//		viper.GetString("grpc.ip"),
	//		viper.GetInt64("grpc_port.private_wallet_service"),
	//	),
	//	grpc.WithInsecure())
	//if err != nil {
	//	log.Error("dial private")
	//	return
	//}
	//privateClient := privatePB.NewPrivateWalletServiceClient(conn)
	//globalRedis.NewRedisPool()
	//cache, err := redis.NewRedis(redis.Config{
	//	MaxIdle:   viper.GetInt("redis.max_idle"),
	//	MaxActive: viper.GetInt("redis.max_active"),
	//	Password:  viper.GetString("redis.password"),
	//	DataBase:  viper.GetInt("redis.data_base"),
	//	Url:       viper.GetString("redis.url"),
	//})
	//if err != nil {
	//	log.Error("open redis", "err", err)
	//	return
	//}
	g := gin.Default()
	g.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Length", "Content-Type", "AppKey", "AccessToken"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	group := g.Group("/api/v1/contract")
	backend := &galaxyWallet{
		//db:                  db,
		//cache:               cache,
		//privateWalletServer: privateClient,
	}
	//u := userServer.NewUser(db.DB, cache)
	//backend.middleware = middleware.NewMiddleware(backend)
	//backend.userServer = server.NewServer(backend)
	_, err = contract.NewServer(group, backend)
	if err != nil {
		log.Error("new contract server", "err", err)
		return
	}
	//log.Info("wallet contract server run")
	if err := g.Run(":8080"); err != nil {
		log.Error("server run", "err", err)
	}
}
