package flyredis

type RedisInterface interface {
	GET(args ...interface{}) Result
	SET(args ...interface{}) Result
	DEL(args ...interface{}) Result
	EXPIRE(args ...interface{}) Result
	KEYS(args ...interface{}) Result

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
	Send(commandName string, args ...interface{}) error
}
