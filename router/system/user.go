package system

import (
	"gin-server/api"
	"gin-server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.Operation())
	userRouterWithoutRecord := Router.Group("user") // 不记录操作日志
	userApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("add", userApi.Add)                       // 新增用户
		userRouter.POST("delete", userApi.Delete)                 // 删除用户
		userRouter.POST("resetPassword", userApi.ResetPassword)   // 重置密码
		userRouter.POST("changePassword", userApi.ChangePassword) // 修改密码
	}
	{
		userRouterWithoutRecord.POST("find/list", userApi.FindList) // 分页查询用户列表
		userRouterWithoutRecord.GET("info", userApi.UserInfo)       // 查询用户权限
		userRouterWithoutRecord.GET("menu", userApi.UserMenu)       // 查询用户菜单
	}
}
