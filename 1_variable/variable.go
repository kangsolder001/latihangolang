package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	// var <nama-variable> <tipe-data>  <-- inisialisasi variable
	var fristName string = "Sendy"  // <-- sudah ada variable fristName dengan tipe data string dan datanya "Sendy"
	var lastName string = "Sandoro" // <-- sudah ada variable fristName dengan tipe data string dan datanya masih kosong

	var (
		fName string
		Lname string
	)

	fName, Lname = "Sendy", "Sandoro"

	namaDepan := "Sendy"      // <-- sudah ada variable fristName dengan tipe data string dan datanya "Sendy"
	namaBelakang := "Sandoro" //<-- sudah ada variable fristName dengan tipe data string dan datanya "Sandoro"
	var age uint8 = 27
	fmt.Println(fristName, lastName)                                                       // cara 1
	text := fmt.Sprintf("My name is %s %s, and my age is %d \n", fristName, lastName, age) // outputan dari Sprintf adalah String ( text )
	fmt.Print(text)
	age = 199
	fmt.Printf("My name is %s %s, and my age is %d \n", fristName, lastName, age) // %s = string , %d = digit ( int, float, int32, int64)
	fmt.Println(fName, Lname)                                                     // cara 2
	fmt.Println(namaDepan, namaBelakang, "adalah nama saya")                      // cara3
	fmt.Println("Tipe data dari age ", reflect.TypeOf(age))
	fmt.Println("Tipe Data dari namaDepan ", reflect.TypeOf(namaDepan))
	fmt.Println("------------------------------------------------------------------")
	//tipe data

	// uint8 -- > u = unsigned, int = integer, 8 = 8 bit -- > unsigned integer 8 bit yang artinya nilainya dimuali dari 0 sampai 2^8 (255)
	// uint8 -- > 0 - 255
	// int8  -- > int = integer, 8 = 8 bit --> nilainya adalah bilangan bulat dari -127 sampai +128

	// Condition IF ELSE

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input Your Point : ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	var (
		point int
		score string
		err   error
	)
	point, err = strconv.Atoi(text)

	if err != nil {
		fmt.Println(err)
	}

	if point > 8 {
		score = "A"
	} else if point > 6 && point <= 8 { // && = and , || = or
		score = "B"
	} else if point <= 6 {
		score = "C"
	}
	fmt.Println("if Else Condition ")
	fmt.Printf("Your Score is %s, and your point is %d \n", score, point)
	fmt.Println("---------------------")

	// Condition Switch Case
	// falltroungh --> untuk meneruskan ke case selanjutnya , meskipun case sudah terpenuhi
	switch {
	case point > 8: // case 1
		score = "A"
	case point > 3 && point <= 8: // case 2
		score = "B"
		fallthrough
	case point <= 6 && point > 0: // case 3
		score = "C"

	default:
		{
			score = "not found"
		}
	}
	fmt.Println("Switch Case Condition ")
	fmt.Printf("Your Score is %s, and your point is %d \n", score, point)
	fmt.Println("---------------------")

}
