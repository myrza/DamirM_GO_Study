/*
Необходимо создать карту с ключами: «слон», «бегемот», «носорог», «лев»;
В качестве значений нужно использовать пустые структуры.
Результат вывести в консоль.
*/

package main

import "fmt"

type Animal struct {
	name     string
	nickName string
	age      int
}

func main() {

	m := map[string]Animal{}

	fmt.Print(m)
	/*
			m := map[string]slc{} //пустая карта нулевой длинны


		m["слон"] = nil
		m["бегемот":] = nil
		m["носорог"] = nil
		m["лев"] = nil

			fmt.Println(m, len(m))
	*/
}
