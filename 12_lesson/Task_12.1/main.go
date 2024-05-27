/*
package main
import "fmt"
func main() {
a := 1
do(a)
}
func do(v any) {
a := 10

Комментарий:
Здесь нужно увеличить значение a на v.
В случае невозможности приведения к int
необходимо сообщить об этом и немедленно
завершить полнение программы.

*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 1

	fmt.Println("1 кейс: передаем в функцию целое число")
	do(a)
	fmt.Println("2 кейс: передаем в функцию строку. Обрабатываем панику")
	do("test")
}
func do(v any) {
	defer func() {
		if recover() != nil {
			err := errors.New("Паника! Входной параметр не целое число")
			fmt.Println(err)
			return
		}
	}()
	a := 10
	val := v.(int)
	a += val
	fmt.Println(a)

}
