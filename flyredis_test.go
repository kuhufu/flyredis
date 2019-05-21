package flyredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestNewPoolWith(t *testing.T) {
	p := NewPoolWith(Opt{
		MaxIdle:     20,
		MaxActive:   10,
		IdleTimeout: 30,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	})

	fmt.Println(p.KEYS("*forum*").Strings())
}