/*
Необходимо распарсить json {"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов»}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int
	Landlord string
	Tenat    string
}

func main() {

	str := string(`{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`)
	c := contract{}
	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v", c)
	fmt.Println()
}
