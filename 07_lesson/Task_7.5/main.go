/*
Необходимо создать пустой срез ёмкостью 10.
Результат вывести в консоль
*/
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0, 10)

	fmt.Println("Длина:", len(a), " Емкость:", cap(a))
}
