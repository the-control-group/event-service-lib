package lib

import (
	"github.com/garyburd/redigo/redis"
	"github.com/the-control-group/redissync"
	"net"
	"time"
)

func NewRedisPool(c Address) (pool *redis.Pool, err error) {
	pool = &redis.Pool{
		MaxIdle:   5,
		MaxActive: 10,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", net.JoinHostPort(c.Host, c.Port))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			// _, err := c.Do("PING")
			return err
		},
	}
	_, err = pool.Dial()
	return
}

func NewReloadLock(pool *redis.Pool, processName, hostname string) *redissync.RedisSync {
	return &redissync.RedisSync{Key: processName + "_" + hostname + "__restarting", Pool: pool, Timeout: 30 * time.Minute, Delay: 1 * time.Second, Expiry: 5 * time.Minute, ErrChan: make(chan error, 1)}
}
