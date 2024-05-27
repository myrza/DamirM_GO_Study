/*
Посмотрите внимательно на код:
...
Запустите код.
Затем посмотрите внимательно на звёзды и скажите: сколько горутин запускалось, а также расскажите о дальнейшей их судьбе.
*/

/*
Объяснение: Запустились обе гоурутины.
Первая гоурутина читает в ch.останавливает работу после stop.
Втроая гоуритина записывает в ch. Однако работу не останавливает
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)

			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	stop <- struct{}{}

	fmt.Println("завершение работы главной горутины")
}
