package middleware

import (
	"gin-server/model/common/response"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Access-Token")
		if token == "" {
			response.UnAuthWithMessage("未登录", c)
			c.Abort()
			return
		}
		j := utils.NewJwt()
		// 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			response.UnAuthWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
