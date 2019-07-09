package cluster

import (
	"github.com/kuhufu/flyredis/crc16"
	"strings"
)

const REDIS_CLUSTER_SLOTS = 16384

type addr struct {
	ip   string
	port int
}

type node struct {
	firstSlot int
	lastSlot  int

	master addr
	slaves []addr
}

func extractClusterSlotsInfo(reply []interface{}) []node {
	nodes := make([]node, len(reply))
	for i := 0; i < len(reply); i++ {

		v := reply[i].([]interface{})

		//master
		master := v[2].([]interface{})
		nodes[i] = node{
			firstSlot: int(v[0].(int64)),
			lastSlot:  int(v[1].(int64)),
			master: addr{
				ip:   string(master[0].([]uint8)),
				port: int(master[1].(int64)),
			},
		}

		var slaves []addr
		//slave
		for i := 3; i < len(v); i++ {
			slave := v[i].([]interface{})
			slaves = append(slaves, addr{
				ip:   string(slave[0].([]uint8)),
				port: int(slave[1].(int64)),
			})
		}

		nodes[i].slaves = slaves
	}
	return nodes
}

func SlotNumber(key string) int {
	//支持 hash tag
	l := strings.Index(key, "{")
	r := strings.Index(key, "}")

	if l != -1 && r != -1 {
		key = key[l+1 : r]
	}
	return int(crc16.Checksum([]byte(key)) % REDIS_CLUSTER_SLOTS)
}

func nodesEqual(a, b []node) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !nodeEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func nodeEqual(a, b node) bool {
	if a.firstSlot != b.firstSlot || a.lastSlot != b.lastSlot || a.master != b.master || len(a.slaves) != len(b.slaves) {
		return false
	}
	for i := 0; i < len(a.slaves); i++ {
		if a.slaves[i] != b.slaves[i] {
			return false
		}
	}
	return true
}

func sortNodes(nodes []node) {
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].firstSlot > nodes[j].firstSlot {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}
}
