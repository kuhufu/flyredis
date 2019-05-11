# flyredis

##### redisgo
```go
conn = redis.Conn{...}

data, err := redis.String(conn.Do("GET", "key"))

data, err = conn.Do("GET", "key")
```

##### flyredis

```go
conn = flyredis.Get()

data, err := conn.Do("GET", "key").String()
// or
data, err = conn.GET("key").String()

data, err = conn.Do("GET", "key").Interface()
// or
data, err = conn.GET("key").Interface()
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
data, err := pool.GET("key").String()
```

 