/*
Необходимо создать модуль и залить в публичный репозиторий на https://github.com/.
Модуль должен иметь функцию Hello.
Используя теги git, необходимо создать пару версий модуля: v1.0.0 и v1.1.0.
Разные версии должны выводить разные сообщения: «Hello, v1.0.0» и «Hello, v1.1.0» соответсвенно.
Написать программу, которая будет одновременно использовать логику версий v1.0.0 и v1.1.0.
Результат должен выводиться в консоль.
*/

package main

import (
	"fmt"

	v "github.com/myrza/hello"
	v2 "github.com/myrza/hello/v2"
)

func main() {

	fmt.Println("module_hello diferent version")
	v2.Version()
	v.Version()
}
