//Необходимо создать пользовательский тип square на базе int. Далее необходимо создать переменную s типа square и значением 30.
//Значение переменной s необходимо увеличить на 15 и вывести
//результат в консоль.

package main

import "fmt"

type Square int

func main() {
	s := Square(30)
	s = s + Square(15)
	fmt.Print(s)
}
