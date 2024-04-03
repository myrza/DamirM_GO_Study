package main

import "fmt"

func main() {
	var firstValue = 16
	var secondValue = 3
	var result = firstValue / secondValue
	var remainder = firstValue % secondValue
	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T", result, remainder, result)
}
