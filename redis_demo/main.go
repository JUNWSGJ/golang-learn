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

func createRandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func init() {

}

func writeData(num int, valueSize int) error {

	rdb = redis.NewClient(&redis.Options{Addr: RedisIp + ":" + RedisPort, Password: ""})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Printf("连接redis失败: %v\n", err)
		return err
	}
	log.Printf("开始写入数据")
	for i := 0; i <= num; i++ {
		// 生成10个字节长度的key
		key := fmt.Sprintf("%d", 1000000000+i)
		value := createRandStr(valueSize)
		if err := rdb.Set(key, value, 1*time.Hour).Err(); err != nil {
			log.Printf("写入key失败: %s", err.Error())
			return err
		}
	}
	log.Println("写入数据结束")
	return nil
}

func main() {
	keyCount := 100000
	valueSize := 10

	// 向redis写入指定数量，指定大小的随机生成的字符串
	if err := writeData(keyCount, valueSize); err != nil {
		log.Printf("写入数据失败：%s", err)
	}

}
