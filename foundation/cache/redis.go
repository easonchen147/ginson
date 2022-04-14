package cache

import (
	"errors"
	"ginson/foundation/cfg"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	client        *redis.Client
	clusterClient *redis.ClusterClient
)

// InitRedis 初始化redis
func InitRedis(cfg *cfg.AppConfig) error {
	if cfg.RedisConfig == nil {
		return nil
	}
	client = redis.NewClient(&redis.Options{
		Addr:         cfg.RedisConfig.Addr,
		Password:     cfg.RedisConfig.Pass,
		DB:           cfg.RedisConfig.Db,
		MinIdleConns: cfg.RedisConfig.MinIdle,
		PoolSize:     cfg.RedisConfig.PoolSize,
		DialTimeout:  time.Second * time.Duration(cfg.RedisConfig.ConnectTimeout),
		ReadTimeout:  time.Second * time.Duration(cfg.RedisConfig.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(cfg.RedisConfig.WriteTimeout),
	})
	return nil
}

func Redis() *redis.Client {
	if client == nil {
		panic(errors.New("cache is not ready"))
	}
	return client
}

// InitRedisCluster 初始化redis cluster
func InitRedisCluster(cfg *cfg.AppConfig) error {
	if cfg.RedisClusterConfig == nil {
		return nil
	}
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        cfg.RedisClusterConfig.Addrs,
		Password:     cfg.RedisClusterConfig.Pass,
		MinIdleConns: cfg.RedisClusterConfig.MinIdle,
		PoolSize:     cfg.RedisClusterConfig.PoolSize,
		DialTimeout:  time.Second * time.Duration(cfg.RedisConfig.ConnectTimeout),
		ReadTimeout:  time.Second * time.Duration(cfg.RedisConfig.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(cfg.RedisConfig.WriteTimeout),
	})
	return nil
}

func RedisCluster() *redis.ClusterClient {
	if clusterClient == nil {
		panic(errors.New("foundation cluster is not ready"))
	}
	return clusterClient
}

func Close() {
	if client != nil {
		_ = client.Close()
	}
	if clusterClient != nil {
		_ = clusterClient.Close()
	}
}
