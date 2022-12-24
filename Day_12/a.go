package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type coord struct {
	row    int
	column int
}

type square struct {
	value   string
	distace int
	coords  coord
}

func A() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	grid, start, end := parseInputs(fileScanner)
	grid[start.row][start.column].distace = 0

	toVisit := []coord{start}

	for {
		node := toVisit[0]
		cuurentSquare := grid[node.row][node.column]
		// fmt.Println("Current Square: ", cuurentSquare)
		distance := cuurentSquare.distace + 1

		if node.row == end.row && node.column == end.column {
			distance := grid[node.row][node.column].distace
			fmt.Println("Distance: ", distance)
			break
		}

		nextNodes := getNextNodes(grid, node)
		// fmt.Println("Possible moves: ", nextNodes)

		for _, node := range nextNodes {
			nodeSquare := getSqaure(grid, node)
			// fmt.Println("NodeSquare : ", nodeSquare)
			if distance < nodeSquare.distace {
				// fmt.Println("Distance is less ", distance, node)
				toVisit = append(toVisit, node)
				grid[node.row][node.column].distace = distance
			}
		}

		toVisit = toVisit[1:]
		// fmt.Println()
	}
}

func getNextNodes(grid [][]square, current coord) []coord {
	var nodes []coord
	directions := []string{"U", "L", "D", "R"}
	currentNode := grid[current.row][current.column]
	// currentLetter = getRealValue(currentLetter)

	for _, dir := range directions {
		node := updatePosition(current, dir)
		if !validPosition(grid, node) {
			continue
		}
		sq := grid[node.row][node.column]
		// nodeLetter = getRealValue(nodeLetter)
		if !possibleMove(currentNode.value, sq.value) {
			continue
		}
		nodes = append(nodes, node)
	}
	return nodes

}

func getSqaure(grid [][]square, node coord) square {
	return grid[node.row][node.column]
}

func possibleMove(current, node string) bool {
	currentNumber := current[0]
	nodeNumber := node[0]
	return nodeNumber <= currentNumber+1
}

func validPosition(grid [][]square, node coord) bool {
	rowLength := len(grid)
	colLength := len(grid[0])
	return node.row >= 0 && node.column >= 0 && node.row < rowLength && node.column < colLength
}

func updatePosition(position coord, direction string) coord {
	switch direction {
	case "R":
		position.column = position.column + 1
	case "L":
		position.column = position.column - 1
	case "U":
		position.row = position.row - 1
	case "D":
		position.row = position.row + 1
	}
	// fmt.Println("Position: ", position)
	return position
}

func parseInputs(fileScanner *bufio.Scanner) ([][]square, coord, coord) {

	var grid [][]square
	rowIndex := 0

	var start coord
	var end coord

	for fileScanner.Scan() {
		var row []square
		input := fileScanner.Text()
		input = string(input)

		letters := strings.Split(input, "")
		for index, letter := range letters {
			if letter == "S" {
				start = coord{row: rowIndex, column: index}
				letter = "a"
			} else if letter == "E" {
				end = coord{row: rowIndex, column: index}
				letter = "z"
			}
			coords := coord{row: rowIndex, column: index}
			temp := square{value: letter, coords: coords, distace: math.MaxUint32}
			row = append(row, temp)
		}
		rowIndex += 1
		grid = append(grid, row)
	}
	return grid, start, end
}
