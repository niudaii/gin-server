package initialize

import (
	"gin-server/global"
	"gin-server/middleware"
	"gin-server/router"
	"github.com/gin-gonic/gin"
)

// Routers 注册路由
func Routers() *gin.Engine {
	r := gin.Default()
	//静态资源
	//r.Static("/export", "export")
	//global.Logger.Info("开启静态资源目录 /export")
	// 跨域
	r.Use(middleware.Cors())
	global.Logger.Info("注册 CORS handler")
	// 获取路由组实例
	g := r.Group("api")
	systemRouter := router.GroupApp.System
	// 不做鉴权
	publicGroup := g.Group("")
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		// 注册基础功能路由
		systemRouter.InitBaseRouter(publicGroup)
	}
	// 身份验证+权限控制
	privateGroup := g.Group("")
	privateGroup.Use(middleware.JwtAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(privateGroup)      // 注册 user 相关路由
		systemRouter.InitOperationRouter(privateGroup) // 注册 operation 相关路由
	}
	return r
}
