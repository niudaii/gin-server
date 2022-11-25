package system

import (
	"gin-server/api/v1"
	"github.com/gin-gonic/gin"
)

type OperationRouter struct{}

func (r *OperationRouter) InitOperationRouter(Router *gin.RouterGroup) {
	routerWithoutRecord := Router.Group("")
	operationApi := v1.GroupApp.System.OperationApi
	{
		routerWithoutRecord.POST("operations", operationApi.FindList) // 查询操作日志列表
	}
}
