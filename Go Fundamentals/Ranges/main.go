package main

import "fmt"

func main() {
   
	athletes := []string{"Stephen", "Klay", "Harrison","Draymond", "Andrew"}

	for i, name:= range athletes{
		fmt.Printf("i = %d, name= %s\n",i, name)
	}

	for _, name := range athletes {
       fmt.Printf("Name =%s\n", name)
	}

	numbers := []int{30,11,40,23,12}

	for _, num := range numbers{
		if num %2 == 0 {
			fmt.Println(num)
		}
	}

	fmt.Println(reduce([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}


func reduce(list []int) int {
	sum := 0

	for _, num := range list{
		sum += num;

	}

	return sum
}