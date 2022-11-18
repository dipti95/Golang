package main

import (
	"fmt"
	"log"
)

func main() {

	//fmt.Println(div(1, 0)) // this will panic
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(7, 2))

}

func safeDiv(a, b int) (q int, err error) {

	// q & err are local variable in safeDiv
	//(Just like a & b)
	defer func() {
		// e's type is any (or empty interface{}) *not* error
		if e := recover(); e != nil {
			log.Println("ERROR", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	// panic("ouch!")

	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
