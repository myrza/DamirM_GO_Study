/*
Необходимо создать модуль и залить в публичный репозиторий на https://github.com/.
Модуль должен иметь функцию Hello. Используя теги git, необходимо создать пару версий модуля: v1.0.0 и v1.1.0.
Разные версии должны выводить разные сообщения: «C:\Repo\DamirM_GO_Study1\10_lesson\Task_10.4\main.go» и «Hello, v1.1.0» соответсвенно.
Написать программу, которая будет одновременно использовать логику версий v1.0.0 и v1.1.0.
Результат должен выводиться в консоль.
*/

package main

import (
	"fmt"

	v100 "github.com/myrza/module_hello/v100"
	v110 "github.com/myrza/module_hello/v110"
)

func main() {

	fmt.Println("module_hello diferent version")
	fmt.Println(v100.Version())
	fmt.Println(v110.Version())