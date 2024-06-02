/*
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

type Meteo interface {
	ReadTemp() int
	ChangeTemp(t int)
}

type Celsius struct {
	temp int
	mu   sync.RWMutex
}

func (c *Celsius) ReadTemp() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.temp
}

func (c *Celsius) ChangeTemp(t int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.temp = t

}

func main() {
	c := Celsius{temp: 0}
	for i := 0; i < 100; i++ {
		go read(&c)
		go change(&c)

	}
	time.Sleep(2 * time.Second)

}

func read(m Meteo) {
	fmt.Println("Текущая температура окружающей среды:", m.ReadTemp())

}

func change(m Meteo) {
	t := rand.Intn(50)
	m.ChangeTemp(t)
}
