/*
Необходимо создать и установить пакет second. Пакет должен иметь функцию Hello, которая возвращает строку «Hello, Louis!».
Далее написать программу, которая будет вызывать функцию Hello пакета first задачи 8.1 и функцию Hello пакета second.
Оба результата должны выводиться в консоль
*/

package main

import (
	"first"
	"fmt"
	"second"
)

func main() {
	fmt.Println(first.Hello())
	fmt.Println(second.Hello())
	//message := fmt.Sprintf("Hello, my %s package!", )
	//println(message)
}
