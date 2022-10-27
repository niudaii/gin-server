package system

import "gin-server/service"

type ApiGroup struct {
	BaseApi
	UserApi
	OperationApi
}

var (
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	operationService = service.ServiceGroupApp.SystemServiceGroup.OperationService
)
