package main

import "fmt"

func main() {
	func() {
		fmt.Print("Hello, Go!")
	}()

}
