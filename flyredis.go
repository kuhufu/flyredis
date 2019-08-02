package flyredis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Option struct {
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	Wait            bool
	MaxConnLifetime time.Duration
	Password        string
	TestOnBorrow    func(c redis.Conn, t time.Time) error
	DialOptions     []DialOption
}

type DialOption = redis.DialOption

func NewResult(reply interface{}, err error) Result {
	return Result{reply: reply, err: err}
}

func NewPool(network, address string, option Option) *Pool {
	dialFunc := func() (redis.Conn, error) {
		return redis.Dial(network, address, append(option.DialOptions, redis.DialPassword(option.Password))...)
	}

	if option.TestOnBorrow == nil {
		option.TestOnBorrow = func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		}
	}

	return &Pool{&redis.Pool{
		MaxIdle:         option.MaxIdle,
		MaxActive:       option.MaxActive,
		IdleTimeout:     option.IdleTimeout,
		Wait:            option.Wait,
		MaxConnLifetime: option.MaxConnLifetime,
		TestOnBorrow:    option.TestOnBorrow,
		Dial:            dialFunc,
	}}
}
