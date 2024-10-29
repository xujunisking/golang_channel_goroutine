package main

import (
	"account/redisdemo/demo02/redisManager"
	"fmt"
)

func main() {

	host := "192.168.1.95"
	port := "6379"
	password := "Ynbys2@)@)"
	err := redisManager.NewRedisConn(host, port, password)
	if err != nil {
		fmt.Println("redis connection err:", err)
		return
	}

	fmt.Println("redis connection success...")

	err = redisManager.StringSet("hero", "hello,world!")
	if err != nil {
		fmt.Println("redis StringSet() err:", err)
		return
	}

	bytes, err1 := redisManager.StringGet("hero")
	if err1 != nil {
		fmt.Println("redis StringGet() err:", err1)
		return
	}

	res := string(bytes)

	fmt.Printf("hero=%v\n", res)
}
