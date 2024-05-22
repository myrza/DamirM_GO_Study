package main

import "fmt"

func hello(f func()) func() {
	return f
}

func main() {
	f := func() {
		fmt.Print("Hello, Go!")
	}
	f2 := hello(f)
	f2()

}
