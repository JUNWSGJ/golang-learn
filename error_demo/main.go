package main

import (
	"golang-demo/error_demo/service"
	"log"
)

func main() {

	if err := service.CreateUser("张三", 20); err != nil {
		log.Printf("创建用户失败，%+v", err)
	}
	if err := service.CreateUser("李四", 25); err != nil {
		log.Printf("创建用户失败，%+v", err)
	}
	if err := service.CreateUser("王五", 25); err != nil {
		log.Printf("创建用户失败，%+v", err)
	}
	if err := service.CreateUser("张三", 25); err != nil {
		log.Printf("创建用户失败，%+v", err)
	}

	user, err1 := service.FindUserById(6)
	if err1 != nil {
		log.Printf("根据ID查找用户失败，%+v", err1)
	}
	if user == nil {
		log.Printf("未查找到Id为6的用户")
	} else {
		log.Printf("查找到Id为6的用户: %v", user == nil)
	}

	users, err2 := service.FindUsersByAge(30)
	if err2 != nil {
		log.Printf("根据年龄查找用户失败，%+v", err1)
	}
	log.Printf("共查找到年龄为25的用户%d个", len(*users))

}
