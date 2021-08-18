package source

import (
	"github.com/gookit/color"
	"slotsevo-admin/global"
	"slotsevo-admin/model"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.SysUser{
	{SlotsModel: global.SlotsModel{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "xiaoliang", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "xiaoliang", HeaderImg: "http://qmplusimg.henrongyi.top/Slots_header.jpg", AuthorityId: "1000"},
}

//@author: xiaoliang
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.SlotsDb.Transaction(func(tx *gorm.DB) error {
		if tx.First(&[]model.SysAuthority{}).RowsAffected > 0 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
