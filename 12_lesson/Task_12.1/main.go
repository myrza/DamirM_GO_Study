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

func main() {

	a := 1
	do(a)
}

func do(v any) {
	a := 10
}
