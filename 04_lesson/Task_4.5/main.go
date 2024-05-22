package main

import "fmt"

func hello() {
	fmt.Println("Hello, Go!")
}
func main() {
	defer fmt.Print("Завершение работы")
	hello()
}
