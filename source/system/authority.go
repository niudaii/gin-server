package system

import (
	"fmt"
	"github.com/niudaii/gin-server/global"
	"github.com/niudaii/gin-server/model/system"
)

var Authority = new(authority)

type authority struct{}

func (a *authority) TableName() string {
	return "authorities"
}

func (a *authority) Initialize() error {
	// 先清空
	if err := global.DB.Where("1 = 1").Delete(&system.Authority{}).Error; err != nil {
		return fmt.Errorf("%v %v", err, a.TableName()+"表初始化失败")
	}
	// 初始化
	entities := []system.Authority{
		{AuthorityId: "1", AuthorityName: "admin"},
		{AuthorityId: "2", AuthorityName: "user"},
	}
	if err := global.DB.Create(&entities).Error; err != nil {
		return fmt.Errorf("%v %v", err, a.TableName()+"表初始化失败")
	}
	return nil
}

func (a *authority) CheckDataExist() bool {
	return false
}
