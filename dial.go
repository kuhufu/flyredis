package flyredis

import (
	"crypto/tls"
	"github.com/gomodule/redigo/redis"
	"net"
	"time"
)

//因为Option中Password字段的存在，这个函数失去了意义
//func DialPassword(password string) DialOption {
//	return redis.DialPassword(password)
//}

func DialNetDial(dial func(network, addr string) (net.Conn, error)) DialOption {
	return redis.DialNetDial(dial)
}

func DialDatabase(db int) DialOption {
	return redis.DialDatabase(db)
}

func DialKeepAlive(d time.Duration) DialOption {
	return DialOption(redis.DialKeepAlive(d))
}

func DialConnectTimeout(d time.Duration) DialOption {
	return redis.DialConnectTimeout(d)
}

func DialReadTimeout(d time.Duration) DialOption {
	return redis.DialReadTimeout(d)
}

func DialWriteTimeout(d time.Duration) DialOption {
	return redis.DialWriteTimeout(d)
}

func DialTLSConfig(c *tls.Config) DialOption {
	return redis.DialTLSConfig(c)
}

func DialTLSSkipVerify(skip bool) DialOption {
	return redis.DialTLSSkipVerify(skip)
}

func DialUseTLS(useTLS bool) DialOption {
	return redis.DialUseTLS(useTLS)
}
