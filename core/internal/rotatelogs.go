package internal

import (
	"NineTo5Server/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
)

type lumberjackLogs struct{}

var LumberjackLogs = new(lumberjackLogs)

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (l *lumberjackLogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := &lumberjack.Logger{
		Filename:   path.Join(global.NineTo5_CONFIG.Zap.Director, level+".log"),
		MaxSize:    global.NineTo5_CONFIG.RotateLogs.MaxSize,
		MaxBackups: global.NineTo5_CONFIG.RotateLogs.MaxBackups,
		MaxAge:     global.NineTo5_CONFIG.RotateLogs.MaxAge,
		Compress:   global.NineTo5_CONFIG.RotateLogs.Compress,
	}

	if global.NineTo5_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
