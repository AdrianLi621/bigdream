package initialize

import (
	"bigdream/huigou/pkg"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

/**
加载日志配置
*/
func LoadLogConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("log")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic("加载log配置文件出错")
	}
	return v
}

/**
实例化日志
*/
func InitLogger(filename string, loglevel string) *zap.Logger {
	v := LoadLogConfig()
	logpath := v.Get("LogPath").(string)
	if !pkg.FileOrDirIsExist(logpath) {
		return nil
	}
	fullpath := logpath + "/" + filename
	hook := lumberjack.Logger{
		Filename:   fullpath, //日志文件路径
		MaxSize:    2048,     //最大字节
		MaxAge:     30,
		MaxBackups: 7,
		Compress:   true,
	}
	writeSyncer := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,      // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writeSyncer,
		level,
	)

	logger := zap.New(core)
	return logger
}
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
