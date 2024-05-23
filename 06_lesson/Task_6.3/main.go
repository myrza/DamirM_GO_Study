/*
Необходимо объявить глобальную структуру contract с полями: ID int, Number string, Date string.
Далее создать экземпляр структуры со значениями полей: ID=1, Number=«#000A\n101», Date=«2024-01-31».
При передачи экземпляра структуры в fmt.Println в консоли должно отображаться: Договор № #000A\n101 от 2024-01-31
*/
package main

import (
	"fmt"
)

type Contract struct {
	ID        int
	Number    string
	Date      string
	OutputRes string
}

func (c *Contract) format() {
	c.OutputRes = "Договор № " + c.Number + " от " + c.Date
}
func main() {

	cntr := Contract{ID: 1, Number: "#000A\\n101", Date: "2024-01-31"}
	cntr.format()
	fmt.Println(cntr.OutputRes)
}
