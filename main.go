package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

type Team struct {
	Name	string
	ID	string
}

func main() {
	fmt.Println("Hello World!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		
		teams := map[string][]Team{
			"Teams" : {
				{Name: "Kowski Krew", ID: "01"},
				{Name: "JingleHeimerSchmidt", ID:"02"},
			},
		}
		temp.Execute(w, teams)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
