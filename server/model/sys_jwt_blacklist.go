package model

import (
	"slotsevo-admin/global"
)

type JwtBlacklist struct {
	global.SlotsModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
