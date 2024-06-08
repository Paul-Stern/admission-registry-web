package main

import (
	"log"
	"net/http"

	"github.com/paul-stern/admission-registry-web/templates"
)

func main() {
	// fmt.Println("Hello, World!")
	http.HandleFunc("/", journal)
	log.Print("Server started. Listening to localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func journal(w http.ResponseWriter, r *http.Request) {
	templates.RenderTable(w)
}
