/*
Не изменяя структуры Xml, Csv и функцию main необходимо
доработать следующий код так, чтобы в консоли увидели:
Данные в формате xml
Данные в формате csv

*/

package main

import "fmt"

type Formater interface {
	Format()
}
type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}
func main() {

	var x, c Formater

	x = Xml{}
	Report(x)

	c = Csv{}
	Report(c)

}
func Report(x Formater) {
	x.Format()
}
