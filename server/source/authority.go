package source

import (
	"github.com/gookit/color"
	"slotsevo-admin/global"
	"slotsevo-admin/model"
	"time"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

var authorities = []model.SysAuthority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "1000", AuthorityName: "系统用户", ParentId: "0", DefaultRouter: "dashboard"},
}

//@author: xiaoliang
//@description: sys_authorities 表数据初始化
func (a *authority) Init() error {
	return global.SlotsDb.Transaction(func(tx *gorm.DB) error {
		if tx.First(&[]model.SysAuthority{}).RowsAffected > 0 {
			color.Danger.Println("\n[Mysql] --> sys_authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authorities 表初始数据成功!")
		return nil
	})
}
