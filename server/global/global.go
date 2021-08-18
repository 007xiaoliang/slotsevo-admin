package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"slotsevo-admin/config"
	"slotsevo-admin/utils/timer"
)

var (
	SlotsDb     *gorm.DB
	SlotsRedis  *redis.Client
	SlotsConfig config.Server
	SlotsVp     *viper.Viper
	//SlotsLog    *oplogging.Logger
	SlotsLog   *zap.Logger
	SlotsTimer = timer.NewTimerTask()
)
