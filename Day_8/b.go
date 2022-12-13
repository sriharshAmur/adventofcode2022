package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	visibleTrees := calculateVisibleTrees(grid)

	max := 0
	for key, _ := range visibleTrees {
		splits := strings.Split(key, "-")
		row, _ := strconv.Atoi(splits[0])
		column, _ := strconv.Atoi(splits[1])
		// fmt.Print("Row and Column: ", row, column)
		score := getScenicScore(grid, row, column)
		if score > max {
			max = score
		}
	}

	fmt.Println("Result: ", max)

}

func getScenicScore(grid [][]int, row, column int) int {
	value := grid[row][column]
	columnLength := len(grid[0])
	rowLength := len(grid)

	// calculate for up direction -> row decreases
	upScore := 0
	for i := row - 1; i >= 0; i-- {
		val := grid[i][column]
		upScore += 1
		if val >= value {
			break
		}
	}

	// calculate for down direction -> row increases
	downScore := 0
	for i := row + 1; i < rowLength; i++ {
		val := grid[i][column]
		downScore += 1
		if val >= value {
			break
		}
	}

	// calculate for left direction -> column decreases
	leftScore := 0
	for i := column - 1; i >= 0; i-- {
		val := grid[row][i]
		leftScore += 1
		if val >= value {
			break
		}
	}

	// calculate for right direction -> column increases
	rightScore := 0
	for i := column + 1; i < columnLength; i++ {
		val := grid[row][i]
		rightScore += 1
		if val >= value {
			break
		}
	}

	score := upScore * downScore * leftScore * rightScore

	return score
}
