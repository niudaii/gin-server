package system

import (
	"fmt"
	"github.com/niudaii/gin-server/global"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

func (a *authoritiesMenus) TableName() string {
	return "authority_menus"
}

func (a *authoritiesMenus) Initialize() error {
	// 清空表
	if err := global.DB.Where("1 = 1").Delete(&AuthorityMenu{}).Error; err != nil {
		return fmt.Errorf("%v %v", err, a.TableName()+"表初始化失败")
	}
	// 初始化表数据
	entities := []AuthorityMenu{
		// 管理员权限
		{AuthorityId: "1", MenuId: 1},   //
		{AuthorityId: "1", MenuId: 10},  //
		{AuthorityId: "1", MenuId: 101}, //
		{AuthorityId: "1", MenuId: 102}, //
		{AuthorityId: "1", MenuId: 103}, //

		// 普通用户权限
		{AuthorityId: "2", MenuId: 1}, //
	}
	global.DB.Create(&entities)
	return nil
}

func (a *authoritiesMenus) CheckDataExist() bool {
	return false
}

type AuthorityMenu struct {
	AuthorityId string `gorm:"column:authority_authority_id"`
	MenuId      int    `gorm:"column:menu_menu_id"`
}
