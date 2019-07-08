package cluster

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"github.com/kuhufu/flyredis/crc16"
	"strconv"
	"time"
)

const REDIS_CLUSTER_SLOTS = 16384

type Addr struct {
	IP   string
	Port int
}

type ClusterNode struct {
	FirstSlot int
	LastSlot  int

	Master Addr
	Slaves []Addr
}

func GetClusterSlotsInfo(reply []interface{}) []ClusterNode {
	nodes := make([]ClusterNode, len(reply))
	for i := 0; i < len(reply); i++ {

		v := reply[i].([]interface{})

		//master
		master := v[2].([]interface{})
		nodes[i] = ClusterNode{
			FirstSlot: int(v[0].(int64)),
			LastSlot:  int(v[1].(int64)),
			Master: Addr{
				IP:   string(master[0].([]uint8)),
				Port: int(master[1].(int64)),
			},
		}

		var slaves []Addr
		//slave
		for i := 3; i < len(v); i++ {
			slave := v[i].([]interface{})
			slaves = append(slaves, Addr{
				IP:   string(slave[0].([]uint8)),
				Port: int(slave[1].(int64)),
			})
		}

		nodes[i].Slaves = slaves
	}
	return nodes
}

func SlotNumber(key string) int {
	return int(crc16.Checksum([]byte(key)) % REDIS_CLUSTER_SLOTS)
}

func NewClient(proc, addr string) *ClusterClient {
	conn, err := flyredis.Dial(proc, addr)
	defer conn.Close()
	if err != nil {
		return nil
	}

	reply, err := conn.Do("cluster", "slots").Values()
	if err != nil {
		return nil
	}
	nodes := GetClusterSlotsInfo(reply)

	pools := make([]*flyredis.Pool, len(nodes))
	for i := 0; i < len(nodes); i++ {
		addr := nodes[i].Master.IP + ":" + strconv.Itoa(nodes[i].Master.Port)
		pools[i] = flyredis.NewPool(&redis.Pool{
			MaxIdle:     50,
			MaxActive:   1000,
			IdleTimeout: 50 * time.Second,
			Dial: func() (conn redis.Conn, err error) {
				return redis.Dial("tcp", addr)
			},
		})
	}

	return &ClusterClient{
		nodes: nodes,
		pools: pools,
	}
}
