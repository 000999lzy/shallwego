package main

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	f := func(ctx context.Context, key string) {
		if v := ctx.Value(key); v != nil {
			fmt.Printf("found value for key %q: %v\n", key, v)
			return
		}

		fmt.Printf("no value found for key %q\n", key)
	}

	k := "myKey"
	ctx := context.Background()
	ctx = context.WithValue(ctx, k, "gogogo")
	f(ctx, k)
	f(ctx, "otherKey")
}
