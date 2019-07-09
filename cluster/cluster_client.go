package cluster

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"log"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	nodes    []node
	pools    map[string]*flyredis.Pool
	slots    [REDIS_CLUSTER_SLOTS]*flyredis.Pool
	slotLock *sync.RWMutex
	close    chan struct{}
}

type commandInfo struct {
	needHash bool
}

var commands = map[string]commandInfo{
	"DEL":       {needHash: true},
	"EXISTS":    {needHash: true},
	"EXPIRE":    {needHash: true},
	"PEXPIRE":   {needHash: true},
	"EXPIREAT":  {needHash: true},
	"PEXPIREAT": {needHash: true},
	"TTL":       {needHash: true},
	"PTTL":      {needHash: true},

	"GET":         {needHash: true},
	"SET":         {needHash: true},
	"SETNX":       {needHash: true},
	"SETEX":       {needHash: true},
	"MGET":        {needHash: true},
	"MSET":        {needHash: true},
	"MGETNX":      {needHash: true},
	"MSETNX":      {needHash: true},
	"APPEND":      {needHash: true},
	"INCR":        {needHash: true},
	"DECR":        {needHash: true},
	"INCRBY":      {needHash: true},
	"DECRBY":      {needHash: true},
	"INCRBYFLOAT": {needHash: true},

	"HGET":         {needHash: true},
	"HSET":         {needHash: true},
	"HDEL":         {needHash: true},
	"HSETNX":       {needHash: true},
	"HMGET":        {needHash: true},
	"HMSET":        {needHash: true},
	"HEXISTS":      {needHash: true},
	"HINCRBY":      {needHash: true},
	"HINCRBYFLOAT": {needHash: true},
	"HLEN":         {needHash: true},
	"HKEYS":        {needHash: true},
	"HVALs":        {needHash: true},
	"HGETALL":      {needHash: true},

	"SADD":      {needHash: true},
	"SREM":      {needHash: true},
	"SCARD":     {needHash: true},
	"SISMEMBER": {needHash: true},
}

func NewClient(proc, addr string, options ...redis.DialOption) (*Client, error) {
	conn, err := flyredis.Dial(proc, addr, options...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err := conn.Do("cluster", "slots").Values()
	if err != nil {
		return nil, err
	}
	nodes := extractClusterSlotsInfo(reply)
	sortNodes(nodes)
	client := &Client{
		nodes:    nodes,
		pools:    map[string]*flyredis.Pool{},
		close:    make(chan struct{}),
		slotLock: &sync.RWMutex{},
	}
	client.updateClusterSlots(nodes)

	go client.watchClusterSlots()

	return client, nil
}

func (c *Client) Do(commandName string, args ...interface{}) flyredis.Result {
	if len(args) < 1 {
		return flyredis.NewResult(nil, errors.New("wrong num of args"))
	}
	key, ok := args[0].(string)
	if !ok {
		return flyredis.NewResult(nil, errors.New("wrong key type"))
	}
	slotNum := SlotNumber(key)
	return c.findNodeBySlot(slotNum).Do(commandName, args...)
}

func (c *Client) findNodeBySlot(slotNum int) *flyredis.Pool {
	c.slotLock.RLock()
	pool := c.slots[slotNum]
	c.slotLock.RUnlock()
	return pool
}

func (c *Client) updateClusterSlots(nodes []node) {
	for _, node := range nodes {
		addr := node.master.ip + ":" + strconv.Itoa(node.master.port)
		var pool *flyredis.Pool
		var ok bool
		if pool, ok = c.pools[addr]; !ok {
			pool = flyredis.NewPool(&redis.Pool{
				MaxIdle:     50,
				MaxActive:   1000,
				IdleTimeout: 50 * time.Second,
				Dial: func() (conn redis.Conn, err error) {
					return redis.Dial("tcp", addr)
				},
			})
			c.pools[addr] = pool
		}
		c.slotLock.Lock()
		for i := node.firstSlot; i <= node.lastSlot; i++ {
			c.slots[i] = pool
		}
		c.slotLock.Unlock()
	}
}

func (c *Client) fetchClusterSlots() []node {
	for _, pool := range c.pools {
		reply, err := pool.Do("cluster", "slots").Values()
		if err != nil {
			log.Println(err)
		}
		nodes := extractClusterSlotsInfo(reply)
		return nodes
	}
	return nil
}

func (c *Client) watchClusterSlots() {
	tick := time.Tick(time.Second * 2)
	for {
		select {
		case <-c.close:
			return
		case <-tick:
			nodes := c.fetchClusterSlots()
			sortNodes(nodes)
			if !nodesEqual(nodes, c.nodes) {
				log.Println("集群发生变化，更新中")
				log.Println("new nodes info:\n", nodes)
				log.Println("old nodes info:\n", c.nodes)
				c.updateClusterSlots(nodes)
				c.nodes = nodes
			}
		}
	}
}

func (c *Client) Close() {
	close(c.close)
}
