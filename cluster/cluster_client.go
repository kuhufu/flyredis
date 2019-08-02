package cluster

import (
	"errors"
	"github.com/kuhufu/flyredis"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Client struct {
	nodes      []node
	pools      map[string]*flyredis.Pool
	poolsLock  *sync.RWMutex
	slots      [REDIS_CLUSTER_SLOTS]*flyredis.Pool
	slotLock   *sync.RWMutex
	close      chan struct{}
	dialOption flyredis.Option
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

func NewClient(network, address string, option flyredis.Option) (*Client, error) {
	pool := newPool(network, address, option)

	reply, err := pool.Do("cluster", "slots").Values()
	if err != nil {
		return nil, err
	}
	nodes := extractClusterSlotsInfo(reply)
	sortNodes(nodes)
	client := &Client{
		nodes:      nodes,
		pools:      map[string]*flyredis.Pool{},
		poolsLock:  &sync.RWMutex{},
		close:      make(chan struct{}),
		slotLock:   &sync.RWMutex{},
		dialOption: option,
	}
	client.pools[address] = pool

	//在新建客户端时获取一次集群信息
	client.updateClusterSlots(nodes)

	//TODO 是否需要定期更新集群信息？还是采用惰性更新，在redis 响应 MOVED消息时才更新？
	//go client.watchClusterSlots()

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
	isNil := false
	var pool = c.GetPoolBySlot(slotNum)
	if pool == nil {
		isNil = true
		var err error
		if pool, err = c.randPool(); err != nil {
			return flyredis.NewResult(nil, err)
		}
	}

	result := pool.Do(commandName, args...)
	err := result.Error()
	if !isRedirectErr(err) {
		if isNil {
			c.SetPoolBySlot(slotNum, pool)
		}
		return result
	}

	//处理 MOVED 和 ASK 重定向
	parts := strings.Split(err.Error(), " ")
	msgType, redirectAddr := parts[0], parts[2]
	redirectPool := c.getOrAddPool(redirectAddr)
	conn := redirectPool.Get()
	switch msgType {
	case "MOVED":
		c.SetPoolBySlot(slotNum, redirectPool)
	case "ASK":
		conn.Do("ASKING")
	}

	result = conn.Do(commandName, args...)
	conn.Close()
	return result
}

func (c *Client) GetPoolBySlot(slotNum int) *flyredis.Pool {
	c.slotLock.RLock()
	pool := c.slots[slotNum]
	c.slotLock.RUnlock()
	return pool
}

func (c *Client) SetPoolBySlot(slotNum int, pool *flyredis.Pool) {
	c.slotLock.Lock()
	c.slots[slotNum] = pool
	c.slotLock.Unlock()
}

func (c *Client) updateClusterSlots(nodes []node) {
	for _, node := range nodes {
		addr := node.master.ip + ":" + strconv.Itoa(node.master.port)
		pool := c.getOrAddPool(addr)
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

func newPool(network, address string, option flyredis.Option) *flyredis.Pool {
	return flyredis.NewPool(network, address, option)
}

func (c *Client) randPool() (*flyredis.Pool, error) {
	c.poolsLock.RLock()
	for _, pool := range c.pools {
		c.poolsLock.RUnlock()
		return pool, nil
	}
	c.poolsLock.RUnlock()
	return nil, errors.New("the pools is empty")
}

func (c *Client) getOrAddPool(address string) (pool *flyredis.Pool) {
	c.poolsLock.RLock()
	pool, ok := c.pools[address]
	if !ok {
		c.poolsLock.RUnlock()
		c.poolsLock.Lock()
		pool = newPool("tcp", address, c.dialOption)
		c.pools[address] = pool
		c.poolsLock.Unlock()
	} else {
		c.poolsLock.RUnlock()
	}
	return
}

func isRedirectErr(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	if strings.HasPrefix(errStr, "MOVED") || strings.HasPrefix(errStr, "ASK") {
		return true
	}
	return false
}
