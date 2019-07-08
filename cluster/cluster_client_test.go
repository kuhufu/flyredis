package cluster

import (
	"fmt"
	"testing"
)

var client = NewClient("tcp", "127.0.0.1:7000")

func TestClusterClient_Do(t *testing.T) {
	fmt.Println(client.Do("GET", "k1").String())
	fmt.Println(client.Do("GET", "k2").String())
	fmt.Println(client.Do("CLUSTER", "slots").Values())
}
