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
	TraceLog    *zap.Logger
	AccessLog   *zap.Logger
	BaseSendLog *zap.Logger
	ErrorLog    *zap.Logger
	SlotsTimer  = timer.NewTimerTask()
	GinMode     string
)
