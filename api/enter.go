package api

import "gin-server/api/system"

type Group struct {
	SystemApiGroup system.ApiGroup
}

var GroupApp = new(Group)
