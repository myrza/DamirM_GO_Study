/*
Необходимо создать модуль и залить в публичный репозиторий на https://github.com/.
Модуль должен иметь функцию Hello. Используя теги git, необходимо создать пару версий модуля: v1.0.0 и v1.1.0.
Разные версии должны выводить разные сообщения: «Hello, v1.0.0» и «Hello, v1.1.0» соответсвенно.
Написать программу, которая будет одновременно использовать логику версий v1.0.0 и v1.1.0.
Результат должен выводиться в консоль.
*/

package main

import (
	"first"
	"fmt"
	"second"
)

func main() {
	fmt.Println(first.Hello())
	fmt.Println(second.Hello())
	//message := fmt.Sprintf("Hello, my %s package!", )
	//println(message)
}
