package system

import "gin-server/service"

type ApiGroup struct {
	BaseApi
	UserApi
	OperationApi
}

var (
	userService      = service.GroupApp.SystemServiceGroup.UserService
	authorityService = service.GroupApp.SystemServiceGroup.AuthorityService
	operationService = service.GroupApp.SystemServiceGroup.OperationService
)
