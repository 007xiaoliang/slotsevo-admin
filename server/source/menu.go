package source

import (
	"github.com/gookit/color"
	"slotsevo-admin/global"
	"slotsevo-admin/model"
	"time"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []model.SysBaseMenu{
	{SlotsModel: global.SlotsModel{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "后端数据", Icon: "s-data"}},
}

//@author: xiaoliang
//@description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return global.SlotsDb.Transaction(func(tx *gorm.DB) error {
		if tx.First(&[]model.SysBaseMenu{}).RowsAffected > 0 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
