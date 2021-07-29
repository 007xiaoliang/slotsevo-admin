package global

import (
	"slotsevo-admin/utils/timer"

	"go.uber.org/zap"

	"slotsevo-admin/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GvaDb     *gorm.DB
	GvaRedis  *redis.Client
	GvaConfig config.Server
	GvaVp     *viper.Viper
	//GvaLog    *oplogging.Logger
	GvaLog   *zap.Logger
	GvaTimer timer.Timer = timer.NewTimerTask()
)
