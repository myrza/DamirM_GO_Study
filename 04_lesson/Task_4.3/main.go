/*
Необходимо создать и вызвать функцию «hello», которая в качестве параметра принимает и вызывает анонимную функцию.
Анонимная функция должна выводить в stdout фразу «Hello, Go!».
*/
package main

import "fmt"

func hello(f func()) {
	f()
}

func main() {
	f := func() {
		fmt.Print("Hello, Go!")
	}
	hello(f)

}
