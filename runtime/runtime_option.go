package runtime

import "context"

const (
	ENABLE_TCO = false
)

type tailcallCtxKeyType struct{}

var tailcallCtxKey = tailcallCtxKeyType{}

func isTailCall(ctx context.Context) bool {
	_, ok := ctx.Value(tailcallCtxKey).(bool)
	return ok
}

func withTailCall(ctx context.Context) context.Context {
	if ENABLE_TCO {
		return context.WithValue(ctx, tailcallCtxKey, true)
	}
	return ctx
}
