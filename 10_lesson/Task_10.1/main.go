/*
Необходимо создать и установить пакет first. Пакет должен иметь функцию Hello, которая возвращает строку «Hello, Dolly!».
Далее написать программу, которая будет вызывать функцию Hello пакета first и выводить результат в консоль
*/

package main

import (
	"first"
	"fmt"
)

func main() {

	fmt.Println(first.Hello())
}
