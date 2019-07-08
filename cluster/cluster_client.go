package cluster

import (
	"errors"
	"github.com/kuhufu/flyredis"
)

type ClusterClient struct {
	pools []*flyredis.Pool
	nodes []ClusterNode
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

func (c *ClusterClient) findNodeBySlot(slotNum int) *flyredis.Pool {
	for i := 0; i < len(c.nodes); i++ {
		if slotNum >= c.nodes[i].FirstSlot && slotNum <= c.nodes[i].LastSlot {
			return c.pools[i]
		}
	}
	return nil
}

func (c *ClusterClient) Do(commandName string, args ...interface{}) flyredis.Result {
	if len(args) < 1 {
		return flyredis.Result{nil, errors.New("wrong arguments")}
	}
	key, ok := args[0].(string)
	if !ok {
		return flyredis.Result{nil, errors.New("wrong key type")}
	}
	slotNum := SlotNumber(key)
	return c.findNodeBySlot(slotNum).Do(commandName, args...)
}
