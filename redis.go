package lib

import (
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

func NewRedisPool(c *Address) (pool *redis.Pool, err error) {
	pool = &redis.Pool{
		MaxIdle:   5,
		MaxActive: 100,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", net.JoinHostPort(c.Host, c.Port))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	_, err = pool.Dial()
	return
}
