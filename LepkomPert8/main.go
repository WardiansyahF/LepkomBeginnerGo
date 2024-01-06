package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	http.HandleFunc("/api/mahasiswa", user)
	fmt.Println("web berjalan dengan port 2514") //ganti 4 digit npm kalian masing”
	http.ListenAndServe(":2514", nil)            //ganti 4 digit npm kalian masing”
}

type lepkom struct {
	Nama   string `json:"nama_mahasiswa"`
	Kursus string `json:"kursus_mahasiswa"`
	Foto   string `json:"foto_mahasiswa"`
}

var data_mahasiswa = []lepkom{
	{
		Nama:   "Wardiansyah", //ganti nama kalian
		Kursus: "Fundamental Web",
		Foto:   "img/gambar1.jpg",
	},
	{
		Nama:   "Bima",
		Kursus: "Golang Intermediate",
		Foto:   "img/gambar1.jpg",
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == http.MethodGet {
		result, err := json.Marshal(data_mahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
}
