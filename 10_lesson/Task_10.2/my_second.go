package second

import (
	"first"
	"fmt"
)

const name = "second"

func init() {
	fmt.Println("Init second!")
}
func Hello() string {
	return name
}

func HelloFromFirst() string {
	return first.Hello()
}
