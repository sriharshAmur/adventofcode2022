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
	sum := 0

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)

		split := strings.Split(input, ",")
		firstPair := split[0]
		secondPair := split[1]

		overlap := checkOverlap(firstPair, secondPair)
		if overlap {
			sum += 1
		}
	}

	fmt.Println("My Score: ", sum)
}

func checkOverlap(firstPair, secondPair string) bool {
	firstStart, firstEnd := getRange(firstPair)
	secondStart, secondEnd := getRange(secondPair)
	// first overlap between second: firstStart between the 2 number or firstEnd between the 2 number
	if (secondStart <= firstStart && firstStart <= secondEnd) || (secondStart <= firstEnd && firstEnd <= secondEnd) {
		return true
	} else if (firstStart <= secondStart && secondStart <= firstEnd) || (firstStart <= secondEnd && secondEnd <= firstEnd) {
		// second overlap between first
		return true
	}

	return false
}
