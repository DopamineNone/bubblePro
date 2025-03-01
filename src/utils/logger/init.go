package logger

import (
	config "github.com/DopamineNone/bubblePro/src/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init() {
	conf := config.GetConf().Log
	core := zapcore.NewCore(getEncoder())
	logger = zap.New(core)
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.TimeKey = "time"
	config.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeDuration = zapcore.SecondsDurationEncoder

	return zapcore.NewJSONEncoder(config)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	writer := &lumberjack.Logger{
		Filename:   filename,
		MaxAge:     maxAge,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
	}
	return zapcore.AddSync(writer)
}
