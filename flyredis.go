package flyredis

import (
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

type Conn struct {
	redis.Conn
}

type Pool struct {
	inner *redis.Pool
}


type Result struct {
	reply interface{}
	err error
}

type Interface interface {
	GET(args ...interface{}) Result
	SET(args ...interface{}) Result

	HGET(args ...interface{}) Result
	HSET(args ...interface{}) Result
	HSETNX(args ...interface{}) Result
	HGETALL(args ...interface{}) Result
	HVALS(args ...interface{}) Result
	HEXISTS(args ...interface{}) Result
	HDEL(args ...interface{}) Result

	SISMEMBER(args ...interface{}) Result
	SADD(args ...interface{}) Result

	Do(commandName string, args ...interface{}) Result
}

func NewPool(pool *redis.Pool) *Pool {
	return &Pool{pool}
}

func Get() Conn {
	return defaultPool.Get()
}

func GET(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.GET(args...)
}

func SET(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.SET(args...)
}

func HGET(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.HGET(args...)
}

func HSET(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.HSET(args...)
}

func HSETNX(args ...interface{}) Result {
	c := Get()
	defer c.Close()
	return c.HSETNX(args...)
}

func HGETALL(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.HGETALL(args...)
}

func HVALS(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.HVALS(args...)
}

func HEXISTS(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.HEXISTS(args...)
}

func HDEL(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.HDEL(args...)
}

func SISMEMBER(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.SISMEMBER(args...)
}

func SADD(args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.SADD(args...)
}

func  Do(commandName string, args ...interface{}) Result {
	c := defaultPool.Get()
	defer c.Close()
	return c.Do(commandName, args...)
}

func  Send(commandName string, args ...interface{}) error{
	c := defaultPool.Get()
	defer c.Close()
	return  c.Send(commandName, args...)
}