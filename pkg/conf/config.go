package conf

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	Dev  = "dev"
	Prod = "prod"
	Test = "test"
)

var AppConf *AppConfig

type AppConfig struct {
	File string

	Env           string `toml:"env"`
	HttpAddr      string `toml:"http_addr"`
	HttpPort      int    `toml:"http_port"`
	LogMode       string `toml:"log_mode"`
	LogFile       string `toml:"log_file"`
	LogLevel      string `toml:"log_level"`
	AccessLogFile string `toml:"access_log_file"`

	DbsConfig          map[string]*dbConfig `toml:"dbs"`
	MongoConfig        *mongoConfig         `toml:"mongo"`
	RedisConfig        *redisConfig         `toml:"redis"`
	RedisClusterConfig *redisClusterConfig  `toml:"redis_cluster"`
	KafkaConfig        *kafkaConfig         `toml:"kafka"`

	Ext extConfig `toml:"ext"`
}

type dbConfig struct {
	Uri         string `toml:"uri"`
	MaxIdleConn int    `toml:"max_idle_conn"`
	MaxOpenConn int    `toml:"max_open_conn"`
}

type redisConfig struct {
	Addr     string `toml:"addr"`
	Pass     string `toml:"pass"`
	Db       int    `toml:"db"`
	MinIdle  int    `toml:"min_idle"`
	PoolSize int    `toml:"pool_size"`
}

type redisClusterConfig struct {
	Addrs    []string `toml:"addrs"`
	Pass     string   `toml:"pass"`
	MinIdle  int      `toml:"min_idle"`
	PoolSize int      `toml:"pool_size"`
}

type kafkaConfig struct {
	Consumers map[string]*kafkaConsumerConfig `toml:"consumers"`
	Producers map[string]*kafkaProducerConfig `toml:"producers"`
}

type kafkaConsumerConfig struct {
	Broker    string `toml:"broker"`
	Topic     string `toml:"topic"`
	Group     string `toml:"group"`
	Partition int    `toml:"partition"`
}

type kafkaProducerConfig struct {
	Broker string `toml:"broker"`
	Topic  string `toml:"topic"`
}

type extConfig struct {
	TokenSecret     string `toml:"token_secret"`
	WxMiniAppId     string `toml:"wx_mini_app_id"`
	WxMiniAppSecret string `toml:"wx_mini_app_secret"`
}

type mongoConfig struct {
	Uri            string `toml:"uri"`
	Db             string `toml:"db"`
	ConnectTimeout uint64 `toml:"connect_timeout"`
	MaxOpenConn    uint64 `toml:"max_open_conn"`
	MaxPoolSize    uint64 `toml:"max_pool_size"`
	MinPoolSize    uint64 `toml:"min_pool_size"`
}

func InitConfig(file string) *AppConfig {
	AppConf = &AppConfig{
		File:          file,
		Env:           Dev,
		HttpAddr:      "0.0.0.0",
		HttpPort:      8080,
		LogMode:       "console",
		LogFile:       "logs/app.log",
		LogLevel:      "debug",
		AccessLogFile: "logs/access.log",
	}
	return AppConf
}

// Load 加载toml配置文件内容
func (cfg *AppConfig) Load() error {
	if _, err := os.Stat(cfg.File); os.IsNotExist(err) {
		return fmt.Errorf("config file %s not existed", cfg.File)
	}

	_, err := toml.DecodeFile(cfg.File, cfg)
	if err != nil {
		return fmt.Errorf("load config file %s failed, error: %v", cfg.File, err)
	}

	return nil
}

func (cfg *AppConfig) IsDevEnv() bool {
	return cfg.Env == "dev"
}

func init() {
	configFile := "app.toml"
	if envFilePath := os.Getenv("CONFIG_FILE"); envFilePath != "" {
		configFile = envFilePath
	}

	// 加载配置
	cfg := InitConfig(configFile)
	err := cfg.Load()
	if err != nil {
		panic(fmt.Sprintf("load config failed, file: %s, error: %s", configFile, err))
	}
}
