package main

import (
	"fmt"
	"strings"
)

func main (){

	dictionary := map[string]string{
		"Go" :"A programming language.",
		"Gopher":"A software engineer who builds with Go.",
		
	}

	fmt.Println(dictionary)
	fmt.Println(dictionary["Go"])

	dictionary["Gopher"] ="software engineer."
    dictionary["Golang"] ="Another name for Go."
	fmt.Println(dictionary)

	for key,value  := range dictionary {
		fmt.Println(key , value)
	}

	delete(dictionary, "Gopher")
	fmt.Println(dictionary)

	//Exercise

	courses := map[int]string{
		1:"Calculus",
		2:"Biology",
		3:"Chemistry",
		4:"Computer Science",
		5:"Communications",
		6:"English",
		7:"Cantonese",
	}

	for _ , val := range courses {
		
		// if val[0:1] == "C"{
		// 	fmt.Println(val)
		// }

		if strings.HasPrefix(val ,"C"){
			fmt.Println(val)
		}
	}

	courses[4] = "Algorithms"
	courses[8] = "Spanish"
	 delete(courses, 1)

	 fmt.Println(courses)
}