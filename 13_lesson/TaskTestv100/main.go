/*
	github.com/myrza/patients@v1.0.0 из json в json без сортировки
	github.com/myrza/patients@v1.1.0 из json в json c сортировкой по годам
	github.com/myrza/patients@v2.0.0 из json в xml без сортировки
	github.com/myrza/patients@v2.1.0 из json в xml c сортировкой по годам
}
*/

package main

import (
	"fmt"

	v "github.com/myrza/patients"
)

func main() {
	src := "patients.txt"
	tgt := "res"
	err := v.Do(src, tgt)
	if err == nil {
		fmt.Println("Файл записан успешно")
	} else {
		fmt.Println(err)
	}

}
