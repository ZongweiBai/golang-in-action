package config

// 全局配置
type Config struct {
	// 日志文件配置
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
	// JWT配置
	Jwt Jwt `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	// socket服务配置
	Socket Socket `mapstructure:"socket" json:"socket" yaml:"socket"`
	// Redis服务配置
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

const (
	LOG_LEVEL_DEBUG  = "DEBUG"
	LOG_LEVEL_INFO   = "INFO"
	LOG_LEVEL_WARN   = "WARN"
	LOG_LEVEL_ERROR  = "ERROR"
	LOG_LEVEL_DPANIC = "DPANIC"
	LOG_LEVEL_PANIC  = "PANIC"
	LOG_LEVEL_FATAL  = "FATAL"
)

type Log struct {
	// 输出日志
	FilePath string `mapstructure:"filePath" json:"filePath" yaml:"filePath"`
	// 日志级别
	Level string `mapstructure:"level" json:"level" yaml:"level"`
}

type Jwt struct {
	// secret
	JwtSecret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	// expires  seconds
	Expires uint64 `mapstructure:"expires" json:"expires" yaml:"expires"`
}

type Socket struct {
	// host
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// port
	Port int `mapstructure:"port" json:"port" yaml:"port"`
}

type Redis struct {
	// host
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// port
	Port int `mapstructure:"port" json:"port" yaml:"port"`
	// password
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	// db
	DB int `mapstructure:"db" json:"db" yaml:"db"`
	// poolsize
	PoolSize int `mapstructure:"poolsize" json:"poolsize"`
}
