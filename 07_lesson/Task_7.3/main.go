/*
Необходимо создать массив, содержащий значения: «яблоко», «груша», «помидор», «абрикос».
Далее «помидор» заменить на «персик».
Результат вывести в консоль
*/
package main

import "fmt"

func main() {

	a := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	fmt.Println("Было: ", a)
	// меняем помидор на персик
	a[2] = "персик"
	fmt.Println("Стало: ", a)
}
