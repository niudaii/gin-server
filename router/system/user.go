package system

import (
	"gin-server/api"
	"gin-server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Use(middleware.Operation())
	routerWithoutRecord := Router.Group("user") // 不记录操作日志
	userApi := api.GroupApp.SystemApiGroup.UserApi
	{
		router.POST("add", userApi.Add)                       // 新增用户
		router.POST("delete", userApi.Delete)                 // 删除用户
		router.POST("resetPassword", userApi.ResetPassword)   // 重置密码
		router.POST("changePassword", userApi.ChangePassword) // 修改密码
	}
	{
		routerWithoutRecord.POST("find/list", userApi.FindList) // 分页查询用户列表
		routerWithoutRecord.GET("info", userApi.UserInfo)       // 查询用户权限
		routerWithoutRecord.GET("menu", userApi.UserMenu)       // 查询用户菜单
	}
}
