package system

import (
	"gin-server/api"
	"github.com/gin-gonic/gin"
)

type OperationRouter struct{}

func (r *OperationRouter) InitOperationRouter(Router *gin.RouterGroup) {
	routerWithoutRecord := Router.Group("operation") // 不记录操作日志
	operationApi := api.GroupApp.SystemApiGroup.OperationApi
	{
		routerWithoutRecord.POST("find/list", operationApi.FindList) // 查询操作日志列表
	}
}
