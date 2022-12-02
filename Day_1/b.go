package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func B() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	temp := 0
	sumArray := []int{}
	for fileScanner.Scan() {
		input := fileScanner.Text()
		if input == "" {
			if temp != 0 {
				sumArray = append(sumArray, temp)
				temp = 0
			}
		} else {
			inputNumber, _ := strconv.Atoi(input)
			temp += inputNumber
		}
	}

	// sort the array
	sort.Slice(sumArray, func(i, j int) bool {
		return sumArray[i] > sumArray[j]
	})

	// get the top 3 and add them up
	sum := sumArray[0] + sumArray[1] + sumArray[2]

	fmt.Println("Max Calorie: ", sum)
}
