package model

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var (
	pool        *redis.Pool
	redisHost   = viper.GetString("db.redis.host")
	redisPasswd = viper.GetString("db.redis.password")
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost, redis.DialPassword(redisPasswd))
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
		MaxIdle:         50,
		MaxActive:       30,
		IdleTimeout:     300 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

//初始化redis连接池
func init() {
	pool = newRedisPool()
}

//对外暴露连接池
func RedisPool() *redis.Pool {
	return pool
}
