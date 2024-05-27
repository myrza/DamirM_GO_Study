/*
Необходимо создать карту с элементами: «слон» = 3, «бегемот» = 0, «носорог» = 5, «лев» = 1.
Изменить «бегемот» на 2 и вывести результат в консоль
*/
package main

import "fmt"

func main() {

	m := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	fmt.Println("Было бегемотов:", m["бегемот"])
	m["бегемот"] = 2
	fmt.Println("Стало бегемотов:", m["бегемот"])
}
