package system

import (
	"errors"
	"fmt"
	"gin-server/global"
	"gin-server/model/system"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var User = new(user)

type user struct{}

func (u *user) TableName() string {
	return "users"
}

func (u *user) Initialize() error {
	entities := []system.User{
		{UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", AuthorityId: "1"},
		{UUID: uuid.NewV4(), Username: "user", Password: "e10adc3949ba59abbe56e057f20f883e", AuthorityId: "2"},
	}
	if err := global.DB.Create(&entities).Error; err != nil {
		return fmt.Errorf("%v %v", err, u.TableName()+"表初始化失败")
	}
	return nil
}

func (u *user) CheckDataExist() bool {
	return !errors.Is(global.DB.Where("username = ?", "user").First(&system.User{}).Error, gorm.ErrRecordNotFound)
}
