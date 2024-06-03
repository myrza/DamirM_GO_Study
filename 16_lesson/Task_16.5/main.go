/*
Нужно запустить 5 горутин и остановить в некоторое время, которое рассчитывается по формуле: текущий момент + 2 секунды
*/

package main

import (
	//"context"
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool) // сигнальный канал для остановки горутин

	for i := 1; i < 11; i++ {

		go func() {
			for {
				select {
				case <-quit:
					fmt.Println("Стоп гоурутина:", i)
					quit <- true
					return
				default:
					fmt.Println("сложные вычисления горутины: ", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("ой, все!")
	quit <- true

	time.Sleep(10 * time.Second) // ждем окончания работы

}
