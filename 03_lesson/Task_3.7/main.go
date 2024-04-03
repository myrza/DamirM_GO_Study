package main

import "fmt"

const (
	myConst1 = 1 << iota
	myConst2
	myConst3
	myConst4
	myConst5
)

func main() {
	fmt.Printf("Результат: %d, %d, %d, %d, %d", myConst1, myConst2, myConst3, myConst4, myConst5)
}
