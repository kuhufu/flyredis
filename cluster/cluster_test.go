package cluster

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"testing"
	"time"
)

var pool = flyredis.NewPool(&redis.Pool{
	MaxIdle:     50,
	MaxActive:   1000,
	IdleTimeout: 10 * time.Second,
	Dial: func() (conn redis.Conn, err error) {
		return redis.Dial("tcp", "127.0.0.1:7000")
	},
})

func TestGetClusterSlotsInfo(t *testing.T) {
	reply, _ := pool.Do("cluster", "slots").Values()
	fmt.Println(reply)
}

func TestSlotNumber(t *testing.T) {
	key := "k2"
	slotNum, err := pool.Do("CLUSTER", "KEYSLOT", key).Int()
	if err != nil {
		t.Fatal(err)
	}
	if SlotNumber(key) != slotNum {
		t.Error("not match redis slot number")
	}
}

func TestGet(t *testing.T) {
	fmt.Println(pool.GET("k1").String())
	fmt.Println(pool.GET("k2").String())
}
