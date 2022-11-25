package system

import (
	"github.com/niudaii/gin-server/global"
	"github.com/niudaii/gin-server/model/system"
)

type AuthorityService struct{}

// GetMenus 获取角色菜单
func (s *AuthorityService) GetMenus(authority *system.Authority) (err error) {
	err = global.DB.Preload("Menus").First(authority).Error
	return
}
