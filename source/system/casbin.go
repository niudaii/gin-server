package system

import (
	"fmt"
	"gin-server/global"
	adapter "github.com/casbin/gorm-adapter/v3"
)

var Casbin = new(casbin)

type casbin struct{}

func (c *casbin) TableName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (c *casbin) Initialize() error {
	// 先清空表
	if err := global.DB.Where("1 = 1").Delete(&adapter.CasbinRule{}).Error; err != nil {
		return fmt.Errorf("%v %v", err, c.TableName()+"表初始化失败")
	}
	// 初始化表数据
	entities := []adapter.CasbinRule{
		// admin 权限
		{Ptype: "p", V0: "1", V1: "/user/info", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/user/menu", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/add", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/delete", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/resetPassword", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/find/list", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/operation/find/list", V2: "POST"},

		// user 权限
		{Ptype: "p", V0: "2", V1: "/user/info", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/user/menu", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/user/changePassword", V2: "POST"},
	}
	if err := global.DB.Create(&entities).Error; err != nil {
		return fmt.Errorf("%v %v", err, c.TableName()+"表初始化失败")
	}

	return nil
}

func (c *casbin) CheckDataExist() bool {
	return false
}
