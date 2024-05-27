/* ????
Необходимо реализовать интерфейс

type Meteo interface {
ReadTemp() string
ChangeTemp(v string)
}

Речь про температуру окружающей среды.
ReadTemp читает температуру, ChangeTemp изменяет температуру. Код должен быть потокобезопасным,
т.е. при работе с температурой нескольких параллельных горутин не должно возникать состояние гонки.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Celsius struct {
	temp int
	mu   sync.RWMutex
}

func (c *Celsius) ReadTemp() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.temp
}
func (c *Celsius) ChangeTemp(temp int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.temp = temp

}

/*
	type Meteo interface {
		ReadTemp() string
		ChangeTemp(v string)
	}
*/
func main() {
	cc := Celsius{temp: 0}

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("Текущая температура окружающей среды:", cc.ReadTemp())
		}()
		go func() {
			cc.ChangeTemp(rand.Intn(50))
		}()

	}
	time.Sleep(2 * time.Second)

}
