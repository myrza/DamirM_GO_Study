// Необходимо создать указатель на строковое значение
package main

import "fmt"

func main() {

	topic := "Строковое значение"
	topicPointer := &topic

	fmt.Printf("Значение: \"%s\". Указатель: %v \n", topic, topicPointer)
}
