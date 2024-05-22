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
