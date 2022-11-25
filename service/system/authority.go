package system

import (
	"gin-server/global"
	"gin-server/model/system"
)

type AuthorityService struct{}

// GetMenus 获取角色菜单
func (s *AuthorityService) GetMenus(authority *system.Authority) (err error) {
	err = global.DB.Preload("Menus").First(authority).Error
	return
}
