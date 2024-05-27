/*
Следующий код программы должен вывести в консоль сообщения:
Утка - Умею летать!
Утка - Умею плавать!
Воробей - Умею летать!

*/

package main

import "fmt"

type Bird interface {
	Fly()
}
type Duck struct{}

func (d Duck) Fly() {
	fmt.Println("Утка - Умею летать!")
}
func (d Duck) Swim() {
	fmt.Println("Утка - Умею плавать!")
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Воробей - Умею летать!")
}
func main() {
	var d, s Bird
	d = Duck{}
	Do(d, true)
	s = Sparrow{}
	Do(s, false)

}
func Do(b Bird, CheckSwim bool) {
	b.Fly()
	if CheckSwim {
		d := Duck{}
		d.Swim()
	}

}
