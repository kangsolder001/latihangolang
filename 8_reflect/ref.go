package main

import (
	"fmt"
	ref "reflect"
)

func main() {
	numb := 8
	refV := ref.ValueOf(numb)

	fmt.Println(refV.Type())
	fmt.Println(ref.TypeOf(numb))
}
