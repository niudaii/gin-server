package internal

import (
	"gin-server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config 自定义 gorm 配置
func (g *_gorm) Config() *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	}
	_default := logger.Default
	switch global.Server.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
