/*
Необходимо представить в виде json структуру contract

contract{
Number: 2,
Landlord: "Остап",
tenat: "Шура",
}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	Tenat    string `json:"tenat"`
}

func main() {
	c := contract{
		Number:   1,
		Landlord: "Остап",
		Tenat:    "Шура",
	}

	res, err := json.Marshal(c)

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v", string(res))

}
