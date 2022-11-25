package system

import (
	"fmt"
	"github.com/niudaii/gin-server/global"
	"github.com/niudaii/gin-server/model/system"
)

var Menu = new(menu)

type menu struct{}

func (m *menu) TableName() string {
	return "menus"
}

func (m *menu) Initialize() error {
	// 先清空
	if err := global.DB.Where("1 = 1").Delete(&system.Menu{}).Error; err != nil {
		return fmt.Errorf("%v %v", err, m.TableName()+"表初始化失败")
	}
	// 初始化表数据
	entities := []system.Menu{
		{MenuId: 1, Name: "demo", Path: "/", ParentId: 0, Meta: system.Meta{Title: "栗子", Show: true}, Component: "demo/Demo"},
		{MenuId: 10, Name: "system", ParentId: 0, Meta: system.Meta{Title: "系统设置", Show: true, HideChildren: true}, Component: "system/AdminIndex", Redirect: "/system/user"},
		{MenuId: 101, Name: "user", ParentId: 10, Meta: system.Meta{Title: "用户管理", Show: true}, Component: "system/User"},
		{MenuId: 102, Name: "operation", ParentId: 10, Meta: system.Meta{Title: "日志管理", Show: true}, Component: "system/Operation"},
	}
	if err := global.DB.Create(&entities).Error; err != nil {
		return fmt.Errorf("%v %v", err, m.TableName()+"表初始化失败")
	}
	return nil
}

func (m *menu) CheckDataExist() bool {
	return false
}
