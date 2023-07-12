package core

import (
	"NineTo5Server/core/internal"
	"NineTo5Server/global"
	"NineTo5Server/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.NineTo5_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.NineTo5_CONFIG.Zap.Director)
		_ = os.Mkdir(global.NineTo5_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.NineTo5_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
