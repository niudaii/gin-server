package initialize

import (
	"fmt"
	"gin-server/global"
	"gin-server/initialize/internal"
	"gin-server/model/system"
	"gin-server/service"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Gorm 初始化数据库
func Gorm() (db *gorm.DB, err error) {
	m := global.Server.Mysql
	if m.Dbname == "" {
		err = fmt.Errorf("db-name为空")
		return
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return
	}
}

// RegisterTables 初始化数据表和数据
func RegisterTables(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		// 系统模块表
		system.User{},
		system.Authority{},
		system.Menu{},
		adapter.CasbinRule{},
		system.Operation{},
	)
	if err != nil {
		return
	}
	// 表数据初始化
	err = service.ServiceGroupApp.SystemServiceGroup.InitDBService.InitMysqlData()
	return
}
