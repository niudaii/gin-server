package system

import (
	"gin-server/global"
	"gin-server/model/common/request"
	"gin-server/model/system"
	"gorm.io/gorm"
)

type OperationService struct{}

type OperationFilter struct {
	system.Operation
	request.PageInfo
}

func (f *OperationFilter) conditions() (db *gorm.DB) {
	db = global.DB.Model(&system.Operation{})
	if f.Operator != "" {
		db = db.Where("operator = ?", f.Operator)
	}
	if f.Path != "" {
		db = db.Where("path = ?", f.Path)
	}
	return
}

func (s *OperationService) Insert(op system.Operation) (err error) {
	err = global.DB.Create(&op).Error
	return
}

func (s *OperationService) SelectList(f *OperationFilter) (list []system.Operation, total int64, err error) {
	db := f.conditions()
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
	err = db.Order("id desc").Find(&list).Error
	return
}
