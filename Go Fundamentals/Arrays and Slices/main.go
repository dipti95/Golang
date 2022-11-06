package main

import "fmt"


func main(){
	
	// var fibonacciArray [8] int

	// fibonacciArray[0]=0;
	// fibonacciArray[1]=1
	// fibonacciArray[2]=1
	// fibonacciArray[3]=2
	// fibonacciArray[4]=3
	// fibonacciArray[5]=5
	// fibonacciArray[6]=8
	// fibonacciArray[7]=13
 
    /** Array :
	   1. contains an ordered list of elements of single type
	   2. contains a specific number of elements
	*/     

	// we can also write like this
	 fibonacciArray := [8]int{0,1,1,2,3,5,8,13}

	fmt.Println(fibonacciArray)
	fmt.Println(fibonacciArray[0])
	fmt.Println(fibonacciArray[7])
	fmt.Println(len(fibonacciArray))
	fmt.Println(fibonacciArray[0:4])

	/** 
	Slice:
	1. Contains an ordered list of elements of single type
	2. Contains a list of items that can expand and shrink.(This is the key difference between a slice and an array!)
	*/

	fruitSlice := []string{"Pear", "Cherry", "Durian"}

	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[2])
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])

	fruitSlice = append(fruitSlice,"Grape", "Mango" )
     

	fmt.Println(fruitSlice)


	// Exercise

	languages := []string{"Go","Javascript", "Ruby", "Python"}

	fmt.Println(languages)
	fmt.Println(len(languages))
	fmt.Println(languages[0])
	fmt.Println(languages[1:3])

	languages = append(languages, "PHP")

	fmt.Println(languages)

	

}