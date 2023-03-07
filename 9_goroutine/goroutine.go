package main

import (
	. "fmt"
)

func main() {
	Println("GoRoutine")

	var input string
	go Scanln(&input)
	Println("Text text text")
	Printf("Text : %s \n", input)
}
