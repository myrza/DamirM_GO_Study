/*
Необходимо создать переменную и указатель на неё. Необходимо через указатель изменить значение переменной
*/
package main

import (
	"fmt"
)

func main() {

	topic := "Строковое значение"
	topicPointer := &topic
	fmt.Printf("Значение: \"%s\". Указатель: %v \n", topic, topicPointer)

	// меняем значение через указатель
	*topicPointer = "Значение изменилось"
	fmt.Printf("Значение: \"%s\". Указатель: %v \n", topic, topicPointer)
}
