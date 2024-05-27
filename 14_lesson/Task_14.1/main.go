/*
Необходимо создать дочернюю горутину, которая выведет в stdout «Привет из дочерней горутины!»
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Привет из дочерней горутины!")
	}()

	time.Sleep(2 * time.Duration(time.Second))
}
