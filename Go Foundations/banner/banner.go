package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)

	// Strings in Go are utf-8

	// len in go is calculating the len of bytes
	// "G☺" has length :4
	// code point = rune ~= unicode character
	s := "G☺"
	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
			// rune (int32)
		}
	}

	b := s[0]

	// "%c" for character
	// "%T" for type
	// Printf => formated print statement
	fmt.Printf("%c of type %T\n", b, b)
	// bytes (uint8)

	x, y := 1, "1"

	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y) // Use #v in debug/log

	fmt.Printf("%20s!", s)

	fmt.Println("g", isPalindroms("g"))
	fmt.Println("go", isPalindroms("go"))
	fmt.Println("gog", isPalindroms("gog"))
	fmt.Println("G☺G", isPalindroms("G☺G"))
}

func isPalindroms(s string) bool {
	rs := []rune(s) // get slice of runes out of s
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}

	return true
}

func banner(text string, width int) {

	//padding := (width - len(text)) / 2 //Bug : len is in bytes
	padding := (width - utf8.RuneCountInString(text)) / 2
	// code point = rune ~= unicode character

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()

}
