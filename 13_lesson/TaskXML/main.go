/*

}
*/

package main

import (
	//"encoding/json"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"sort"
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
type patients struct {
	List []xml_patient `xml:"Patient"`
}

func main() {
	fmt.Println("Начало работы ")
	f, err := os.Open("patients.txt")
	if err != nil {

		log.Fatalln(f)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	res := make([]paitent, 0, 3)

	for dec.More() {
		var p paitent
		err := dec.Decode(&p)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, p)

	}
	// отсортируем по годам

	sort.Slice(res[:], func(i, j int) bool {
		return res[i].Age < res[j].Age
	})

	fmt.Printf("res: %+v \n", res)
	fmt.Println(res[1].Name)

	fileName := "patients.xml"
	f, err = os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	f.WriteString(xml.Header)

	x1 := xml_patient{
		Name:  res[0].Name,
		Age:   res[0].Age,
		Email: res[0].Email,
	}
	x2 := xml_patient{
		Name:  res[1].Name,
		Age:   res[1].Age,
		Email: res[1].Email,
	}
	x3 := xml_patient{
		Name:  res[2].Name,
		Age:   res[2].Age,
		Email: res[2].Email,
	}
	d := patients{}
	d.List = append(d.List, x1)
	d.List = append(d.List, x2)
	d.List = append(d.List, x3)

	fmt.Printf("%+v", d)

	enc := xml.NewEncoder(f)
	enc.Indent("", "    ")
	err = enc.Encode(d)
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
	f, err = os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	//res1 := xml_patients{}
	err = xml.NewDecoder(f).Decode(&d)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v", d)
	fmt.Println()

	/*
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
	   enc.Indent("", " ")
	   err = enc.Encode(res)

	   	if err != nil {
	   		log.Fatalln(err)
	   	}

	   f.Close()
	*/
}
