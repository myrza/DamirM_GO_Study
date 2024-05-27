/*
Необходимо запустить 5 горутин.
Не используя time.Sleep нужно обеспечить вывод в консоль каждой горутиной своего уникального сообщения.
Например:
горутина: 1
горутина: 2
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var j int
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			j++
			fmt.Println("Горутна ", j)
		}()
	}
	wg.Wait()

}
