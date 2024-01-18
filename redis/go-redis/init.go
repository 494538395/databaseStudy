package main

import "github.com/redis/go-redis/v9"

var rdb redis.UniversalClient

var (
	address      = []string{"127.0.0.1:26379"}
	DB           = 5
	PoolSize     = 2
	MinIdleConns = 2
	Password     = "123456"
	isCluster    = false
)

func initRedis() {
	opts := &redis.UniversalOptions{
		Addrs:        address,
		DB:           DB,
		PoolSize:     PoolSize,
		MinIdleConns: MinIdleConns,
		Password:     Password,
	}

	if isCluster {
		// 集群模式
		//obj := redis.NewClusterClient(opts.Cluster())
	} else {
		// 自动识别模式
		rdb = redis.NewUniversalClient(opts)
	}
}
