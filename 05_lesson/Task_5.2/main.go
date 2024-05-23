// Необходимо создать переменную. В консоль вывести её значение и адрес
package main

import "fmt"

func main() {

	topic := "Строковое значение"
	topicPointer := &topic

	fmt.Printf("Значение: \"%s\". Указатель: %v \n", topic, topicPointer)
}
