/*
Используя пакет atomic, необходимо реализовать счётчик, с которым параллельно могут работать несколько горутин.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var j int32
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&j, 1)
		}()
	}
	wg.Wait()
	fmt.Println("j = ", j)
}
