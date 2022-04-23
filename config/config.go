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
	// MQTT服务配置
	MQTT MQTT `mapstructure:"mqtt" json:"mqtt" yaml:"mqtt"`
	// 数据源配置
	DataSource DataSource `mapstructure:"datasource" json:"datasource" yaml:"datasource"`
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
	Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
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

type MQTT struct {
	Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	// broker
	Broker   string `mapstructure:"broker" json:"broker" yaml:"broker"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	ClientId string `mapstructure:"clientId" json:"clientId" yaml:"clientId"`
	UserName string `mapstructure:"userName" json:"userName" yaml:"userName"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Topics   string `mapstructure:"topics" json:"topics" yaml:"topics"`
}

type DataSource struct {
	Enabled     bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Dialect     string `mapstructure:"dialect" json:"dialect" yaml:"dialect"`
	Scheme      string `mapstructure:"scheme" json:"scheme" yaml:"scheme"`
	Url         string `mapstructure:"url" json:"url" yaml:"url"`
	MaxPoolSize int    `mapstructure:"maxPoolSize" json:"maxPoolSize" yaml:"maxPoolSize"`
	MinIdle     int    `mapstructure:"minIdle" json:"minIdle" yaml:"minIdle"`
}
