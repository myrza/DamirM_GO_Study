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

type user struct {
	ID                 int
	Name, Addss, Phone string
}

type employee struct {
	ID int
	user
}

func main() {
	u := user{ID: 1, Name: "Вася", Addss: "Москва, 3 улица Строителей 2", Phone: "911-111-111"}
	e := employee{ID: 1, user: user{ID: 1, Name: "Пупкин", Addss: "Москва, Марьина роща 2", Phone: "777-222-222"}}

	fmt.Println("Пользователь: ", u.Name, u.Addss, u.Phone, " \nСотрудник: ", e.Name, e.Addss, e.Phone)
}
