/*

}
*/

package main

import (
	//"encoding/json"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type paitent struct {
	Name  string `json:"name"`
	Age   int    `"json:"age"`
	Email string `"json:"email"`
}

type xml_patient struct {
	Name  string `xml:"Name"`
	Age   int    `xml:"Age"`
	Email string `xml:"Email"`
}
type xml_patients struct {
	List []xml_patient `xml:"patient"`
}

func Do(srcFile string, tgtFaile string) error {
	fmt.Println(srcFile, tgtFaile)
	srcF, err := os.Open("patients")
	if err != nil {
		return errors.New("Ошибка при попытке прочитать файл!")
	}
	defer srcF.Close()
	fmt.Printf("%+v", srcF)

	dec := json.NewDecoder(srcF)
	res := make([]paitent, 0, 3)

	for dec.More() {
		var p paitent
		err := dec.Decode(&p)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, p)

	}
	fmt.Println("read file")
	tgtF, err := ioutil.TempFile("./", tgtFaile)
	if err != nil {
		return errors.New("Не удалось создать файл для записи результата!")
	}

	err = json.NewEncoder(tgtF).Encode(res)
	if err != nil {
		return errors.New("Не удалось записать результат в файл")
	}
	tgtF.Close()

	return nil

}
func main() {
	fmt.Println("Начало работы ")
	teststrin := "patients.txt"
	testerr := Do(teststrin, "restest")
	fmt.Println(testerr)

	f, err := os.Open(teststrin) //os.Open("patients.txt")
	if err != nil {

		fmt.Printf("ошибка в файле: %+v \n", f)
		fmt.Println("Ошибка", err)
		log.Fatalln(f)
	}
	defer f.Close()

	/*dec := json.NewDecoder(f)
	res := make([]paitent, 0, 3)

	for dec.More() {
		var p paitent
		err := dec.Decode(&p)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, p)

	}
	fmt.Printf("res: %+v \n", res)

	f, err = ioutil.TempFile("./", "sample-")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.NewEncoder(f).Encode(res)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(f, err)
	f.Close()

	log.Println("файл записан")

	//===========================================
	// отсортируем по годам

	sort.Slice(res[:], func(i, j int) bool {
		return res[i].Age < res[j].Age
	})

	f, err = ioutil.TempFile("./", "sample2-")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.NewEncoder(f).Encode(res)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(f, err)
	f.Close()
	//=================================================================//
	// записать в файл xml!!! последнее!! потом будем собирать модуль и подключать их//
	fileName := "patients.xml"
	f, err = os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	f.WriteString(xml.Header)
	enc := xml.NewEncoder(f)
	for dec.More() {
		var p paitent
		err := dec.Decode(&p)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, p)

	}
	enc.Indent("", " ")
	err = enc.Encode(res)
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
	*/
}
