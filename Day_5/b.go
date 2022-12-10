package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// "strconv"

func B() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var stacks [][]string = createStacks(fileScanner)

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)
		number, from, to := getStep(input)
		stacks = moveCratesInOrder(number, from, to, stacks)
	}
	var top []string
	for _, stack := range stacks {
		top = append(top, stack[0])
	}
	topString := strings.Join(top[:], "")
	fmt.Println("Top: ", topString)
}

func moveCratesInOrder(number, from, to int, stacks [][]string) [][]string {
	crates := make([]string, number)
	copy(crates, stacks[from-1][:number])
	newFrom := stacks[from-1][number:]
	stacks[from-1] = newFrom
	toStack := stacks[to-1]

	toNewStack := append(crates, toStack...)
	stacks[to-1] = toNewStack
	return stacks
}
