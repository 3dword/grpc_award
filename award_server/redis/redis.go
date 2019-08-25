package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)


/** 原有代码
func GetConn() redis.Conn{
	conn , err := redis.Dial("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", 6379))

	if err != nil {
		log.Println("connect to redis error ", err)
		return nil
	}

	return conn
}
 */

var pool *redis.Pool

/** 优化：使用 redis 连接池 */
func GetConn() redis.Conn{

	p := GetPool()
	if p == nil {
		return nil
	}

	return p.Get()
}

func GetPool() *redis.Pool {
	if pool == nil {
		return &redis.Pool{
			MaxIdle:     2000,
			IdleTimeout: 3 * time.Minute,
			Dial: func() (redis.Conn, error) {
				conn , err := redis.Dial("tcp",
					fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
					redis.DialConnectTimeout(800*time.Millisecond),
					redis.DialWriteTimeout(800*time.Millisecond),
					redis.DialReadTimeout(800*time.Millisecond),
				)

				if err != nil {
					log.Println("connect to redis error ", err)
				}

				return conn , err
			},
		}
	}
	return pool
}