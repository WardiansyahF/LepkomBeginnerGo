package main

import "fmt"

var i int

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print("Input Nilai I : ")
		fmt.Scan(&i)
		if i <= 10{
			if i % 2 == 0{
				fmt.Println(i,"adalah Angka Genap")
			}else if i % 2 ==1{
				fmt.Println(i,"adalah Angka Ganjil")
			}else{
				fmt.Println(i,"adalah Angka Nol")
			}
		}else if i > 10{
			fmt.Println("Pertanyaan selesai, karena angka I yang diinput sudah melebihi dari 10. Terimakasih")
			break;
		}
	}
}