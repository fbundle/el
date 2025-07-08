package runtime

import "context"

func setTailCall(ctx context.Context) context.Context {
	return context.WithValue(ctx, "tailcall", true)
}

func getTailCall(ctx context.Context) bool {
	v, ok := ctx.Value("tailcall").(bool)
	return ok && v
}
