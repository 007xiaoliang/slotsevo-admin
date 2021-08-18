package service

import (
	"errors"
	"gorm.io/gorm"
	"slotsevo-admin/global"
	"slotsevo-admin/model"
	"time"
)

//@author: xiaoliang
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.SlotsDb.Create(&jwtList).Error
	return
}

//@author: xiaoliang
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.SlotsDb.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@author: xiaoliang
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.SlotsRedis.Get(userName).Result()
	return err, redisJWT
}

//@author: xiaoliang
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: userName string
//@return: err error, redisJWT string

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.SlotsConfig.JWT.ExpiresTime) * time.Second
	err = global.SlotsRedis.Set(userName, jwt, timer).Err()
	return err
}
