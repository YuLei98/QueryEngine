package engineprocessors

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局pool

var POOL *redis.Pool

func InitPool(address string, maxIdle int, maxActive int, idleTmieout time.Duration) {
	POOL = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   //表示和数据库的最大连接数，0:表示没有限制
		IdleTimeout: idleTmieout, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
