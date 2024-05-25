/*
Необходимо создать карту с элементами: «слон» = 3, «бегемот» = 0, «носорог» = 5, «лев» = 1.
Перебрать значения: «слон», «бегемот», «осьминог» и вывести сообщения в консоль по шаблону: «:Животное: %s, количество: %d (есть в карте: %v)»
*/
package main

import "fmt"

func main() {

	animal := map[string]int{"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1}
	//fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)", animal["слон"], animal["слон"], animal)

	val, ok := animal["слон"]
	if ok {
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v) \n", "слон", val, animal)
	} else {
		fmt.Printf("Животное НЕ в карте: %v \n", animal)
	}

	val, ok = animal["бегемот"]
	if ok {
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v) \n", "бегемот", val, animal)
	} else {
		fmt.Printf("Животное НЕ в карте: %v \n", animal)
	}

	val, ok = animal["осьминог"]
	if ok {
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v) \n", "осьминог", val, animal)
	} else {
		fmt.Printf("Животное НЕ в карте: %v \n", animal)
	}
}
