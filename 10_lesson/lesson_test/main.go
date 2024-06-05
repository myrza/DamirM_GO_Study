/*
тест переадресации
*/

package main

import (
	v100 "github.com/myrza/module_hello5/v2@v1.0.0"
	v110 "github.com/myrza/module_hello5@v1.1.0"
)

func main() {

	v110.Version()
	v100.Version()
}
