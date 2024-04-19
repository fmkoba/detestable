package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("running on http://localhost:8080")
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", fs))
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/index.html"))
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("GET /")
}
