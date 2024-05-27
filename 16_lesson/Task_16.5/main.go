/*?????
Не используя context и буферизованные каналы необходимо написать программу,
которая будет запускать 10 рабочих горутин и одну капризную управляющую горутину.
Каждая рабочая горутина с задержкой в 1 секунду должна выводить в stdout сообщение «сложные вычисления горутины: 1»,
где 1 - порядковый номер горутины.
Управляющая горутина через 3 секунды после своего запуска должна в stdout вывести «ой, всё!»,
после чего рабочие горутины должны в stdout вывести «stop горутина: 1», где 1 - порядковый номер горутины,
и завершить своё выполнение.
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
