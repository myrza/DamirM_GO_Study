/*
Необходимо убрать повторяющийся код - поля Addss и Phone из структур:
type user struct {
ID int
Name string
Addss string
Phone string
}
type employee struct {
ID int
Name string
Addss string
Phone string
}

После проведения рефакторинга строка fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone) должна выводить в консоль
адрес и телефон пользователя и сотрудника соответственно
*/
package main

import (
	"fmt"
)

type contact struct {
	Addss, Phone string
}
type user struct {
	ID   int
	Name string
	contact
}

type employee struct {
	ID   int
	Name string
	contact
}

func main() {
	u := user{ID: 1, Name: "Вася", contact: contact{"Москва, 3 улица Строителей 2;", "911-111-111"}}
	e := employee{ID: 1, Name: "Никанор Босой", contact: contact{"Москва, Марьина роща 2;", "777-222-222"}}

	fmt.Println("Пользователь: ", u.Addss, u.Phone, " \nСотрудник: ", e.Addss, e.Phone)
}
