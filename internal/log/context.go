package log

import (
	"context"

	"github.com/mateeullahmalik/goa-demo/internal/log/hooks"
)

type (
	// private type used to define context keys
	ctxKey int
)

const (
	// PrefixKey is the prefix of the log record
	PrefixKey ctxKey = iota
)

// ContextWithPrefix returns a new context with PrefixKey value.
func ContextWithPrefix(ctx context.Context, prefix string) context.Context {
	return context.WithValue(ctx, PrefixKey, prefix)
}

func init() {
	AddHook(hooks.NewContextHook(PrefixKey, func(ctxValue interface{}, msg string, fields hooks.ContextHookFields) (string, hooks.ContextHookFields) {
		fields["prefix"] = ctxValue
		return msg, fields
	}))
}
