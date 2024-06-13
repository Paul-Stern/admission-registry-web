package main

import (
	"log"
	"net/http"

	"github.com/paul-stern/admission-registry-web/web"
)

func main() {
	// fmt.Println("Hello, World!")
	http.HandleFunc("/", web.Journal)
	log.Print("Server started. Listening to localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
