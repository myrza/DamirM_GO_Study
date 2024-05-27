/*
Нужно запустить 5 горутин и остановить, используя контекст с отменой
*/

package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Стоп гоурутина:", i)
					return
				default:
					fmt.Println("Длинный расчет")
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Stop")
		cancel()

	}()
	wg.Wait()

}
