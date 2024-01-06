package main

import "fmt"

//Mohamad Rizki Adiansyah
func main() {
	var gajiSekarang, ekspektasiGaji int

	fmt.Print("Masukkan gaji anda : ")
	fmt.Scan(&gajiSekarang)

	fmt.Print("Msasukkan gaji yang anda inginkan : ")
	fmt.Scan(&ekspektasiGaji)

	naikkanGaji(&gajiSekarang, ekspektasiGaji)
	fmt.Printf("\nGaji anda sekarang %d\n", gajiSekarang)

}

func naikkanGaji(gajiAwal *int, gajiAkhir int) {
	*gajiAwal = gajiAkhir
}
