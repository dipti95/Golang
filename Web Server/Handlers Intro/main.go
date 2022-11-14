package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !!!!!")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact Info<h1>")
}

// Exercise
var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func cityList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of most populous cities.")
	for _, name := range cities {
		fmt.Fprintf(w, "%s\n", name)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)

	http.HandleFunc("/cities", cityList)

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}
