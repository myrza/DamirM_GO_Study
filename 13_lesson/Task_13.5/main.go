/*
В рамках задачи будем работать с картотекой известного врача. Нужно будет написать модуль с несколькими версиями: v1.0.0, v1.1.0, v2.0.0, v2.1.0. Модуль должен прочитать файл со следующим содержимым:

{"name":"Ёжик","age":10,"email":"ezh@mail.ru"}
{"name":"Зайчик","age":2,"email":"zayac@mail.ru"}
{"name":"Лисичка","age":3,"email":"alice@mail.ru"}

v1.0.0 должна создавать файл с содержимым:

[{«name»:»Ёжик","age":10,"email":"ezh@mail.ru"},
{"name":"Зайчик","age":2,"email":"zayac@mail.ru"},
{«name":"Лисичка","age":3,"email":"alice@mail.ru"}]

v1.1.0 должна сортировать данные по полю age по
возрастанию:
[{«name":"Зайчик","age":2,"email":"zayac@mail.ru"},
{«name":"Лисичка","age":3,"email":"alice@mail.ru"}{«name»:»
Ёжик","age":10,"email":"ezh@mail.ru"}]

v2.0.0 должна создавать файл с содержимым:
<?xml version="1.0" encoding="UTF-8"?>
<patients>
<Patient>
<Name>Ёжик</Name>
<Age>10</Age>
<Email>ezh@mail.ru</Email>
</Patient>
<Patient>
<Name>Зайчик</Name>
<Age>2</Age>
<Email>zayac@mail.ru</Email>
</Patient>
<Patient>
<Name>Лисичка</Name>
<Age>3</Age>
<Email>alice@mail.ru</Email>
</Patient>
</patients>
v2.1.0 должна сортировать данные по полю age по возрастанию:
<?xml version="1.0" encoding="UTF-8"?>
<patients>
<Patient>
<Name>Зайчик</Name>
<Age>2</Age>
<Email>zayac@mail.ru</Email>
</Patient>
<Patient>
<Name>Лисичка</Name>
<Age>3</Age>
<Email>alice@mail.ru</Email>
</Patient>
<Patient>
<Name>Ёжик</Name>
<Age>10</Age>
<Email>ezh@mail.ru</Email>
</Patient>
</patients>
Модуль должен содержать функцию Do, которая принимает два
строковых параметра: путь файла откуда прочитать данные, путь
файла, в который записать данные в требуемом формате; и
возвращать ошибку. Пример использования модуля:
package main
import (
format ...
)
func main() {
err := format.Do("patients", «result")
if err != nil {
…
}
}
Должна быть возможность подключить любую из версий: v1.0.0, v1.1.0,
v2.0.0, v2.1.0.

}
*/

package main

import (
	"encoding/xml"
	"fmt"
)

type Contractor interface {
}
type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	c1 := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	}
	c2 := contract{
		Number:   2,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	}
	d := contracts{}
	d.List = append(d.List, c1)
	d.List = append(d.List, c2)

	res, err := xml.Marshal(d)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(xml.Header, string(res))
	fmt.Println()
}
