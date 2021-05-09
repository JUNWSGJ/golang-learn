package main

import (
	"log"
	"sync"
	"time"
)

/**
  学习 sync.Once 用法
 是 Go 标准库提供的使函数只执行一次的实现, 常应用于单例模式，例如初始化配置、保持数据库连接等。
 */


type Config struct {
	Value int
}

var (
	once   sync.Once
	config *Config
)

func loadConfigFromFile()(*Config, error){
	// load config from file...
	return &Config{
		Value: 1,
	}, nil
}

func GetConfig() *Config {
	once.Do(func(){
		var err error
		config , err = loadConfigFromFile()
		if err!=nil {
			log.Println("init config err...")
		}
		log.Println("init config done")
	})
	return config
}

func main() {
	// 10个goroutine同时去获取配置，但是config只会被加载一次
	for i := 0; i < 10; i++ {
		go func() {
			_ = GetConfig()
		}()
	}
	time.Sleep(time.Second)
}
