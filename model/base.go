package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uint           `json:"-" gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      `json:"-"`                   // 创建时间
	UpdatedAt time.Time      `json:"-"`                   // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`      // 删除时间
}
