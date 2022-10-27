package system

import (
	"fmt"
	"gin-server/global"
	"gin-server/source/system"
	"go.uber.org/zap"
)

type InitDBService struct{}

func (i *InitDBService) InitMysqlData() error {
	return MysqlDataInitialize(
		system.User,
		system.Menu,
		system.Authority,
		system.AuthoritiesMenus,
		system.Casbin,
	)
}

const (
	Mysql           = "mysql"
	InitDataExist   = "[%v] --> %v 表的初始数据已存在"
	InitDataFailed  = "[%v] --> %v 表初始化数据失败"
	InitDataSuccess = "[%v] --> %v 表初始化数据成功"
)

type InitData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}

// MysqlDataInitialize 初始化表数据
func MysqlDataInitialize(inits ...InitData) error {
	for i := 0; i < len(inits); i++ {
		if inits[i].CheckDataExist() {
			global.Logger.Info(fmt.Sprintf(InitDataExist, Mysql, inits[i].TableName()))
			continue
		}
		if err := inits[i].Initialize(); err != nil {
			global.Logger.Error(fmt.Sprintf(InitDataFailed, Mysql, inits[i].TableName()), zap.Error(err))
		}
		global.Logger.Info(fmt.Sprintf(InitDataSuccess, Mysql, inits[i].TableName()))
	}
	return nil
}
