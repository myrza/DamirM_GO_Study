/*
Нужно создать, используя оборачивание, ошибку «ошибка3:ошибка2:ошибка1»
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
	fmt.Println(err3)

}
