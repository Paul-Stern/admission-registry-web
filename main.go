package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paul-stern/admission-registry-web/model"
	"github.com/paul-stern/admission-registry-web/templates"
)

func main() {
	// fmt.Println("Hello, World!")
	test()
	http.HandleFunc("/", journal)
	log.Print("Server started. Listening to localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func journal(w http.ResponseWriter, r *http.Request) {
	templates.RenderTable(w)
}
func test() {
	for i := 0; i < 10; i++ {
		// fmt.Printf("age: %d\nmonth: %d\nday: %d\n\n", rand.Intn(100), rand.Intn(11)+1, rand.Intn(27)+1)
		fmt.Printf("%v+\n", model.RandEntry())
	}
}
