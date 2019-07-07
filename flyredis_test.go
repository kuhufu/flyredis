package flyredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

var key = "flyredis_test_key"

func TestNewPool(t *testing.T) {
	p := NewPool(&redis.Pool{
		MaxIdle:     20,
		MaxActive:   10,
		IdleTimeout: 30,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	})

	fmt.Println(p.KEYS("*forum*").Strings())
}

func TestSET(t *testing.T) {
	if err := SET(key, "test_data").Error(); err != nil {
		log.Fatal(err)
	}
}

// redis GET 当 key不存在时，将返回 nil, 也就是说 GET(key).Error() 和 GET(key).Reply() 都等于 nil
// 而当进行类型转换时：res, err := GET(key).String(), res等于 nil, err不等于nil, error string 为：redigo: nil returned
func TestGET(t *testing.T) {
	if _, err := GET(key).String(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(GET(key).Value())
}

func TestDEL(t *testing.T) {
	if reply, err := DEL(key).Int(); reply == 0 {
		if err != nil {
			log.Fatal(err)
		}
		log.Fatalf(`key "%v" not exist`, key)
	}
}

func TestDial(t *testing.T) {
	conn, err := Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(conn.GET(key).String())
}
