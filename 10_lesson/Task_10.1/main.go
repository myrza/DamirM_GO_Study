/*
Необходимо создать карту с ключами: «слон», «бегемот», «носорог», «лев»;
В качестве значений нужно использовать пустые структуры.
Результат вывести в консоль.
*/

package main

import (
	"first"
	"fmt"
)

func main() {

	message := fmt.Sprintf("Hello, my %s package!", first.Hello())
	println(message)
}
