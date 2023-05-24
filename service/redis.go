package service

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	RedisClient *redis.Client
	Context     context.Context
)

func InitRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"),
		DB:       viper.GetInt("redis.DB"),
		// PoolSize:     viper.GetInt("redis.PoolSize"),
		// MinIdleConns: viper.GetInt("redis.MinIdleConns"),
	})

	Context = context.Background()

	// err := rdb.Set(Context, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(Context, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("=> key", val)

	// val2, err := rdb.Get(Context, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("=> key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("=> key2", val2)
	// }

	// Output: key value
	// key2 does not exist

	fmt.Println("=> Redis 初始化成功")
	RedisClient = rdb
}
