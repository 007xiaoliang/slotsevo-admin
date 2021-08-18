package source

import (
	"slotsevo-admin/global"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

var carbines = []gormadapter.CasbinRule{
	{Ptype: "p", V0: "1000", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "1000", V1: "/user/register", V2: "POST"},
	{Ptype: "p", V0: "1000", V1: "/menu/getMenu", V2: "POST"},
	{Ptype: "p", V0: "1000", V1: "/menu/getMenuList", V2: "POST"},
}

//@author: xiaoliang
//@description: casbin_rule 表数据初始化
func (c *casbin) Init() error {
	_ = global.SlotsDb.AutoMigrate(gormadapter.CasbinRule{})
	return global.SlotsDb.Transaction(func(tx *gorm.DB) error {
		if tx.First(&[]gormadapter.CasbinRule{}).RowsAffected > 0 {
			color.Danger.Println("\n[Mysql] --> casbin_rule 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> casbin_rule 表初始数据成功!")
		return nil
	})
}
