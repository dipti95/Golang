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

func getRectangleArea(width , length int)string{
    product := width * length;

	if product < 50 {
		return fmt.Sprintf("The area is %d, which is less than 50", product)
	}else {
		return fmt.Sprintf("The area is %d, which is greater or equal to 50", product)
	}
}

func main(){
   fmt.Println(add(4, 5));
   fmt.Println(sayHello("Dipti"))
   fmt.Println(sayLoudly("I am developer"))
   fmt.Println(getRectangleArea(5,10))
   fmt.Println("Test")
   
}