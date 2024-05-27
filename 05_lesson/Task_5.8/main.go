/*
Необходимо создать пользовательский тип square на базе int.
Далее необходимо создать переменную s типа square и значением 34.
Значение переменной s необходимо увеличить на 10 и вывести результат в консоль.
Тип square при выводе в консоль должен автоматически дописывать м², то есть результат должен выглядеть: 44 м².
*/
package main

import "fmt"

type Square int

func String(s Square) string {
	return fmt.Sprintf("%d м²", s)
}
func main() {
	var s Square = 34
	s += 10
	fmt.Println("Результат: ", String(s))

}
