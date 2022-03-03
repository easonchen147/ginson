package cache

import (
	"errors"
	"ginson/pkg/conf"
	"github.com/go-redis/redis/v8"
)

var (
	client        *redis.Client
	clusterClient *redis.ClusterClient
)

// InitRedis 初始化redis
func InitRedis(cfg *conf.AppConfig) error {
	if cfg.RedisConfig == nil {
		return nil
	}
	client = redis.NewClient(&redis.Options{
		Addr:         cfg.RedisConfig.Addr,
		Password:     cfg.RedisConfig.Pass,
		DB:           cfg.RedisConfig.Db,
		MinIdleConns: cfg.RedisConfig.MinIdle,
		PoolSize:     cfg.RedisConfig.PoolSize,
	})
	return nil
}

func Client() *redis.Client {
	if client == nil {
		panic(errors.New("redis is not ready"))
	}
	return client
}

// InitRedisCluster 初始化redis cluster
func InitRedisCluster(cfg *conf.AppConfig) error {
	if cfg.RedisClusterConfig == nil {
		return nil
	}
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        cfg.RedisClusterConfig.Addrs,
		Password:     cfg.RedisClusterConfig.Pass,
		MinIdleConns: cfg.RedisClusterConfig.MinIdle,
		PoolSize:     cfg.RedisClusterConfig.PoolSize,
	})
	return nil
}

func ClusterClient() *redis.ClusterClient {
	if clusterClient == nil {
		panic(errors.New("redis cluster is not ready"))
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
