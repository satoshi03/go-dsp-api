package main

import (
	"runtime"

	"github.com/guregu/kami"
	"golang.org/x/net/context"

	"github.com/satoshi03/go-dsp-api/common/consts"
	"github.com/satoshi03/go-dsp-api/config"
	"github.com/satoshi03/go-dsp-api/fluent"
	"github.com/satoshi03/go-dsp-api/redis"
)

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	conf := config.Read()

	ctx := redis.Open(context.Background(), conf.Redis.Host, conf.Redis.Port, consts.CtxRedisKey)
	defer redis.Close(ctx, consts.CtxRedisKey)

	ctx = fluent.Open(ctx, conf.Fluent.Host, conf.Fluent.Port, consts.CtxFluentKey)
	defer fluent.Close(ctx, consts.CtxFluentKey)

	kami.Context = ctx
	kami.Serve()
}
