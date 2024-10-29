package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.95:6379",
		Password: "Ynbys2@)@)", //默认密码为空
		DB:       0,            //默认数据库
		PoolSize: 10,           // 设置连接池大小
	})

	ctx := context.Background()

	err := rdb.Set(ctx, "hero", "hello,world!", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := rdb.Get(ctx, "hero").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hero=", val)
}
