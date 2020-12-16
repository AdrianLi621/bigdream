package initialize

import (
	"bigdream/huigou/pkg"
	"github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	once.Do(func() {
		if Logger == nil {
			Logger = WriteLogger()
		}
	})
	return Logger
}

func WriteLogger() *zap.SugaredLogger {
	file := "./storage/logs/aa.log"
	if !pkg.MakeFile(file) {
		panic("创建日志文件失败")
	}
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "file",
		EncodeLevel: zapcore.CapitalLevelEncoder, //基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder, //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) { //一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	warnIoWriter, err := rotatelogs.New(
		file+".%Y%m%d",
		rotatelogs.WithLinkName(file),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	// 创建Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(warnIoWriter), zap.DebugLevel),
	)
	logger := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	return logger.Sugar()
}
