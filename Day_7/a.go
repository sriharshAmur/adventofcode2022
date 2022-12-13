package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	directories []*directory
	files       []file
	size        int
	parent      *directory
}

func (dir *directory) calculateSize() int {
	size := 0

	// calculate the size of all the files
	for _, file := range dir.files {
		size += file.size
	}

	// calculate the size of all the directories within
	for _, folder := range dir.directories {
		size += folder.calculateSize()
	}
	dir.size = size
	return size
}

func (dir *directory) print() {
	fmt.Println()
	fmt.Println("- ", dir.name, " Size: ", dir.size, " (dir) Parent - ", dir.parent)
	for _, file := range dir.files {
		fmt.Println("	- ", file.name, " (file, size=", file.size, ")")
	}
	for _, direct := range dir.directories {
		direct.print()
	}

	fmt.Println()
}

func A() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var root directory
	var size int

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
	size = calculateSize(root)
	fmt.Println("Result: ", size)
}

func calculateSize(root directory) int {
	var size int
	if root.size <= 100000 {
		size += root.size
	}
	for _, dir := range root.directories {
		size += calculateSize(*dir)
	}
	return size
}

func inputHandler(current **directory, input string, fileScanner *bufio.Scanner) {
	inputString := strings.Split(input, " ")
	symbol := inputString[0]

	if symbol == "$" {
		command := inputString[1]
		if command == "cd" {
			name := inputString[2]
			if name == ".." {
				*current = (*current).parent
			}
			for _, dir := range (*current).directories {
				if dir.name == name {
					*current = dir
				}
			}

		} else if command == "ls" {
			if next := listItems(current, fileScanner); next == "" {
			} else {
				inputHandler(current, next, fileScanner)
			}
		}
	}
}

func listItems(current **directory, fileScanner *bufio.Scanner) string {
	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)
		inputString := strings.Split(input, " ")
		symbol := inputString[0]

		if symbol == "$" {
			return input
		} else if symbol == "dir" {
			name := inputString[1]
			child := &directory{name: name}
			child.parent = *current
			(*current).directories = append((*current).directories, child)
		} else {
			// file
			name := inputString[1]
			size, _ := strconv.Atoi(symbol)
			file := file{name: name, size: size}
			(*current).files = append((*current).files, file)
		}
	}
	return ""
}
