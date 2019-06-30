package flyredis

type RedisInterface interface {
	KeyCommand

	StringCommand

	HashCommand

	SetCommand

	Do(commandName string, args ...interface{}) Result
	Send(commandName string, args ...interface{}) error
}

type KeyCommand interface {
	DEL(args ...interface{}) Result
	EXPIRE(args ...interface{}) Result
	EXISTS(args ...interface{}) Result
	KEYS(args ...interface{}) Result
}

type StringCommand interface {
	GET(args ...interface{}) Result
	SET(args ...interface{}) Result
}

type HashCommand interface {
	HGET(args ...interface{}) Result
	HSET(args ...interface{}) Result
	HDEL(args ...interface{}) Result
	HSETNX(args ...interface{}) Result
	HGETALL(args ...interface{}) Result
	HVALS(args ...interface{}) Result
	HEXISTS(args ...interface{}) Result
}

type SetCommand interface {
	SISMEMBER(args ...interface{}) Result
	SADD(args ...interface{}) Result
}
