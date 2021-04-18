package client

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	RedisDb *redis.Pool
)

func SetUpRedisClient() {
	RedisDb = &redis.Pool{
		// 连接数
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 0,
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")

			return err
		},
	}
}
