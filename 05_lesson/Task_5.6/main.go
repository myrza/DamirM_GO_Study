// Необходимо создать пользовательский тип square на базе int.
// Далее необходимо создать переменную s типа square и значением 25.
// Необходимо вывести значение переменной в консоль
package main

import "fmt"

type Square int

func main() {
	s := Square(25)
	fmt.Print(s)
}
