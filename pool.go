package flyredis

import (
	"context"
	"github.com/gomodule/redigo/redis"
)

type Pool struct {
	inner *redis.Pool
}

var _ RedisInterface = (Pool)(nil)

func (p Pool) Get() Conn {
	return Conn{p.inner.Get()}
}

func (p Pool) Close() error {
	return p.inner.Close()
}

func (p Pool) ActiveCount() int {
	return p.inner.ActiveCount()
}

func (p Pool) IdleCount() int {
	return p.inner.IdleCount()
}

func (p Pool) Stats() redis.PoolStats {
	return p.inner.Stats()
}

func (p Pool) GetContext(ctx context.Context) (Conn, error) {
	conn, err := p.inner.GetContext(ctx)
	return Conn{conn}, err
}

func (p Pool) Do(commandName string, args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.Do(commandName, args...)
}

func (p Pool) Send(commandName string, args ...interface{}) error {
	c := p.Get()
	defer c.Close()
	return c.Send(commandName, args...)
}

func (p Pool) GET(args ...interface{}) Result {
	return p.Do("GET", args...)
}

func (p Pool) SET(args ...interface{}) Result {
	return p.Do("SET", args...)
}

func (p Pool) DEL(args ...interface{}) Result {
	return p.Do("DEL", args...)
}

func (p Pool) EXPIRE(args ...interface{}) Result {
	return p.Do("EXPIRE", args...)
}

func (p Pool) EXISTS(args ...interface{}) Result {
	return p.Do("EXISTS", args...)
}

func (p Pool) KEYS(args ...interface{}) Result {
	return p.Do("KEYS", args...)
}

func (p Pool) HGET(args ...interface{}) Result {
	return p.Do("HGET", args...)
}

func (p Pool) HSET(args ...interface{}) Result {
	return p.Do("HSET", args...)
}

func (p Pool) HSETNX(args ...interface{}) Result {
	return p.Do("HSETNX", args...)
}

func (p Pool) HGETALL(args ...interface{}) Result {
	return p.Do("HGETALL", args...)
}

func (p Pool) HVALS(args ...interface{}) Result {
	return p.Do("HVALS", args...)
}

func (p Pool) HEXISTS(args ...interface{}) Result {
	return p.Do("HEXISTS", args...)
}

func (p Pool) HDEL(args ...interface{}) Result {
	return p.Do("HDEL", args...)
}

func (p Pool) SISMEMBER(args ...interface{}) Result {
	return p.Do("SISMEMBER", args...)
}

func (p Pool) SADD(args ...interface{}) Result {
	return p.Do("SADD", args...)
}
