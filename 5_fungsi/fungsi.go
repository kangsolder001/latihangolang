package main

import (
	"fmt"
)

func main() {
	angka_pertama := 1
	angka_kedua := 2

	fmt.Println("function single return")
	hasil_penjumlahan := Hitung_penjumlahan(angka_pertama, angka_kedua)
	fmt.Printf("Hasil Penjumlahan dari angka %d + %d = %d \n", angka_pertama, angka_kedua, hasil_penjumlahan)
	hasil_pengurangan := Hitung_pengurangan(angka_pertama, angka_kedua)
	fmt.Printf("Hasil Pengurangan dari angka %d + %d = %d \n", angka_pertama, angka_kedua, hasil_pengurangan)

	//func multiple return
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("function multiple return ")
	hasil_penjumlahan, hasil_pengurangan = Hitung_jumlah_kurang(angka_pertama, angka_kedua)

	fmt.Printf("Hasil Penjumlahan dari angka %d + %d = %d \n", angka_pertama, angka_kedua, hasil_penjumlahan)
	fmt.Printf("Hasil Pengurangan dari angka %d + %d = %d \n", angka_pertama, angka_kedua, hasil_pengurangan)

	fmt.Println("------------------------------Fucntion Veriadict---------------------------")
	hasil_veriadict := Calculate(1, 2, 3, 4, 5)
	fmt.Printf("Hasil dari penjumlahan angka diatas adalah : %d", hasil_veriadict)

}

// func <nama_func> (<input_data>) <output_data> {}

// return = A + B
func Hitung_penjumlahan(angka_pertama, angka_kedua int) int {
	var (
		hasil int
	)
	hasil = angka_pertama + angka_kedua
	return hasil
}

// return = A - B
func Hitung_pengurangan(angka_pertama, angka_kedua int) int {
	var (
		hasil int
	)
	hasil = angka_pertama - angka_kedua
	return hasil
}

// return A+B, A-B

func Hitung_jumlah_kurang(angka_pertama, angka_kedua int) (int, int) {
	var (
		hasil_1 int
		hasil_2 int
	)
	hasil_1 = angka_pertama + angka_kedua
	hasil_2 = angka_pertama - angka_kedua

	return hasil_1, hasil_2
}

// function veriadict

func Calculate(angka ...int) int {
	var (
		hasil int
	)
	for _, value := range angka {
		hasil = hasil + value
	}
	return hasil
}
