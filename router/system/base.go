package system

import (
	"gin-server/api"
	"gin-server/middleware"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (r *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	router := Router.Group("base").Use(middleware.Operation())
	routerWithoutRecord := Router.Group("base")
	baseApi := api.ApiGroupApp.SystemApiGroup.BaseApi
	{
		router.POST("login", baseApi.Login)
	}
	{
		routerWithoutRecord.GET("captcha", baseApi.Captcha)
		routerWithoutRecord.GET("logout", baseApi.Logout)
	}
}
