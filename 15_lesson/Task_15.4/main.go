/*
Необходимо создать функцию start, которая в консоль выводит некоторое сообщение.
Необходимо запустить 10 горутин, которые будут запускать функцию start и выводить в консоль факт своего запуска,
причём необходимо обеспечить однократный запуск функции start
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func start(m string) {
	fmt.Println(m)
}
func main() {
	var j int32
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			j++
			start(fmt.Sprintf("%v-я горутина запущена", j))
		}()

	}
	time.Sleep(2 * time.Second)

}
