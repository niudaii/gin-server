package v1

import "gin-server/api/v1/system"

type Group struct {
	System system.ApiGroup
}

var GroupApp = new(Group)
