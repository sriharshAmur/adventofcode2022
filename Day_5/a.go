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
	var stacks [][]string = createStacks(fileScanner)
	printStacks(stacks)
	fmt.Println()

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)
		number, from, to := getStep(input)
		// fmt.Println(number, from, to)
		stacks = moveCrates(number, from, to, stacks)
	}
	var top []string
	for _, stack := range stacks {
		top = append(top, stack[0])
	}
	topString := strings.Join(top[:], "")
	fmt.Println("Top: ", topString)
}

func printStacks(stacks [][]string) {
	for _, stack := range stacks {
		fmt.Println(stack)
	}
}

func moveCrates(number, from, to int, stacks [][]string) [][]string {
	crates := stacks[from-1][:number]
	newFrom := stacks[from-1][number:]
	stacks[from-1] = newFrom

	toStack := stacks[to-1]
	for _, crate := range crates {
		toStack = append([]string{crate}, toStack...)
	}

	stacks[to-1] = toStack
	return stacks

}

func getStep(input string) (int, int, int) {
	input = input[5:]
	firstSplit := strings.Split(input, "from")
	secondPart := firstSplit[1]
	secondSplit := strings.Split(secondPart, "to")
	numberString := strings.Trim(firstSplit[0], " ")
	fromString := strings.Trim(secondSplit[0], " ")
	toString := strings.Trim(secondSplit[1], " ")

	number, _ := strconv.Atoi(numberString)
	from, _ := strconv.Atoi(fromString)
	to, _ := strconv.Atoi(toString)

	return number, from, to
}

func createStacks(fileScanner *bufio.Scanner) [][]string {
	var input []string
	for fileScanner.Scan() {
		inputString := fileScanner.Text()
		inputString = string(inputString)
		if inputString == "" {
			break
		}
		input = append(input, inputString)
	}

	inputLength := len(input)
	noOfStacks := input[inputLength-1]
	stacksCount := strings.Split(noOfStacks, "  ")
	inputStacks := input[:inputLength-1]

	var stacks [][]string

	for i := range stacksCount {
		var stack []string
		for _, line := range inputStacks {
			index := (i * 3) + i + 1
			letter := string(line[index])
			if letter != " " {
				stack = append(stack, letter)
			}
		}
		stacks = append(stacks, stack)
	}

	return stacks
}
