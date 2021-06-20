package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"math/rand"
	"time"
)

var (
	RedisIp    = "127.0.0.1"
	RedisPort  = "6379"
	expireTime = 600
	rdb        *redis.Client
)

func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func main() {

	rdb = redis.NewClient(&redis.Options{Addr: RedisIp + ":" + RedisPort, Password: ""})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Printf("连接redis失败: %v\n", err)
	}

	log.Println("开始写入数据")
	for i := 0; i <= 100000; i++ {
		key := fmt.Sprintf("%d", 1000000000+i)
		value := RandStr(10)
		if err := rdb.Set(key, value, 1*time.Hour).Err(); err != nil {
			log.Printf("设置key失败: %s", err.Error())
		}
	}

	log.Println("写入数据结束")

}
