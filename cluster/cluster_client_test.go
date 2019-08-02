package cluster

import (
	"fmt"
	"github.com/kuhufu/flyredis"
	"log"
	"strconv"
	"testing"
	"time"
)

var c, err = NewClient("tcp", "127.0.0.1:7000", flyredis.Option{
	MaxIdle:     10,
	MaxActive:   20,
	IdleTimeout: 20 * time.Second,
})

func TestClient_Do_SET(t *testing.T) {
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 50000; i++ {
		if err := c.Do("SET", strconv.Itoa(i), i).Error(); err != nil {
			t.Error(err)
		}
	}
}

func TestClient_Do_GET(t *testing.T) {
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 50000; i++ {
		if res, err := c.Do("GET", strconv.Itoa(i)).Int(); res != i {
			t.Error("key:", i, "error:", err)
		}
	}
}

func TestClient_Do_Race(t *testing.T) {
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 50000; i++ {
		i := i
		go func() {
			c.Do("GET", strconv.Itoa(i))
		}()
		fmt.Println("finished")
	}
}
