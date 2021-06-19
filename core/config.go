package core

// 全局配置
type Config struct {
	// 日志文件配置
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
}

const (
	LOG_LEVEL_DEBUG = "DEBUG"	
	LOG_LEVEL_INFO = "INFO"
	LOG_LEVEL_WARN = "WARN"
	LOG_LEVEL_ERROR = "ERROR"
	LOG_LEVEL_DPANIC = "DPANIC"
	LOG_LEVEL_PANIC = "PANIC"
	LOG_LEVEL_FATAL = "FATAL"
)

type Log struct {
	// 输出日志
	FilePath string `mapstructure:"filePath" json:"filePath" yaml:"filePath"`
	// 日志级别
	Level    string `mapstructure:"level" json:"level" yaml:"level"`
}
