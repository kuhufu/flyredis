package cluster

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestClient_Do_SET(t *testing.T) {
	c, err := NewClient("tcp", "127.0.0.1:7000")
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
	c, err := NewClient("tcp", "127.0.0.1:7000")
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 50000; i++ {
		if res, _ := c.Do("GET", strconv.Itoa(i)).Int(); res != i {
			t.Error()
		}
	}
}

func TestClient_Do_Race(t *testing.T) {
	c, err := NewClient("tcp", "127.0.0.1:7000")
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
