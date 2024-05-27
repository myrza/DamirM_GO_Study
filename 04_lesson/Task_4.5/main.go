/*
Необходимо создать и вызвать функцию «hello», которая выводит «Hello, Go!» в stdout.
Также, используя defer необходимо вывести фразу «завершение работы».
*/
package main

import "fmt"

func hello() {
	fmt.Println("Hello, Go!")
}
func main() {
	defer fmt.Print("Завершение работы")
	hello()
}
