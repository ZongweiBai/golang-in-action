package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/natefinch/lumberjack"
	"time"
)

// 初始化zaplog
func InitLogger() (*zap.Logger, *zap.SugaredLogger) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	logger = logger.WithOptions(zap.AddCaller())
	return logger, logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
    // Filename: 日志文件的位置
    // MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
    // MaxBackups：保留旧文件的最大个数
    // MaxAges：保留旧文件的最大天数
    // Compress：是否压缩/归档旧文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./application.log",
		MaxSize:    100,
		MaxBackups: 100,
		MaxAge:     30,
		Compress:   true,
	}
	// 不切割文件
	// file, _ := os.Create("./application.log")
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = customTimeEncoder
	encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoder)
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05.001"))
}
