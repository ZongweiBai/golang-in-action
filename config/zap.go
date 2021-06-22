package config

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// 初始化zaplog
func InitLogger() (*zap.Logger, *zap.SugaredLogger) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	// 读取log level
	level := zap.DebugLevel
	if CONFIG.Log.Level != "" {
		switch CONFIG.Log.Level {
		case LOG_LEVEL_INFO:
			level = zap.InfoLevel
		case LOG_LEVEL_WARN:
			level = zap.WarnLevel
		case LOG_LEVEL_ERROR:
			level = zap.ErrorLevel
		case LOG_LEVEL_DPANIC:
			level = zap.DPanicLevel
		case LOG_LEVEL_PANIC:
			level = zap.PanicLevel
		case LOG_LEVEL_FATAL:
			level = zap.FatalLevel
		default:
			level = zap.DebugLevel
		}
	}

	core := zapcore.NewCore(encoder, writeSyncer, level)

	logger := zap.New(core)
	logger = logger.WithOptions(zap.AddCaller())
	return logger, logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	filePath := "./application.log"
	if CONFIG.Log.FilePath != "" {
		filePath = CONFIG.Log.FilePath + "/application.log"
	}
	// Filename: 日志文件的位置
	// MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
	// MaxBackups：保留旧文件的最大个数
	// MaxAges：保留旧文件的最大天数
	// Compress：是否压缩/归档旧文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
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
