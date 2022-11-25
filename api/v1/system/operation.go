package system

import (
	"github.com/niudaii/gin-server/model/common/response"
	"github.com/niudaii/gin-server/service/system"
	"github.com/gin-gonic/gin"
)

type OperationApi struct{}

// FindList 查询操作日志列表
func (a *OperationApi) FindList(c *gin.Context) {
	var req system.OperationFilter
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMessage("参数校验失败", err, c)
		return
	}
	if list, total, err := operationService.SelectList(&req); err != nil {
		response.ErrorWithMessage("查询操作日志列表失败", err, c)
	} else {
		response.Ok(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "查询操作日志列表成功", c)
	}
}
