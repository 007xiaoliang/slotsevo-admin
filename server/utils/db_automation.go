package utils

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

//@author: xiaoliang
//@function: ClearTable
//@description: 清理数据库表数据
//@param: target db(数据库对象) *gorm.DB,tableName(表名) string,compareField(比较字段) string , interval string  间隔
//@return: err

func ClearTable(db *gorm.DB, tableName string, compareField string, interval string) error {
	if db == nil {
		return errors.New("db Cannot be empty")
	}
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0")
	}
	return db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", tableName, compareField), time.Now().Add(-duration)).Error
}
