package service

import "gin-server/service/system"

type Group struct {
	System system.ServiceGroup
}

var GroupApp = new(Group)
