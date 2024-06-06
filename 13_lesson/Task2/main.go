/*

}
*/

package main

import (
	//"encoding/json"
	"encoding/json"
	"io/ioutil"
	"log"
)

type patient struct {
	Name  string `json:"name"`
	Age   int    `"json:"age"`
	Email string `"json:"email"`
}

func main() {
	c := patient{
		Name:  "косой",
		Age:   16,
		Email: "dom60m@gmail.com",
	}
	f, err := ioutil.TempFile("./", "sample-")
	if err != nil {
		log.Fatalln(err)

	}
	err = json.NewEncoder(f).Encode(c)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("файл записан")

}
