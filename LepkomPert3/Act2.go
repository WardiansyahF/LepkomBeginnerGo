package main

import (
	"fmt"
)

var NilaiUts, NilaiUas, NilaiTotal float64

func main() {
	fmt.Print("Masukkan Nilai UTS: ")
	fmt.Scan(&NilaiUts)
	fmt.Print("Masukkan Nilai UAS: ")
	fmt.Scan(&NilaiUas)

	NilaiTotal := NilaiUts*0.3 + NilaiUas*0.7

	fmt.Println("Nilai UTS:", NilaiUts)
	fmt.Println("Nilai UAS:", NilaiUas)
	fmt.Println("Nilai Total Anda:", NilaiTotal)

	if NilaiTotal <= 20 {
		fmt.Println("Grade Anda E")
	} else if NilaiTotal >= 21 && NilaiTotal <= 40 {
		fmt.Println("Grade Anda D")
	} else if NilaiTotal >= 41 && NilaiTotal <= 60 {
		fmt.Println("Grade Anda C")
	} else if NilaiTotal >= 61 && NilaiTotal <= 80 {
		fmt.Println("Grade Anda B")
	} else {
		fmt.Println("Grade Anda A")
	}

}
