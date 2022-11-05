package main

import (
	"fmt"
	"strings"
)


func add(num1 , num2 int) int {
	return num1 +num2
}

func sayHello(name string) string {
	return "Hello "+name
}

func sayLoudly(phase string) string{
	return strings.ToUpper(phase)
}

func main(){
   fmt.Println(add(4, 5));
   fmt.Println(sayHello("Dipti"))
   fmt.Println(sayLoudly("I am developer"))
   
}