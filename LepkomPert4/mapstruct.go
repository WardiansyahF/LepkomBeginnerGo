package main

import "fmt"

//Mohamad Rizki Adiansyah
type mahasiswa struct {
	Nama  string
	Kelas string
}

func main() {
	var data = map[string]mahasiswa{
		"50421836": {
			"Mohamad Rizki Adiansyah",
			"2IA02",
		},
		"10115966": {
			"Ariel Noah",
			"4IA20",
		},
	}
	var search string
	fmt.Print("Masukkan NPM anda ? ")
	fmt.Scanf("%s", &search)

	var value, ok = data[search]

	if ok {
		fmt.Printf("Nama anda %s \nKelas anda %s", value.Nama, value.Kelas)
	} else {
		fmt.Println("data tidak ditemukan")
	}
}
