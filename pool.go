package flyredis

import (
	"context"
	"github.com/gomodule/redigo/redis"
)

type Pool struct {
	inner *redis.Pool
}

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
	c := p.Get()
	defer c.Close()
	return c.GET(args...)
}

func (p Pool) SET(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.SET(args...)
}

func (p Pool) DEL(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.DEL(args...)
}

func (p Pool) EXPIRE(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.EXPIRE(args...)
}

func (p Pool) KEYS(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.KEYS(args...)
}

func (p Pool) HGET(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.HGET(args...)
}

func (p Pool) HSET(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.HSET(args...)
}

func (p Pool) HSETNX(args ...interface{}) Result {
	c := Get()
	defer c.Close()
	return c.HSETNX(args...)
}

func (p Pool) HGETALL(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.HGETALL(args...)
}

func (p Pool) HVALS(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.HVALS(args...)
}

func (p Pool) HEXISTS(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.HEXISTS(args...)
}

func (p Pool) HDEL(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.HDEL(args...)
}

func (p Pool) SISMEMBER(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.SISMEMBER(args...)
}

func (p Pool) SADD(args ...interface{}) Result {
	c := p.Get()
	defer c.Close()
	return c.SADD(args...)
}
