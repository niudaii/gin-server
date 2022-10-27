package system

import (
	"gin-server/api"
	"gin-server/middleware"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (r *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base").Use(middleware.Operation())
	baseRouterWithoutRecord := Router.Group("base")
	baseApi := api.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
	}
	{
		baseRouterWithoutRecord.GET("captcha", baseApi.Captcha)
		baseRouterWithoutRecord.GET("logout", baseApi.Logout)
	}
	return baseRouter
}
