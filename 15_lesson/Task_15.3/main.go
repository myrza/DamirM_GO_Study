/*
Используя Mutex, необходимо реализовать счётчик, с которым параллельно могут работать несколько горутин
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var j int32
	var mu sync.Mutex

	for i := 0; i < 500; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			j++
		}()
		go func() {
			mu.Lock()
			defer mu.Unlock()
			j++
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("j = ", j)
}
