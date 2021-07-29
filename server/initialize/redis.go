package initialize

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"slotsevo-admin/global"
)

func Redis() {
	redisCfg := global.SlotsConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.SlotsLog.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.SlotsLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.SlotsRedis = client
	}
}
