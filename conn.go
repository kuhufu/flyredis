package flyredis

import "github.com/gomodule/redigo/redis"

type Conn struct {
	redis.Conn
}

func (c *Conn) Do(commandName string, args ...interface{}) Result {
	reply, err := c.Conn.Do(commandName, args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) EXPIRE(args ...interface{}) Result {
	return c.Do("EXPIRE", args...)
}

func (c *Conn) GET(args ...interface{}) Result {
	return c.Do("GET", args...)
}

func (c *Conn) SET(args ...interface{}) Result {
	return c.Do("SET", args...)
}

func (c *Conn) KEYS(args ...interface{}) Result {
	return c.Do("KEYS", args...)
}

func (c *Conn) HGET(args ...interface{}) Result {
	return c.Do("HGET", args...)
}

func (c *Conn) HSET(args ...interface{}) Result {
	return c.Do("HSET", args...)

}

func (c *Conn) HSETNX(args ...interface{}) Result {
	return c.Do("HSETNX", args...)
}

func (c *Conn) HGETALL(args ...interface{}) Result {
	return c.Do("HGETALL", args...)
}

func (c *Conn) HVALS(args ...interface{}) Result {
	return c.Do("HVALS", args...)
}

func (c *Conn) HEXISTS(args ...interface{}) Result {
	return c.Do("HEXISTS", args...)
}

func (c *Conn) HDEL(args ...interface{}) Result {
	return c.Do("HDEL", args...)
}

func (c *Conn) SISMEMBER(args ...interface{}) Result {
	return c.Do("SISMEMBER", args...)
}

func (c *Conn) SADD(args ...interface{}) Result {
	return c.Do("SADD", args...)
}
