package v1

import "gin-server/api/v1/system"

type Group struct {
	SystemApiGroup system.ApiGroup
}

var GroupApp = new(Group)
