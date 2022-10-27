package middleware

import (
	"gin-server/model/common/response"
	"gin-server/service"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := service.ServiceGroupApp.SystemServiceGroup.CasbinService.Casbin()
		//判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if utils.RunMode == utils.DebugMode || success {
			c.Next()
		} else {
			response.OkWithMessage("权限不足", c)
			c.Abort()
			return
		}
	}
}
