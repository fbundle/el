package runtime

import "context"

type tailcallCtxKeyType struct{}

var tailcallCtxKey = tailcallCtxKeyType{}

func isTailCall(ctx context.Context) bool {
	_, ok := ctx.Value(tailcallCtxKey).(bool)
	return ok
}

func withTailCall(ctx context.Context) context.Context {
	return ctx
	return context.WithValue(ctx, tailcallCtxKey, true)
}
