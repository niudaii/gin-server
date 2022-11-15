package system

import "time"

type Operation struct {
	Id        uint      `json:"-" gorm:"primarykey"` // 主键ID
	CreatedAt time.Time `json:"createdAt"`           // 创建时间
	Operator  string    `json:"operator"`
	Ip        string    `json:"ip"`
	Agent     string    `json:"-"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Status    int       `json:"status"`
	Body      string    `json:"body" gorm:"type:text"`
	Resp      string    `json:"resp" gorm:"type:text"`
}
