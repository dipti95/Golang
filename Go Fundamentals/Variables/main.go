package main

import "fmt"

func main() {
	
	var language string = "Go"
	const released int = 2009
	var isAProgrammingLanguage = true
   

	fmt.Println("Language:",language)
	fmt.Println("Released year:", released)
	fmt.Println("Is a programming language:",isAProgrammingLanguage)


	// var product = "T-shirt"
	// var cost = 20

	// Inside the function we can also write variable shown below
	product:= "T-shirt"
	cost:= 20



  fmt.Println("product's value is:", product)

  fmt.Printf("product's type is: %T\n", product)
  // we can change the value of cost because cost is not constant
  cost = 15
  fmt.Println("product's cost is:", cost)
 
  fmt.Printf("cost type is: %T\n", cost)

}
