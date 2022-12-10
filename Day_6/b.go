package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func B() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var result int
	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)
		result = getMarker14(input)
	}
	// result = getMarker14("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	result += 14

	fmt.Println("Result: ", result)
}

func getMarker14(input string) int {
	inputString := strings.Split(input, "")
	length := len(inputString)
	sum := 0
	for i := 13; i < length; i += 1 {
		var equal bool = false
		for j := i - 13; j < i; j += 1 {
			for k := j + 1; k <= i; k += 1 {
				if inputString[j] == inputString[k] {
					equal = true
					break
				}
			}
			if equal {
				break
			}
		}
		if equal {
			sum += 1
		} else {
			break
		}
	}
	return sum
}
