package flyredis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"time"
)

var defaultPool = Pool{&redis.Pool{
	MaxIdle:     50,
	MaxActive:   1000,
	IdleTimeout: 30 * time.Second,
	Dial: func() (conn redis.Conn, err error) {
		return redis.Dial("tcp", "127.0.0.1:6379")
	},
}}

type Opt struct {
	Dial            func() (redis.Conn, error)
	DialContext     func(ctx context.Context) (redis.Conn, error)
	TestOnBorrow    func(c redis.Conn, t time.Time) error
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	Wait            bool
	MaxConnLifetime time.Duration
}

func NewPool(pool *redis.Pool) *Pool {
	return &Pool{pool}
}

func Get() Conn {
	return defaultPool.Get()
}

func Do(commandName string, args ...interface{}) Result {
	return defaultPool.Do(commandName, args...)
}

func Send(commandName string, args ...interface{}) error {
	return defaultPool.Send(commandName, args...)
}

func GET(args ...interface{}) Result {
	return defaultPool.GET(args...)
}

func SET(args ...interface{}) Result {
	return defaultPool.SET(args...)
}

func KEYS(args ...interface{}) Result {
	return defaultPool.KEYS(args...)
}

func HGET(args ...interface{}) Result {
	return defaultPool.HGET(args...)
}

func HSET(args ...interface{}) Result {
	return defaultPool.HSET(args...)
}

func HSETNX(args ...interface{}) Result {
	return defaultPool.HSETNX(args...)
}

func HGETALL(args ...interface{}) Result {
	return defaultPool.HGETALL(args...)
}

func HVALS(args ...interface{}) Result {
	return defaultPool.HVALS(args...)
}

func HEXISTS(args ...interface{}) Result {
	return defaultPool.HEXISTS(args...)
}

func HDEL(args ...interface{}) Result {
	return defaultPool.HDEL(args...)
}

func SISMEMBER(args ...interface{}) Result {
	return defaultPool.SISMEMBER(args...)
}

func SADD(args ...interface{}) Result {
	return defaultPool.SADD(args...)
}
