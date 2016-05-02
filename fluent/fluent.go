package fluent

import (
	"log"
	"time"

	flib "github.com/fluent/fluent-logger-golang/fluent"
	"golang.org/x/net/context"
)

func Open(ctx context.Context, host string, port int, key string) context.Context {
	// TODO: implement async run by go routine
	f, err := flib.New(flib.Config{FluentPort: port, FluentHost: host})
	if err != nil {
		panic(err)
	}
	return context.WithValue(ctx, key, f)
}

func Send(ctx context.Context, key, tag string, data map[string]interface{}) {
	f := ctx.Value(key).(*flib.Fluent)
	data["created_at"] = time.Now().Format("2006-01-02 15:04:05 -0700")
	if err := f.Post(tag, data); err != nil {
		log.Println(err)
	}
}

func Close(ctx context.Context, key string) context.Context {
	f := ctx.Value(key).(*flib.Fluent)
	if err := f.Close(); err != nil {
		log.Println(err)
	}
	return context.WithValue(ctx, key, nil)
}
