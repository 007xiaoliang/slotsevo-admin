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
	SlotsDb     *gorm.DB
	SlotsRedis  *redis.Client
	SlotsConfig config.Server
	SlotsVp     *viper.Viper
	//SlotsLog    *oplogging.Logger
	SlotsLog   *zap.Logger
	SlotsTimer timer.Timer = timer.NewTimerTask()
)
