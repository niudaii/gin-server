package initialize

import (
	"github.com/niudaii/gin-server/global"
	"github.com/niudaii/gin-server/initialize/internal"
	"github.com/niudaii/gin-server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 初始化zap日志库
func Zap() (logger *zap.Logger) {
	if ok := utils.PathExists(global.Common.Zap.Director); !ok { // 判断是否有Director文件夹
		_ = os.Mkdir(global.Common.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	if global.Common.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
