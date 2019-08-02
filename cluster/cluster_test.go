package cluster

import (
	"github.com/kuhufu/flyredis"
	"regexp"
	"strings"
	"testing"
	"time"
)

var pool = flyredis.NewPool("tcp", "127.0.0.1:7000", flyredis.Option{
	MaxIdle:     10,
	MaxActive:   20,
	IdleTimeout: 20 * time.Second,
})

func TestGetClusterSlotsInfo(t *testing.T) {
	_, err := pool.Do("cluster", "slots").Values()
	if err != nil {
		t.Error(err)
	}
}

func TestSlotNumber(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{name: "{in}.name", in: "{in}.name", want: true},
		{name: "{{in}}.name", in: "{{in}}.name", want: true},
		{name: "{in}}.name", in: "{in}}.name", want: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			num, _ := pool.Do("CLUSTER", "KEYSLOT", test.in).Int()
			myNum := SlotNumber(test.in)
			if num != myNum {
				t.Errorf("redis:%v, my:%v\n", num, myNum)
			}
		})
	}
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
		{key: "{user}:in", want: "user"},
		{key: "{user:in", want: "{user:in"},
		{key: "user}:in", want: "user}:in"},
		{key: "{user}:{in}", want: "user"},
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
	key := `{user}:in`
	a := strings.Index(key, "{")
	c := strings.Index(key, "}")

	if a != -1 && c != -1 {
		key = key[a+1 : c]
	}

	tests := []struct {
		key  string
		want string
	}{
		{key: "{user}:in", want: "user"},
		{key: "{user:in", want: "{user:in"},
		{key: "user}:in", want: "user}:in"},
		{key: "{user}:{in}", want: "user"},
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
	key := `{user}:in`
	for i := 0; i < b.N; i++ {
		if r.MatchString(key) {
			key = strings.Trim(r.FindString(key), "{}")
		}
	}
}

func BenchmarkFoo3(b *testing.B) {
	key := `{user}:in`
	for i := 0; i < b.N; i++ {
		a := strings.Index(key, "{")
		c := strings.Index(key, "}")

		if a != -1 && c != -1 {
			key = key[a+1 : c]
		}
	}

}
