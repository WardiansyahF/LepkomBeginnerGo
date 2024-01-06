package main

//Mohamad Rizki Adiansyah
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := 3030

	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Server running on http://localhost:%d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))

}
