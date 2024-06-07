package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.0.1.84:38079", // Redis 服务器地址
		Password: "",                // 设置空密码
		DB:       0,                 // 使用默认数据库
	})

	// 使用 Ping 命令检查与 Redis 服务器的连接是否正常
	ctx := context.Background()
	scan, uintRes, err := rdb.Scan(ctx, 0, "{123}:*", 100).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("uintRes-->", uintRes)
	fmt.Println("scan-->", scan)

	for _, key := range scan {
		rdb.Rename(ctx, key, "jerry:list"+key)
	}
}
