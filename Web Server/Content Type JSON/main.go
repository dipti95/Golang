package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


var dictionary = map[string]string{
	"Go" :"A programming language.",
	"Gopher":"A software engineer who builds with Go.",
	"Golang":"Another name for Go.",
}

//Exercise
var cityPopulations = map[string]uint32{
	"Tokyo":       37435191,
	"Delhi":       29399141,
	"Shanghai":    26317104,
	"Sao Paulo":   21846507,
	"Mexico City": 21671908,
}

func getDictionary(w http.ResponseWriter, r* http.Request ){
    w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dictionary)
}

func index(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "HTML")

	w.WriteHeader(http.StatusOK)

	for key, val := range cityPopulations{
		fmt.Fprintf(w, "<h2>%s: %d</h2>", key, val)
	}
}

func getCities(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cityPopulations)
}


func main(){

   http.HandleFunc("/", getDictionary)
   http.HandleFunc("/getcities", index)
   http.HandleFunc("/cities", getCities)

	fmt.Println("Server Listening port 3000...")
	http.ListenAndServe(":3000", nil)
}