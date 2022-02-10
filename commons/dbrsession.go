package commons

import (
	"context"
	"github.com/gocraft/dbr"
)

type ctxKey int

const (
	PGKey ctxKey = iota
)

func newContext(ctx context.Context, key ctxKey, db *dbr.Connection) context.Context {
	return context.WithValue(ctx, key, db.NewSession(&dumbEventReceiver{}))
}

func fromContext(ctx context.Context, key ctxKey) dbr.SessionRunner {
	if ctx == nil {
		panic("context is nil")
	}
	if val, ok := ctx.Value(key).(dbr.SessionRunner); ok {
		return val
	}
	panic("session is not defined in context")
}

func DBSessionNewContext(ctx context.Context, db *dbr.Connection) context.Context {
	return newContext(ctx, PGKey, db)
}

func DBSessionFromContext(ctx context.Context) dbr.SessionRunner {
	return fromContext(ctx, PGKey)
}
