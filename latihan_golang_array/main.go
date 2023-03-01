package main

import "fmt"

func main() {

	var bilanganBulat [5]int
	bilanganBulat[0] = 0
	bilanganBulat[1] = 1
	bilanganBulat[2] = 2
	bilanganBulat[3] = 3
	bilanganBulat[4] = 4

	var bilanganArrayKeDua = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for untuk array
	fmt.Println("Array pertama")
	for index, element := range bilanganBulat { // i = index_array , nilai = element_array
		fmt.Printf("Index %d bernilai %d \n", index, element)
	}
	fmt.Println("Array kedua")
	for index, element := range bilanganArrayKeDua { // i = index_array , nilai = element_array
		fmt.Printf("Index %d bernilai %d \n", index, element)
	}
	// array 2D
	var array2d = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(array2d[0][1])

	// penggunaan for dalam for
	for i := 0; i < len(array2d); i++ {
		for j := 0; j < len(array2d); j++ {
			fmt.Printf("%d,", array2d[i][j])
		}
		fmt.Println()
	}

	var fruits = []string{"apple", "grape", "banana", "melon"} // data awal

	var slice = fruits[1:2] // data slice  --> grape
	fmt.Println(fruits)     //--> print semua fruits
	fmt.Println(slice)
	slice[0] = "salak"  // --> si grape diganti dengan salak
	fmt.Println(fruits) // --> print semua fruits

	fmt.Printf("CAP fruit %d , LEN fruit %d \n", cap(fruits), len(fruits))
	fmt.Printf("CAP slice %d , LEN Slice %d \n", cap(slice), len(slice))
}
