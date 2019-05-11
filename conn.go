package flyredis

func (c *Conn) GET(args ...interface{}) Result {
	reply, err := c.Conn.Do("GET", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) SET(args ...interface{}) Result {
	reply, err := c.Conn.Do("SET", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) HGET(args ...interface{}) Result {
	reply, err := c.Conn.Do("HGET", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) HSET(args ...interface{}) Result {
	reply, err := c.Conn.Do("HSET", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) HSETNX(args ...interface{}) Result {
	reply, err := c.Conn.Do("HSETNX", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) HGETALL(args ...interface{}) Result {
	reply, err := c.Conn.Do("HGETALL", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) HVALS(args ...interface{}) Result {
	reply, err := c.Conn.Do("HVALS", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) HEXISTS(args ...interface{}) Result {
	reply, err := c.Conn.Do("HEXISTS", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) HDEL(args ...interface{}) Result {
	reply, err := c.Conn.Do("HDEL", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) SISMEMBER(args ...interface{}) Result {
	reply, err := c.Conn.Do("SISMEMBER", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) SADD(args ...interface{}) Result {
	reply, err := c.Conn.Do("SADD", args...)
	return Result{reply: reply, err: err}
}

func (c *Conn) Do(commandName string, args ...interface{}) Result {
	reply, err := c.Conn.Do(commandName, args...)
	return Result{reply: reply, err: err}
}