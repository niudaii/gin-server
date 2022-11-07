package router

import "gin-server/router/system"

type Group struct {
	System system.RouterGroup
}

var GroupApp = new(Group)
