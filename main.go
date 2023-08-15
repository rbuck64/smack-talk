package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
)

func main() {
	fmt.Println("Hello World!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!\n")
		io.WriteString(w, r.Method)
	}

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
