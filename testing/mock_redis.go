package lib_test

import (
	"sync"
	"errors"
)

import (
	"github.com/garyburd/redigo/redis"
)

type MockCommand struct {
	Command string
	Args []interface{}
}

func NewRedisPool() *redis.Pool {
	var conn = MockRedisConn{
		Reply: "OK",
	}
	return &redis.Pool{
		Dial: func() (c redis.Conn, err error) {
			return &conn, nil
		},
	}
}

type MockRedisConn struct {
	sync.Mutex
	err error
	Closed bool
	Buffer []MockCommand
	Commands []MockCommand
	Reply interface{}
}

func (c *MockRedisConn) Close() error {
	c.Lock()
	defer c.Unlock()
	if c.Closed {
		c.err = errors.New("Closed called on closed conn") 
		return c.err
	}
	c.Closed = true
	return nil
}

func (c *MockRedisConn) Err() error {
	c.Lock()
	defer c.Unlock()
	return c.err
}

func (c *MockRedisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c.Lock()
	defer c.Unlock()
	if c.Closed {
		c.err = errors.New("Do called on closed conn") 
		return nil, c.err
	}
	c.Commands = append(c.Commands, MockCommand{commandName, args})
	return c.Reply, nil
}

func (c *MockRedisConn) Send(commandName string, args ...interface{}) error {
	c.Lock()
	defer c.Unlock()
	if c.Closed {
		c.err = errors.New("Send called on closed conn") 
		return c.err
	}
	c.Buffer = append(c.Buffer, MockCommand{commandName, args})
	return nil
}

func (c *MockRedisConn) Flush() error {
	c.Lock()
	defer c.Unlock()
	if c.Closed {
		c.err = errors.New("Flush called on closed conn") 
		return c.err
	}
	c.Commands = append(c.Commands, c.Buffer...)
	c.Buffer = []MockCommand{}
	return nil
}

func (c *MockRedisConn) Receive() (reply interface{}, err error) {
	c.Lock()
	defer c.Unlock()
	if c.Closed {
		c.err = errors.New("Receive called on closed conn") 
		return nil, c.err
	}
	return c.Reply, nil
}
