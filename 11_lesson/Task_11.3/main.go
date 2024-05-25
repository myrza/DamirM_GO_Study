/*
Нужно создать, используя оборачивание, ошибку «ошибка3:ошибка2:ошибка1».
Не используя unwrap, нужно определить была ли ошибка «ошибка1»
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("Ошибка1")
	err2 := fmt.Errorf("Ошибка2:%w", err1)
	err3 := fmt.Errorf("Ошибка2:%w", err2)

	if errors.Is(err3, err1) {
		fmt.Printf("%v была", err1)
	} else {
		fmt.Printf("%v - не было", err1)
	}

}
