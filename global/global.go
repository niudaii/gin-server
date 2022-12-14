package global

import (
	"gin-server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Common config.Common
	Server config.Server
	Logger *zap.Logger
	DB     *gorm.DB
)
