package main

import (
	"runtime"

	"github.com/guregu/kami"
	"golang.org/x/net/context"

	"github.com/satoshi03/go-dsp-api/fluent"
	"github.com/satoshi03/go-dsp-api/redis"
)

var (
	rediskey  = "redisdb"
	fluentkey = "fluent"
)

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	ctx := redis.Open(context.Background(), "127.0.0.1", 6379, rediskey)
	defer redis.Close(ctx, rediskey)

	ctx = fluent.Open(ctx, "127.0.0.1", 24224, fluentkey)
	defer fluent.Close(ctx, fluentkey)

	kami.Context = ctx
	kami.Serve()
}
