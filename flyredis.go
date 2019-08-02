package flyredis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Option struct {
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
	Wait bool
	MaxConnLifetime time.Duration
	Password string
}

func NewResult(reply interface{}, err error) Result {
	return Result{reply: reply, err: err}
}

func NewPool(network, address string, option Option) *Pool {
	return &Pool{&redis.Pool{
		MaxIdle:         option.MaxIdle,
		MaxActive:       option.MaxActive,
		IdleTimeout:     option.IdleTimeout,
		Wait:            option.Wait,
		MaxConnLifetime: option.MaxConnLifetime,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return dial(network, address, option.Password)
		},
	}}
}

func dial(network, address, password string) (redis.Conn, error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, err := c.Do("AUTH", password); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, err
}
