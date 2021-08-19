package utils

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"runtime/debug"
	"slotsevo-admin/global"
	"strings"
	"time"
)

//@author: xiaoliang
//@function: ClearLog
//@description: 清理日志文件
//@param: keepDays 保留天数, logPath log文件路径
//@return: err
func ClearLog(keepDays int, logPath string) error {
	defer func() {
		if err := recover(); err != nil {
			global.ErrorLog.Error("[Recovery from panic]",
				zap.Any("error", err),
				zap.Any("stack", debug.Stack()),
			)
		}
	}()
	return walkDir(logPath, keepDays, remove)
}

// 获取指定路径下以及所有子目录下的所有日志文件，并且删除过期的日志
func walkDir(dir string, keepDays int, f func(string) error) (err error) {
	timeStr := time.Now().AddDate(0, 0, -keepDays).Format("2006-01-02")
	currentDate, _ := time.Parse("2006-01-02", timeStr)
	err = filepath.Walk(dir, func(fname string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			fileDate, _ := time.Parse("2006-01-02", strings.Split(strings.Split(fname, "_")[1], ".")[0])
			if !fileDate.After(currentDate) { // 日志日期在当前日期-keepDays之前则删除日志文件
				_ = f(fname)
			}
		}
		return nil
	})

	return err
}

func remove(file string) error {
	// 删除文件
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}
