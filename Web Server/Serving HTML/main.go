package main

import (
	"fmt"
	"net/http"
)



func main (){

	fileserver := http.FileServer(http.Dir("Web Server/Serving HTML/static"))

	http.Handle("/",fileserver )

	fmt.Println("Server is starting... ")
	http.ListenAndServe(":3000", nil)
}