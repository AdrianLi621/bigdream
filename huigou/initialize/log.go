package initialize

import (
	"bigdream/huigou/pkg"
	"github.com/lestrrat/go-file-rotatelogs"
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
func InitLogger(filename string) *zap.SugaredLogger {
	v := LoadLogConfig()
	logpath := v.Get("LogPath").(string)
	if !pkg.FileOrDirIsExist(logpath) {
		return nil
	}
	fullpath := logpath + "/" + filename
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "file",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	warnIoWriter, _ := rotatelogs.New(
		fullpath+".%Y%m%d",
		rotatelogs.WithMaxAge(time.Hour*24*5),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	// 创建Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(warnIoWriter), zap.DebugLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}
