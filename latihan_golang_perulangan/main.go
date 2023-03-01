package main

import (
	"fmt"
)

func main() {
	fmt.Println("For dengan biasa")
	// for <nilai_awal>;<nilai_akhir>;<pertambahan_nilai>
	for i := 0; i <= 5; i++ { // perulangan biasa ( i++ = i = i+1)
		fmt.Printf("number : %d \n", i)
	}
	fmt.Println("-------------------------------------")
	fmt.Println("For dengan break")
	i := 0

	for {
		fmt.Printf("number : %d \n", i)
		if i == 5 {
			break // break untuk memberhentikan for
		}
		i++
	}
	fmt.Println("-------------------------------------")
	fmt.Println("For dengan kondisi")
	i = 0
	// for <kondisi>
	for i <= 5 {
		fmt.Printf("number : %d \n", i)
		i++
	}

	//for dalam for
	for i := 0; i <= 5; i++ {
		fmt.Printf("Number i = %d . ", i)
		for j := 0; j <= i; j++ {

			if j == 3 {
				fmt.Print("*")
			} else {
				fmt.Printf("%d", j)
			}
		}
		fmt.Println("")

	}
}
