package system

type ServiceGroup struct {
	InitDBService
	UserService
	CasbinService
	AuthorityService
	OperationService
}
