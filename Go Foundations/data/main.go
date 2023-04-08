package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func getSum(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	inputStr := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputStr = append(inputStr, scanner.Text())

		err = scanner.Err()
		if err != nil {
			break
		}
	}

	sum := 0

	for _, value := range inputStr {
		val, _ := strconv.Atoi(value)
		sum += val
	}

	return sum
}

func main() {
	lines, err := ReadLines("data.txt")

	sum := getSum("input.txt")
	fmt.Println(sum)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}

}
