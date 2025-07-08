package runtime

import "context"

type tailcallCtxKey struct{}

var tailcallCtxKeyVal = tailcallCtxKey{}

func setTailCall(ctx context.Context) context.Context {
	return context.WithValue(ctx, tailcallCtxKeyVal, true)
}

func getTailCall(ctx context.Context) bool {
	v, ok := ctx.Value(tailcallCtxKeyVal).(bool)
	return ok && v
}
