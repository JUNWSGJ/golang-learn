package main

import (
	"fmt"
)

// 使用wire后
func main() {
	u := InitializeUserService("foo", 1)
	fmt.Println(u.UserExist(1))
}
