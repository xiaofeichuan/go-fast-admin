package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// InitRedis 初始化redis
func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       10, // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("连接Redis出现错误：" + err.Error())
	} else {
		fmt.Println("连接Redis成功:" + pong)
	}
	return client
}
