package flyredis

import (
	"github.com/gomodule/redigo/redis"
)

type Result struct {
	reply interface{}
	err   error
}

func (r Result) Value() (interface{}, error) {
	return r.reply, r.err
}

func (r Result) Bool() (reply bool, err error) {
	return redis.Bool(r.reply, r.err)
}

func (r Result) Int() (reply int, err error) {
	return redis.Int(r.reply, r.err)
}

func (r Result) Int64() (reply int64, err error) {
	return redis.Int64(r.reply, r.err)
}

func (r Result) Float64() (reply float64, err error) {
	return redis.Float64(r.reply, r.err)
}

func (r Result) String() (reply string, err error) {
	return redis.String(r.reply, r.err)
}

func (r Result) IntMap() (reply map[string]int, err error) {
	return redis.IntMap(r.reply, r.err)
}

func (r Result) StringMap() (reply map[string]string, err error) {
	return redis.StringMap(r.reply, r.err)
}

func (r Result) Values() (reply []interface{}, err error) {
	return redis.Values(r.reply, r.err)
}

func (r Result) Bytes() (reply []byte, err error) {
	return redis.Bytes(r.reply, r.err)
}

func (r Result) Ints() (reply []int, err error) {
	return redis.Ints(r.reply, r.err)
}

func (r Result) Int64s() (reply []int64, err error) {
	return redis.Int64s(r.reply, r.err)
}

func (r Result) Float64s() (reply []float64, err error) {
	return redis.Float64s(r.reply, r.err)
}

func (r Result) Strings() (reply []string, err error) {
	return redis.Strings(r.reply, r.err)
}
