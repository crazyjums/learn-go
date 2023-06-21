package learn_context

import (
	"context"
	"fmt"
)

func CtxWithValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")   // 传递上下文
	ctx = context.WithValue(ctx, "key2", "value2") // 传递上下文
	go func(ctx context.Context) {
		fmt.Println(ctx.Value("key"))
		fmt.Println(ctx.Value("key2"))
	}(ctx)
}
