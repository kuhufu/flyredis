package cluster

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestClusterClient_Do(t *testing.T) {
	var client, err = NewClient("tcp", "127.0.0.1:7000")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(client)
	time.Sleep(time.Hour)
}
