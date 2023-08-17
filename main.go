package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func main() {
	fmt.Println("Hello World!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
