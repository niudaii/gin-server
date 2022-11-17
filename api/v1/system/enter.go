package system

import "gin-server/service"

type ApiGroup struct {
	BaseApi
	UserApi
	OperationApi
}

var (
	userService      = service.GroupApp.System.UserService
	authorityService = service.GroupApp.System.AuthorityService
	operationService = service.GroupApp.System.OperationService
)
