package redisManager

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

func NewRedisConn(host, port, password string) error {
	address := fmt.Sprintf("%s:%s", host, port)
	RedisConn = &redis.Pool{
		// Maximum number of idle connections in the pool.(连接池idle连接数)
		MaxIdle: 30,
		// Maximum number of connections allocated by the pool at a given time.(连接池在给定时间内分配的最大连接数)
		// When zero, there is no limit on the number of connections in the pool(当连接池中的连接数为0时，没有限制)
		MaxActive: 30,

		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: 300 * time.Second,

		// If Wait is true and the pool is at the MaxActive limit, then Get() waits
		// for a connection to be returned to the pool before returning.
		Wait: true,

		// Close connections older than this duration. If the value is zero, then
		// the pool does not close connections based on age.
		MaxConnLifetime: 200 * time.Second,

		//创建和配置链接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			if password != "" && password != "none" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}
