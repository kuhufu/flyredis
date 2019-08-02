# flyredis
## 单机
#### example

```go
//redisgo
data, err := conn.Do("GET", "key")
//flyredis
data, err := conn.Do("GET", "key").Value()
```
```go
//redisgo
data, err := redis.String(conn.Do("GET", "key"))
//flyredis
data, err := conn.Do("GET", "key").String()
```

```go
//flyredis
data, err := conn.Do("GET", "key").Bool()
data, err := conn.Do("GET", "key").Int()
data, err := conn.Do("GET", "key").String()
data, err := conn.Do("HGETALL", "key").Ints()
data, err := conn.Do("HGETALL", "key").Strings()
data, err := conn.Do("HGETALL", "key").Values()
......
```

```go
//flyredis
data, err := conn.GET("key").String()
data, err := conn.SET("key", "value")
data, err := conn.HGET("key", "field")
data, err := conn.HSET("key", "filed", "value")
```

### 连接池
```go
var pool = flyredis.NewPool("tcp", "127.0.0.1:7000", flyredis.Option{
	MaxIdle:     10,
	MaxActive:   20,
	IdleTimeout: 20 * time.Second,
})
```
```go
conn := pool.Get()
defer conn.Close()
data, err := conn.GET("key").String()
```

## 集群
```go
client := flyredis.cluster.NewClient("tcp", "127.0.0.1:7000")
fmt.Println(client.Do("GET", "k1").String())
```