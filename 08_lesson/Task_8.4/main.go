/*
Необходимо создать карту с ключами: «слон», «бегемот», «носорог», «лев»;
В качестве значений нужно использовать пустые структуры.
Далее добавить «выдра» и результат вывести в консоль.
*/
package main

import "fmt"

type empty_st struct{}

func main() {
	animals := map[string]empty_st{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	animals["выдра"] = empty_st{}

	fmt.Println(animals)

}
