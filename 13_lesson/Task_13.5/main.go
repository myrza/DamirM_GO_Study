/*
Необходимо представить в виде xml структуру contracts

type contracts struct {
List []contract `xml:"contract"`
}
contract{
Number: 1,
Landlord: "Остап Бендер",
tenat: "Шура Балаганов",

}
*/

package main

import (
	"encoding/xml"
	"fmt"
)

type Contractor interface {
}
type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	c1 := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	}
	c2 := contract{
		Number:   2,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	}
	d := contracts{}
	d.List = append(d.List, c1)
	d.List = append(d.List, c2)

	res, err := xml.Marshal(d)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(xml.Header, string(res))
	fmt.Println()
}
