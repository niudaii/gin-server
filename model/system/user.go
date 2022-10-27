package system

import (
	"gin-server/model"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	model.BaseModel
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"username" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	AuthorityId string    `json:"authorityId" binding:"required"` // 权限 Id
	Authority   Authority `json:"authority"`
}
