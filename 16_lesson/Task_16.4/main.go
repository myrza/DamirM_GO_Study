/*
Нужно создать контекст со значениями

some key1: some value1
some key2: some value2

Контекст следует передать в функцию, которая выведет значения some key1 и some key2 в stdout.
*/

package main

import (
	"context"
	"fmt"
)

type ctxKey string

func main() {
	ctx := context.Background()
	var key1 ctxKey = "some key1"
	var key2 ctxKey = "some key2"

	ctx = context.WithValue(ctx, key1, "some value1")
	do(ctx, key1)
	ctx = context.WithValue(ctx, key2, "some value2")
	do(ctx, key2)
}
func do(ctx context.Context, key ctxKey) {

	fmt.Println(key, ":", ctx.Value(key))
}
