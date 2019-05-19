package flyredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"log"
	"testing"
)

func TestPubSub(t *testing.T) {
	conn := flyredis.Get().Conn

	subConn := redis.PubSubConn{Conn: conn}
	err := subConn.Subscribe("chann")
	if err != nil {
		log.Println(err)
	}

	for {
		switch m := subConn.Receive().(type) {
		case redis.Message:
			fmt.Println(m.Channel, string(m.Data))
		case redis.Subscription:
			fmt.Println(m.Channel, m.Kind, m.Count)
		case error:
			fmt.Println(err)
		}

	}

}
