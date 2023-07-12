package global

import (
	"NineTo5Server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	NineTo5_VP     *viper.Viper
	NineTo5_CONFIG config.Server
	NineTo5_LOG    *zap.Logger
	NineTo5_DB     *gorm.DB
)
