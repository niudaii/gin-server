package middleware

import (
	"bytes"
	"gin-server/global"
	"gin-server/model/system"
	"gin-server/service"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

// Operation 记录用户请求
func Operation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err == nil {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		operation := system.Operation{
			Status: c.Writer.Status(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Query:  c.Request.URL.RawQuery,
			Body:   string(body),
			Ip:     c.ClientIP(),
			Agent:  c.Request.UserAgent(),
		}
		claims, _ := utils.GetClaims(c)
		if claims != nil {
			operation.Operator = claims.Username
		}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		// 先放行获取返回包
		c.Next()
		operation.Status = c.Writer.Status()
		operation.Resp = writer.body.String()
		operation.Resp = operation.Resp[:utils.Min(1000, len(operation.Resp))]
		if err := service.GroupApp.System.OperationService.Insert(operation); err != nil {
			global.Logger.Error("插入操作记录失败", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
