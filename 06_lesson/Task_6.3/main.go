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
	ID     int
	Number string
	Date   string
}

func String(c Contract) string {
	return fmt.Sprintf("Договор № %v от %v ", c.Number, c.Date)
}

func main() {

	cntr := Contract{ID: 1, Number: "#000A\\n101", Date: "2024-01-31"}
	fmt.Println(String(cntr))
}
