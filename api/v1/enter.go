package v1

import "github.com/niudaii/gin-server/api/v1/system"

type Group struct {
	System system.ApiGroup
}

var GroupApp = new(Group)
