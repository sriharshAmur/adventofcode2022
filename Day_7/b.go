package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var root directory

	fileScanner.Scan()
	input := fileScanner.Text()
	input = string(input)
	inputString := strings.Split(input, " ")
	name := inputString[2]
	root.name = name
	current := &root
	pointer := &current

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)
		inputHandler(pointer, input, fileScanner)
	}
	root.calculateSize()
	// root.print()
	sizes := getSizes(&root)
	sort.Ints(sizes)
	minimumSize := 30000000 - (70000000 - root.size)
	fmt.Println("Required space:", minimumSize)

	for _, size := range sizes {
		if size >= minimumSize {
			fmt.Println("Result: ", size)
			break
		}
	}
}

func getSizes(root *directory) []int {
	var sizes []int
	sizes = append(sizes, root.size)
	for _, dir := range root.directories {
		sizes = append(sizes, getSizes(dir)...)
	}
	return sizes
}
