package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Jude"
	course := "Somesome"
	module := 3.3

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
}

