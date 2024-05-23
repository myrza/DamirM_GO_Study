// Необходимо объявить глобальную структуру contract с полями: ID int, Number string, Date string.
// Далее создать экземпляр структуры со значениями полей:
// ID=1, Number=«#000A\n101», Date=«2024-01-31».
// В консоль нужно вывести структуру таким образом, чтобы данные отображались в виде:
// {ID:1Number:#000A\n101 Date:2024-01-31}

package main

import "fmt"

type Contract struct {
	ID     int
	Number string
	Date   string
}

func main() {

	cntr := Contract{ID: 1, Number: "#000A\n101", Date: "2024-01-31"}

	fmt.Printf("ID:%v Number:%#v Date:%v", cntr.ID, cntr.Number, cntr.Date)

}
