package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	str := []string{}

	for scanner.Scan() {
		str = append(str, scanner.Text())

		err = scanner.Err()

		if err != nil {
			return nil, err
		}

	}
	return str, nil

}

func main() {
	lines, err := ReadLines("data.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}

}
