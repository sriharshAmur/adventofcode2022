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

	var grid [][]int

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)
		letters := strings.Split(input, "")
		var row []int
		for _, letter := range letters {
			number, _ := strconv.Atoi(letter)
			row = append(row, number)
		}
		grid = append(grid, row)
	}

	columnLength := len(grid[0])
	rowLength := len(grid)

	visibleTrees := calculateVisibleTrees(grid)

	fmt.Println(visibleTrees)
	result := 2*(rowLength+columnLength) - 4 + len(visibleTrees)

	fmt.Println("Result: ", result)
	// fmt.Println("Result: ", len(grid))
}

func calculateVisibleTrees(grid [][]int) map[string]int {
	columnLength := len(grid[0])
	rowLength := len(grid)
	visibleTrees := make(map[string]int)

	// find all trees visible from the columns view
	for i := 1; i < columnLength-1; i++ {
		firstTree := grid[0][i]
		lastTree := grid[rowLength-1][i]
		max := firstTree
		key := ""
		for j := 1; j < rowLength-1; j++ {
			tree := grid[j][i]
			if tree > max {
				max = tree
				key = strconv.Itoa(j) + "-" + strconv.Itoa(i)
				visibleTrees[key] = max
			}
		}

		max = lastTree
		key = ""
		for j := rowLength - 2; j > 0; j-- {
			tree := grid[j][i]
			if tree > max {
				max = tree
				key = strconv.Itoa(j) + "-" + strconv.Itoa(i)
				visibleTrees[key] = max

			}
		}

	}

	// find all trees visible from the rows view
	for i := 1; i < rowLength-1; i++ {
		firstTree := grid[i][0]
		max := firstTree
		key := ""
		for j := 1; j < columnLength-1; j++ {
			tree := grid[i][j]
			if tree > max {
				max = tree
				key = strconv.Itoa(i) + "-" + strconv.Itoa(j)
				visibleTrees[key] = max
			}
		}

		lastTree := grid[i][columnLength-1]
		max = lastTree
		key = ""
		for j := columnLength - 2; j > 0; j-- {
			tree := grid[i][j]
			if tree > max {
				max = tree
				key = strconv.Itoa(i) + "-" + strconv.Itoa(j)
				visibleTrees[key] = max
			}
		}

	}
	return visibleTrees
}
