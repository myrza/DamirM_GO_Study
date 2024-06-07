/*
}
*/
package main

import (
	"fmt"

	v1 "github.com/myrza/patientsmodule"
)

func main() {
	src := "patients.txt"
	tgt := "res"
	err := v1.Do(src, tgt)
	if err == nil {
		fmt.Println("Файл записан успешно")
	} else {
		fmt.Println(err)
	}

}
