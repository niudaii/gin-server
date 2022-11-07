package service

import "gin-server/service/system"

type Group struct {
	SystemServiceGroup system.ServiceGroup
}

var GroupApp = new(Group)
