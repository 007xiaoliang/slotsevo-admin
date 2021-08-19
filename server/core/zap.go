package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"slotsevo-admin/global"
	"slotsevo-admin/utils"
	"time"
)

var level zapcore.Level

func Zap() (traceLogger *zap.Logger, accessLogger *zap.Logger, baseSendLogger *zap.Logger, errorLogger *zap.Logger) {
	if ok, _ := utils.PathExists(global.SlotsConfig.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.SlotsConfig.Zap.Director)
		_ = os.Mkdir(global.SlotsConfig.Zap.Director, os.ModePerm)
	}

	switch global.SlotsConfig.Zap.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		traceLogger = zap.New(getEncoderCore("trace"), zap.AddStacktrace(level))
		accessLogger = zap.New(getEncoderCore("access"), zap.AddStacktrace(level))
		baseSendLogger = zap.New(getEncoderCore("baseSend"), zap.AddStacktrace(level))
		errorLogger = zap.New(getEncoderCore("error"), zap.AddStacktrace(level))
	} else {
		traceLogger = zap.New(getEncoderCore("trace"))
		accessLogger = zap.New(getEncoderCore("access"))
		baseSendLogger = zap.New(getEncoderCore("baseSend"))
		errorLogger = zap.New(getEncoderCore("error"))
	}
	if global.SlotsConfig.Zap.ShowLine {
		traceLogger = traceLogger.WithOptions(zap.AddCaller())
		accessLogger = accessLogger.WithOptions(zap.AddCaller())
		baseSendLogger = baseSendLogger.WithOptions(zap.AddCaller())
		errorLogger = baseSendLogger.WithOptions(zap.AddCaller())
	}
	return
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.SlotsConfig.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.SlotsConfig.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.SlotsConfig.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.SlotsConfig.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.SlotsConfig.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.SlotsConfig.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(logName string) (core zapcore.Core) {
	writer, err := utils.GetWriteSyncer(logName) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.SlotsConfig.Zap.Prefix + "2006/01/02 15:04:05.000"))
}
