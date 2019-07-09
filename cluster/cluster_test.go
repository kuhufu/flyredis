package cluster

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"regexp"
	"strings"
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
	fmt.Println(extractClusterSlotsInfo(reply))
}

func TestSlotNumber(t *testing.T) {
	key := "{k2}:key"
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

func TestNodeEqual(t *testing.T) {
	tests := []struct {
		a    node
		b    node
		want bool
	}{
		{node{}, node{}, true},
		{node{firstSlot: 1}, node{}, false},
		{node{slaves: []addr{}}, node{}, true},
	}
	for _, test := range tests {
		if nodeEqual(test.a, test.b) != test.want {
			t.Error("err")
		}
	}
}

var r, _ = regexp.Compile(`{(.+?)}`)

func TestHashTag(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{key: "{user}:key", want: "user"},
		{key: "{user:key", want: "{user:key"},
		{key: "user}:key", want: "user}:key"},
		{key: "{user}:{key}", want: "user"},
	}
	for _, test := range tests {
		key := test.key
		if r.MatchString(key) {
			key = strings.Trim(r.FindString(key), "{}")
		}
		if key != test.want {
			t.Error("in:", test.key, "out:", key, "want:", test.want)
		}
	}
}

func TestHashTag3(t *testing.T) {
	key := `{user}:key`
	a := strings.Index(key, "{")
	c := strings.Index(key, "}")

	if a != -1 && c != -1 {
		key = key[a+1 : c]
	}

	tests := []struct {
		key  string
		want string
	}{
		{key: "{user}:key", want: "user"},
		{key: "{user:key", want: "{user:key"},
		{key: "user}:key", want: "user}:key"},
		{key: "{user}:{key}", want: "user"},
	}
	for _, test := range tests {
		key := test.key
		a := strings.Index(key, "{")
		c := strings.Index(key, "}")
		if a != -1 && c != -1 {
			key = key[a+1 : c]
		}
		if key != test.want {
			t.Error("in:", test.key, "out:", key, "want:", test.want)
		}
	}
}

//BenchmarkFoo-8    	30000000	        60.4 ns/op
//BenchmarkFoo3-8   	100000000	        13.6 ns/op
func BenchmarkFoo(b *testing.B) {
	key := `{user}:key`
	for i := 0; i < b.N; i++ {
		if r.MatchString(key) {
			key = strings.Trim(r.FindString(key), "{}")
		}
	}
}

func BenchmarkFoo3(b *testing.B) {
	key := `{user}:key`
	for i := 0; i < b.N; i++ {
		a := strings.Index(key, "{")
		c := strings.Index(key, "}")

		if a != -1 && c != -1 {
			key = key[a+1 : c]
		}
	}

}
