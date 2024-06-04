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

	fmt.Println("Используем пакет  first:", first.Hello())
	fmt.Println("Используем пакет second:", second.Hello())
}
