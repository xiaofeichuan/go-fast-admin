package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisTool struct{}

var Redis = new(RedisTool)
var (
	ctx = context.Background()
	rdb *redis.Client
)

// InitRedis 初始化
func InitRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		//Redis connection error
		panic(err)
	} else {
		fmt.Println("Redis connection successful，" + pong)
	}
	return rdb
}

// Set 设置缓存
func (*RedisTool) Set(key string, value interface{}, expiration time.Duration) {
	//序列化
	data, _ := json.Marshal(value)
	rdb.Set(ctx, key, string(data), expiration)
}

// Get 获取缓存
func (*RedisTool) Get(key string, value interface{}) {
	str, _ := rdb.Get(ctx, key).Result()
	//反序列化
	json.Unmarshal([]byte(str), &value)

}

// Del 删除缓存
func (*RedisTool) Del(key string) {
	rdb.Del(ctx, key)
}

// DelByPattern 模糊删除缓存
func (*RedisTool) DelByPattern(pattern string) {
	keys, _ := rdb.Keys(ctx, pattern+"*").Result()
	for i := 0; i < len(keys); i++ {
		rdb.Del(ctx, keys[i])
	}
}

// GetAllKeys 获取所有key
func (*RedisTool) GetAllKeys() (keys []string) {
	iter := rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		keys = append(keys, key)
	}
	return keys
}

// GetValue 获取缓存
func (*RedisTool) GetValue(key string) string {
	str, _ := rdb.Get(ctx, key).Result()
	return str
}
