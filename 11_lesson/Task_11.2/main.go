/*
Нужно создать, используя оборачивание, ошибку «ошибка3:ошибка2:ошибка1».
Из созданной цепочки ошибок нужно получить ошибку «ошибка2:ошибка1» и вывести в stdout
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("Ошибка1")
	err2 := fmt.Errorf("Ошибка2:%w", err1)
	err3 := fmt.Errorf("Ошибка3:%w", err2)

	fmt.Println(errors.Unwrap(err3))

}
