/*
Необходимо создать и вызвать анонимную функцию, которая выводит «Hello, Go!» в stdout.
*/
package main

import "fmt"

func main() {
	func() {
		fmt.Print("Hello, Go!")
	}()

}
