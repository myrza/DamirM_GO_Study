/*
Необходимо создать карту с элементами: «слон» = 3, «бегемот» = 0, «носорог» = 5, «лев» = 1.
Изменить «бегемот» на 2 и вывести результат в консоль
*/
package main

import (
	"fmt"
)

func main() {
	animal := map[string]int{"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1}
	animal["бегемот"] = 2
	fmt.Print("Бегемот, новое количество:", animal["бегемот"])
}