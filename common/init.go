package common

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	RedisClient   *redis.Client
	MongoDBClient *mongo.Client
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.Addr"),
		Password:     viper.GetString("redis.Password"),
		DB:           viper.GetInt("redis.DB"),
		MinIdleConns: viper.GetInt("redis.MinIdleConns"),
		PoolSize:     viper.GetInt("redis.PoolSize"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := client.Ping(ctx).Result()
	// 错误处理
	if err != nil {
		fmt.Println("=> REDIS 初始化失败...")
		panic(err)
	}
	fmt.Println("=> REDIS 初始化成功...")
	fmt.Println("=>", pong)
	RedisClient = client
}

const (
	PublishKey = "websocket"
)

// 发布消息到 REDIS
func Publis(ctx context.Context, channel string, message string) error {
	err := RedisClient.Publish(ctx, channel, message).Err()
	fmt.Println("=> 正在发布消息", message)
	return err
}

// 订阅 REDIS 的消息
func Subscribe(ctx context.Context, channel string) {
	// There is no error because go-redis automatically reconnects on error.
	pubsub := RedisClient.Subscribe(ctx, "mychannel1")

	// Close the subscription when we are done.
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}
