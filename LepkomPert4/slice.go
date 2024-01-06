package main

import "fmt"

//Mohamamd Rizki Adiansyah
func main() {
	var kursus = []string{"dbms", "server os", "networking", "web", "desktop", "erp"}
	var kursus_saya = kursus[1:5]

	kursus_saya = append(kursus_saya, "Manajemen Informatika")

	fmt.Println("Isi dari kursus adalah", kursus)
	fmt.Println("Panjang kursus", len(kursus), "dan kapasitas", cap(kursus))

	fmt.Println("Isi dari kursus saya adalah", kursus_saya)
	fmt.Println("Panjang kursus", len(kursus_saya), "dan kapasitas", cap(kursus_saya))

}
