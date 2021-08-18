package source

import (
	"github.com/gookit/color"
	"gorm.io/gorm"
	"slotsevo-admin/global"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

var authorityMenus = []AuthorityMenus{
	{"1000", 1},
}

//@author
//@description: sys_authority_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return global.SlotsDb.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.First(&[]AuthorityMenus{}).RowsAffected > 0 {
			color.Danger.Println("\n[Mysql] --> sys_authority_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authority_menus 表初始数据成功!")
		return nil
	})
}
