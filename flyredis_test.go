package flyredis

import (
	"fmt"
	"testing"
	"time"
)

var key = "flyredis_test_key"

func TestNewPool(t *testing.T) {
	p := NewPool("tcp", "127.0.0.1:6379", Option{
		MaxIdle:         10,
		MaxActive:       20,
		IdleTimeout:     20 * time.Second,
	})

	fmt.Println(p.KEYS("*forum*").Strings())
}

// redis GET 当 key不存在时，将返回 nil, 也就是说 GET(key).Error() 和 GET(key).Reply() 都等于 nil
// 而当进行类型转换时：res, err := GET(key).String(), res等于 nil, err不等于nil, error string 为：redigo: nil returned
