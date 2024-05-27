/*
Необходимо создать карту с ключами: «слон», «бегемот», «носорог», «лев»;
В качестве значений нужно использовать пустые структуры.
Результат вывести в консоль.
*/

package main

import "fmt"

type Animal struct {
	name string
}

func main() {
	a := Animal{}
	m := map[string]Animal{
		"слон":    a,
		"бегемот": a,
		"носорог": a,
		"лев":     a,
	}

	fmt.Print(m)

}
