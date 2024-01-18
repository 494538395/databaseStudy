package main

import "github.com/redis/go-redis/v9"

func Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	cmd := rdb.Eval(globalCtx, script, keys, args...)

	return cmd.Result()
}

func ZScore(key string, member string) (float64, error) {
	cmd := rdb.ZScore(globalCtx, key, member)

	return cmd.Result()
}

// ZRevRangeByScore 返回指定分数区间内的用户 (倒序)，不带用户分数
func ZRevRangeByScore(key string, min, max string, offset, count int64) ([]string, error) {
	opt := &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	cmd := rdb.ZRevRangeByScore(globalCtx, key, opt)

	return cmd.Result()
}

// ZRevRank 返回 ZSet 中成员的反向排名
func ZRevRank(key string, member string) (int64, error) {
	cmd := rdb.ZRevRank(globalCtx, key, member)

	return cmd.Result()
}

func ZRevRangeByScoreWithScores(key string, min, max string, offset, count int64) ([]redis.Z, error) {
	opt := &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	cmd := rdb.ZRevRangeByScoreWithScores(globalCtx, key, opt)

	return cmd.Result()
}
