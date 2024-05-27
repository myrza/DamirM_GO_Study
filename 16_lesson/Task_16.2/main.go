/*
Нужно запустить 5 горутин и остановить через 2 секунды
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Стоп гоурутина:", i, ctx.Err())
					return
				default:
					fmt.Println("Длинный расчет")
					time.Sleep(time.Second)
				}
			}
		}()
	}

	wg.Wait()

}
