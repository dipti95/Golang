package main

import (
	"fmt"
	"time"
)


func showMsg(msg string){
   for i :=0;i <=5;i++{
	time.Sleep(1* time.Second)

	fmt.Println(msg)
   }
}

// Exercise
func getFruit(fruits []string){
	
	for _, fruit := range fruits{
		time.Sleep(2* time.Second)
		fmt.Println(fruit)
	}
}

func main(){

	// Since this is a Goroutine, this will run concurrently (i.e., simultaneous and indpenedent of other program code)
    // As such, this Goroutine will not wait until it is finished executing
    // As a result, the `showMsg` call below will begin running at about the same time as the previous call!

   go showMsg("Hello!")


   // This function call does not require the above function call to finish before running itself
   showMsg("World")

   basket1 := []string{"Mango","Banana", "Pineapple"}
   basket2 := []string{"Pears","Grapes","Apple"}
   go getFruit(basket1)

   getFruit(basket2)
}