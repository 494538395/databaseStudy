package main

import (
	"context"
	"fmt"

	"database-study/redis/go-redis/lua"

	"github.com/gogf/gf/util/gconv"
)

var globalCtx = context.Background()

func init() {
	initRedis()
}

func main() {
	key := "test-key"
	data := []interface{}{"500", 60, "user01", 61, "user02", 62, "user03", 64, "user04", 65, "user05"}

	cmd, err := Eval(lua.ZsetADDMember, []string{key}, data)
	if err != nil {
		panic(err)
	}

	fmt.Println(cmd)

	users, err := ZRevRangeByScore(key, gconv.String(64), gconv.String(64), 0, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	rank, err := ZRevRank(key, "user02")
	if err != nil {
		panic(err)
	}
	fmt.Println(rank)

	zs, err := ZRevRangeByScoreWithScores(key, gconv.String(50), gconv.String(65), 1, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(zs)

}
