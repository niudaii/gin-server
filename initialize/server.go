package initialize

import (
	"fmt"
	"gin-server/global"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

// initServer gin endless 热重启(linux、mac)
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

// RunServer 启动服务
func RunServer() {
	gin.SetMode(global.Server.RunMode)
	router := Routers()
	address := fmt.Sprintf(":%v", global.Server.Port)
	s := initServer(address, router)
	global.Logger.Info("启动 WebServer 成功", zap.String("address", address))
	global.Logger.Error(s.ListenAndServe().Error())
}
