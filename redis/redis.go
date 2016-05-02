package redis

import (
	"fmt"
	"log"
	"time"

	rlib "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
)

func Open(ctx context.Context, host string, port int, key string) context.Context {
	pool := newPool(host, port)

	return context.WithValue(ctx, key, pool)
}

func GetConn(ctx context.Context, key string) rlib.Conn {
	pool := ctx.Value(key).(*rlib.Pool)
	return pool.Get()
}

func Close(ctx context.Context, key string) context.Context {
	redis := GetConn(ctx, key)
	if err := redis.Close(); err != nil {
		log.Println("failed to close redis server:", err)
	}

	return context.WithValue(ctx, key, nil)
}

func newPool(host string, port int) *rlib.Pool {
	return &rlib.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (rlib.Conn, error) {
			c, err := rlib.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				log.Println("failed to dial redis server:", err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c rlib.Conn, t time.Time) error {
			_, err := c.Do("PING")
			log.Println("redis server connection error:", err)
			return err
		},
	}
}

func GetCmd(cli rlib.Conn, key string) (string, error) {
	return rlib.String(cli.Do("GET", key))
}
