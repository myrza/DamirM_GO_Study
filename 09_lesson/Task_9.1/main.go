/*
Необходимо написать функцию fruitMarket, которая будет принимать название фруктов, а возвращать их количество.
Например, «апельсины» - 5.
Сама функция должна создать карту: апельсин=5, яблоки=3, сливы=1, груши=0.
Если запрашиваемых фруктов нет в карте, то в консоль должно выводится сообщение об отсутствии
*/

package main

import "fmt"

func fruitMarket(fn string) int {
	fmap := map[string]int{"апельсин": 5, "яблоки": 3, "сливы": 1, "груши": 0}

	val, ok := fmap[fn]
	if ok == false {
		fmt.Printf("Фрукт %s не обнаружен\n", fn)
	}

	return val
}

func main() {
	fn := "апельсин1"
	//cnt := fruitMarket(fn)
	fmt.Printf("%s всего:%v \n", fn, fruitMarket(fn))
}
