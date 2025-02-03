package database

import "github.com/gomodule/redigo/redis"

var RedisPool redis.Pool

func init() {
	InitRedis()
}

func InitRedis() {
	//使用redis 连接池
	RedisPool = redis.Pool{
		MaxIdle:         20,     //最大空闲数=初始数
		MaxActive:       50,     //最大存活数(总数)
		MaxConnLifetime: 60 * 5, //最大生命周期
		IdleTimeout:     60,     //空闲超时时间
		Dial: func() (redis.Conn, error) { //连接数据库
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
