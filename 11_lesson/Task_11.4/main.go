/*
Нужно создать, используя оборачивание, ошибку «ошибка3:ошибка2:ошибка1».
Также нужно создать свою ошибку в виде структуры myFirstError, которая обязательно должна иметь метод Error() string.
Необходимо убедиться, что в созданной цепочке ошибок не было ошибки типа myFirstError
*/

package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	message string
}

func (e myFirstError) Error() string {
	return e.message
}

func main() {
	err1 := errors.New("Ошибка1")
	err2 := fmt.Errorf("Ошибка2:%w", err1)
	err3 := fmt.Errorf("Ошибка2:%w", err2)

	mfe := myFirstError{message: "Моя первая ошибка"}

	if errors.Is(err3, mfe) {
		fmt.Printf("%v была в цепочке ошибок", mfe)
	} else {
		fmt.Printf("%v - не было в цепочке ошибок", mfe)
	}

}
