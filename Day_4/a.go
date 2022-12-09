package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
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

		contained := checkContain(firstPair, secondPair)
		if contained {
			sum += 1
		}
	}

	fmt.Println("My Score: ", sum)
}

func checkContain(firstPair, secondPair string) bool {
	firstStart, firstEnd := getRange(firstPair)
	secondStart, secondEnd := getRange(secondPair)

	// first contained between second
	if secondStart <= firstStart && secondEnd >= firstEnd {
		return true
	}

	// second container between first
	if secondStart >= firstStart && secondEnd <= firstEnd {
		return true
	}
	return false
}

func getRange(pair string) (int, int) {
	split := strings.Split(pair, "-")
	firstNumber, _ := strconv.Atoi(split[0])
	secondNumber, _ := strconv.Atoi(split[1])
	return firstNumber, secondNumber
}
