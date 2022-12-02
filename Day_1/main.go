package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	max := 0
	temp := 0
	for fileScanner.Scan() {
		input := fileScanner.Text()

		if input == "" {
			if temp > max {
				max = temp
			}
			temp = 0
		} else {
			inputNumber, _ := strconv.Atoi(input)
			temp += inputNumber
		}
	}

	fmt.Println("Max Calorie: ", max)
}
