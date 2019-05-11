package flyredis

import (
	"context"
	"github.com/gomodule/redigo/redis"
)

func (p Pool) Get() Conn {
	return Conn{p.inner.Get()}
}

func (p Pool) Close() error {
	return p.inner.Close()
}

func (p Pool) ActiveCount() int {
	return p.inner.ActiveCount()
}

func (p Pool) IdleCount() int {
	return p.inner.IdleCount()
}

func (p Pool) Stats() redis.PoolStats {
	return p.inner.Stats()
}

func (p Pool) GetContext(ctx context.Context) (Conn, error) {
	conn, err := p.inner.GetContext(ctx)
	return Conn{conn}, err
}
