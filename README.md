# flyredis

#### example

```go
//redisgo
data, err := conn.Do("GET", "key")
//flyredis
data, err := conn.Do("GET", "key").Interface()
```
```go
//redisgo
data, err := redis.String(conn.Do("GET", "key"))
//flyredis
data, err := conn.DO("GET", "key").String()
```

```go
//flyredis
data, err := conn.Do("GET", "key").Bool()
data, err := conn.Do("GET", "key").Int()
data, err := conn.DO("GET", "key").String()
data, err := conn.Do("HGETALL", "key").Ints()
data, err := conn.DO("HGETALL", "key").Strings()
data, err := conn.DO("HGETALL", "key").Values()
......
```

```go
//flyredis
data, err := conn.GET("key")
data, err := conn.SET("key", "value")
data, err := conn.HGET("key", "field")
data, err := conn.HSET("key", "filed", "value")
```

### 注意
flyredis 有一个内置的 redis连接池 只连接本地 `127.0.0.1:6379` 的无密码 redis 服务，方便开发测试。

在生产环境中请使用如下方式
```go
pool := flyredis.NewPool(&redis.Pool{
    MaxIdle:     50,
    MaxActive:   1000,
    IdleTimeout: 30 * time.Second,
    Dial: func() (conn redis.Conn, err error) {
        return redis.Dial("tcp", "127.0.0.1:6379")
    },
})
```
```go
conn := pool.Get()
defer conn.Close()
data, err := conn.GET("key").String()
```