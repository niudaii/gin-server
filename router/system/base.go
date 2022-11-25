package system

import (
	"gin-server/api/v1"
	"gin-server/middleware"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (r *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	router := Router.Group("base").Use(middleware.Operation())
	routerWithoutRecord := Router.Group("base")
	baseApi := v1.GroupApp.System.BaseApi
	{
		router.POST("login", baseApi.Login)
	}
	{
		routerWithoutRecord.GET("captcha", baseApi.Captcha)
		routerWithoutRecord.GET("logout", baseApi.Logout)
	}
}
