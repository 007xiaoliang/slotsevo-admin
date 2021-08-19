package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"slotsevo-admin/global"
	"slotsevo-admin/initialize/internal"
	"slotsevo-admin/service"
)

//@author: xiaoliang
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB
func Gorm() *gorm.DB {
	switch global.SlotsConfig.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

//
//@author: xiaoliang
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB
func GormMysql() *gorm.DB {
	m := global.SlotsConfig.Mysql
	if m.Dbname == "" {
		return nil
	}
	db, err := service.InitDB(gormConfig(m.LogMode))
	if err != nil {
		global.TraceLog.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
	}
	return db
}

//@author: xiaoliang
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config
func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.SlotsConfig.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = internal.Default.LogMode(logger.Info)
			break
		}
		config.Logger = internal.Default.LogMode(logger.Silent)
	}
	return config
}
