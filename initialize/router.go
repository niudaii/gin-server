package initialize

import (
	"github.com/niudaii/gin-server/global"
	"github.com/niudaii/gin-server/middleware"
	"github.com/niudaii/gin-server/router"
	"github.com/gin-gonic/gin"
)

// Routers 注册路由
func Routers() *gin.Engine {
	r := gin.Default()
	//静态资源
	//r.Static("/export", "export")
	// 跨域
	r.Use(middleware.Cors())
	global.Logger.Info("注册 CORS handler")
	// 获取路由组实例
	g := r.Group("api")
	systemRouter := router.GroupApp.System
	PublicGroup := g.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		// 注册基础功能路由 不做鉴权
		systemRouter.InitBaseRouter(PublicGroup)
	}
	// 身份验证+权限控制
	PrivateGroup := g.Group("")
	PrivateGroup.Use(middleware.JwtAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup)      // 注册 user 相关路由
		systemRouter.InitOperationRouter(PrivateGroup) // 注册 operation 相关路由
	}
	return r
}
