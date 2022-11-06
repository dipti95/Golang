package main

import (
	"fmt"
	"strconv"
)


func main()  {
	
//example1
	for i:=0;i<=20;i++{
		if i%2 == 0{
		fmt.Println(i)
		}
	   
	}
// example2
	for i:=0; i <=20; i+=2{
		fmt.Println(i)
	}
// example3
 i:=0
for i <=20{
	fmt.Println(i)
	i++
}

fmt.Println(fizzbuzz(16))


}

func fizzbuzz(n int)[] string{
	//Exercise
var result[]string


for i:=1;i <=n;i++{
	if i % 3==0 {
		result = append(result, "Buzz")
	}else if i %5==0{
        result = append(result, "Fizz")
	}else if i%15 ==0{
		result = append(result, "FizzBuzz")
	}else {
		result = append(result, strconv.Itoa(i))
	}

	
}
return result
}

