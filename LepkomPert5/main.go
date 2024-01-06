package main

import(
	"fmt"
	"log"
	"net/http"
)
func main()  {
	PORT := 1514

	http.Handle("/",http.FileServer(http.Dir("polymer")))
	fmt.Printf("Server running on http://localhost:%d",PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",PORT),nil))

}