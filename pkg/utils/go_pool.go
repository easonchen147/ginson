package utils

import (
	"context"
	"ginson/pkg/log"
	"github.com/panjf2000/ants/v2"
)

var goPool *ants.Pool

func init() {
	pool, err := ants.NewPool(1000, ants.WithLogger(poolLogger), ants.WithPanicHandler(func(i interface{}) {
		log.Error(context.Background(), "goroutine panic: %v", i)
	}))
	if err != nil {
		log.Error(context.Background(), "new go pool failed, error: %v", err)
		return
	}
	goPool = pool
}

type goPoolLogger struct{}

func (g *goPoolLogger) Printf(format string, args ...interface{}) {
	log.Debug(context.Background(), format, args)
}

var poolLogger = &goPoolLogger{}

func GoInPool(task func()) error {
	return goPool.Submit(task)
}
