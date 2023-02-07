package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

const (
	OK     = 200
	ERROR  = 400
	UNAUTH = 401
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(code, Response{
		code,
		data,
		msg,
	})
}

func Ok(data interface{}, msg string, c *gin.Context) {
	Result(OK, data, msg, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(OK, map[string]interface{}{}, msg, c)
}

func ErrorWithMessage(msg string, err error, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

func UnAuth(data interface{}, msg string, c *gin.Context) {
	Result(UNAUTH, data, msg, c)
}

func UnAuthWithMessage(msg string, c *gin.Context) {
	Result(UNAUTH, map[string]interface{}{}, msg, c)
}
