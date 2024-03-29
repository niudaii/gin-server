package main

import (
	"gin-server/global"
	"gin-server/initialize"
	"go.uber.org/zap"
	"log"
)

func main() {
	err := initialize.CommonConfig() // 初始化common配置文件
	if err != nil {
		log.Fatalf("解析 common 配置失败: %v", err)
	}
	err = initialize.ServerConfig() // 初始化server配置文件
	if err != nil {
		log.Fatalf("解析 server 配置失败: %v", err)
	}
	global.Common.Zap.Director += "/server"
	global.Logger = initialize.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.Logger)
	global.DB, err = initialize.Gorm() // gorm连接数据库
	if err != nil {
		global.Logger.Error("连接数据库失败", zap.Error(err))
		return
	}
	global.Logger.Info("连接数据库成功")
	err = initialize.RegisterTables(global.DB) // 初始化表和数据
	if err != nil {
		global.Logger.Error("初始化数据表和数据失败", zap.Error(err))
		return
	}
	global.Logger.Info("初始化数据表和数据成功")
	db, _ := global.DB.DB()
	defer db.Close()
	initialize.RunServer()
}
