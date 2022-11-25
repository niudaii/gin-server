package system

import (
	"gin-server/api/v1"
	"gin-server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("").Use(middleware.Operation())
	routerWithoutRecord := Router.Group("")
	userApi := v1.GroupApp.System.UserApi
	{
		router.POST("user/create", userApi.Create)                 // 创建用户
		router.POST("user/delete", userApi.Delete)                 // 删除用户
		router.POST("user/resetPassword", userApi.ResetPassword)   // 重置密码
		router.POST("user/changePassword", userApi.ChangePassword) // 修改密码
	}
	{
		routerWithoutRecord.GET("user/info", userApi.GetInfo) // 查询用户权限
		routerWithoutRecord.GET("user/menu", userApi.GetMenu) // 查询用户菜单
		routerWithoutRecord.POST("users", userApi.FindList)   // 分页查询用户列表
	}
}
