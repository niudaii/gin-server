package system

import (
	"errors"
	"fmt"
	"gin-server/global"
	"gin-server/model/common/request"
	"gin-server/model/system"
	"gin-server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) Login(username, password string) (user system.User, err error) {
	err = global.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	return
}

func (s *UserService) Insert(req system.User) (err error) {
	var user system.User
	// 判断用户名是否注册
	if !errors.Is(global.DB.Where("username = ?", req.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("用户名已注册")
		return
	}
	// 否则附加uuid,密码md5简单加密 注册
	req.Password = utils.Md5(req.Password)
	req.UUID = uuid.NewV4()
	err = global.DB.Create(&req).Error
	return
}

func (s *UserService) Select(uuid uuid.UUID) (user system.User, err error) {
	err = global.DB.Preload("Authority").First(&user, "uuid = ?", uuid).Error
	return
}

func (s *UserService) SelectList(f *request.PageInfo) (list []system.User, total int64, err error) {
	db := global.DB.Model(&system.User{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// 分页
	if f.Page > 0 && f.PageSize > 0 {
		limit := f.PageSize
		offset := f.PageSize * (f.Page - 1)
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Preload("Authority").Find(&list).Error
	return
}

func (u *UserService) Delete(uuid uuid.UUID) (err error) {
	err = global.DB.Where("uuid = ?", uuid).Delete(&system.User{}).Error
	return
}

func (s *UserService) ResetPassword(uuid uuid.UUID) (err error) {
	err = global.DB.Model(&system.User{}).Where("uuid = ?", uuid).Update("password", utils.Md5("fgnb")).Error
	return
}

func (s *UserService) ChangePassword(uuid uuid.UUID, password string) (err error) {
	err = global.DB.Model(&system.User{}).Where("uuid = ?", uuid).Update("password", password).Error
	return
}
