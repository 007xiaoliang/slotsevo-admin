package service

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"slotsevo-admin/config"
	"slotsevo-admin/global"
	"slotsevo-admin/model"
	"slotsevo-admin/utils"
)

//@author: xiaoliang
//@function: writeConfig
//@description: 回写配置
//@param:
//@return: error
func writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.SlotsConfig.Mysql = mysql
	cs := utils.StructToMap(global.SlotsConfig)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

//@author: xiaoliang
//@function: createTable
//@description: 创建数据库(mysql)
//@param: dsn string, driver string, createSql
//@return: error
func createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func initDB(InitDBFunctions ...model.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

//@author: xiaoliang
//@function: InitDB
//@description: 创建数据库并初始化
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu
func InitDB(gormConfig *gorm.Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	conf := global.SlotsConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", conf.Username, conf.Password, conf.Path)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.Dbname)
	if err = createTable(dsn, "mysql", createSql); err != nil {
		return db, err
	}

	MysqlConfig := config.Mysql{
		Path:     conf.Path,
		Dbname:   conf.Dbname,
		Username: conf.Username,
		Password: conf.Password,
		Config:   conf.Config,
	}

	if err = writeConfig(global.SlotsVp, MysqlConfig); err != nil {
		return db, err
	}

	linkDns := conf.Username + ":" + conf.Password + "@tcp(" + conf.Path + ")/" + conf.Dbname + "?" + conf.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		_ = writeConfig(global.SlotsVp, MysqlConfig)
		return db, nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		global.SlotsDb = db
	}
	// 注释数据库初始化，直接改用sql语句操作数据库初始化数据库
	//err = global.SlotsDb.AutoMigrate(
	//	model.SysUser{},
	//	model.SysAuthority{},
	//	model.SysApi{},
	//	model.SysBaseMenu{},
	//	model.SysBaseMenuParameter{},
	//	model.JwtBlacklist{},
	//)
	//if err != nil {
	//	_ = writeConfig(global.SlotsVp, MysqlConfig)
	//	return db, err
	//}
	//err = initDB(
	//	source.Admin,
	//	source.Api,
	//	source.Authority,
	//	source.AuthoritiesMenus,
	//	source.Casbin,
	//	source.DataAuthorities,
	//	source.BaseMenu)
	//if err != nil {
	//	_ = writeConfig(global.SlotsVp, MysqlConfig)
	//	return db, err
	//}
	return db, nil
}
