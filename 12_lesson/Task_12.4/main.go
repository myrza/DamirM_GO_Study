/*
В представленном ниже коде возникает паника. Нужно понять где и объяснить почему.
Также нужно предложить не менее 3-х вариантов решения проблемы с объяснением причины устранения ошибки
*/
/*
	Анализ ошибки: Ошибка возникает при вызове функции Sing. На вход требуется передавать интерфейс
*/
package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}
type Duck struct {
	voice string
}

func (d Duck) Sing() string {
	return d.voice
}

func main() {
	//var d *Duck

	var d Duck
	d = Duck{"кря-кря"}

	// 1 решение: решение: получаем голос из структуры
	fmt.Println("1 решение: Получаем голос из структуры d.voice:", d.voice)

	// 2 решение: Новая функция обрабатывающая тип
	fmt.Println("2 решение: Новая функция обрабатывающая тип SingNew(d):", SingNew(d))

	// 3 решение: Создаем интерфейс и передаем в изначальную функцию
	var b Bird
	b = d
	song, err := Sing(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("3 решение: Создаем интерфейс и передаем в изначальную функцию :", song, err)

}
func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
func SingNew(d Duck) string {
	return d.voice

}
