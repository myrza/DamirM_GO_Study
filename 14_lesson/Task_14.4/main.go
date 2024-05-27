/*
Если запустить следующий код программы:
package main
import "fmt"
func main() {
ch := make(chan int)
stop := make(chan struct{})
go func() {
<-ch
stop <- struct{}{}
}()
<-stop
fmt.Println("happy end")
}
возникнет ошибка блокировки:
Не меняя логику дочерней горутины, нужно дописать6 логику главной горутины так, чтобы ошибки блокировки не возникало.
Вместо ошибки в консоль должна выводится фраза «happy end».
Так как канал ch является «особым» , то в него ЗАПРЕЩАЕТСЯ писать значения.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	k := 0
	go func() {

		<-ch
		stop <- struct{}{}
	}()

	for {
		select {
		case <-stop:
			fmt.Println(<-stop)
			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(time.Second)
			k++
			if k > 5 {
				fmt.Println("happy end") // Вместо ошибки в консоль должна выводится фраза «happy end».
				return
			}

		}
	}

}
