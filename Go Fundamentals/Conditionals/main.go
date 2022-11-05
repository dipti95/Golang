package main

import "fmt"


func main(){
	language := "Go"
    
	if language == "Java"{
		fmt.Println("language is Java")
	}else if language == "C"{
		fmt.Println("language is C")
	}else if language == "C++"{
		fmt.Println("language is C++")
	}else {
		fmt.Println("language is Go")
	}

	number:= 500

	if number < 0{
		fmt.Println( number,"is negative")
	}else if number < 100{
		fmt.Println( number,"is positive")
	}else {
		fmt.Println( number,"is positive and is large number!")
	}
}
