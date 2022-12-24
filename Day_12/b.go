package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func B() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	grid, _, end := parseInputs(fileScanner)
	distance := math.MaxInt32

	startingPoints := getPossibleStartingPoints(grid)

	for _, point := range startingPoints {
		gridCopy := copyGrid(grid)
		dis := getShortestPath(gridCopy, point, end)
		if dis < distance {
			distance = dis
		}
	}

	fmt.Println("distance: ", distance)
}

func copyGrid(grid [][]square) [][]square {
	cp := [][]square{}
	for _, row := range grid {
		rows := []square{}
		rows = append(rows, row...)
		cp = append(cp, rows)
	}
	return cp
}

func getPossibleStartingPoints(grid [][]square) []coord {
	points := []coord{}
	for _, row := range grid {
		for _, col := range row {
			if col.value == "a" {
				points = append(points, col.coords)
			}
		}
	}
	return points
}

func getShortestPath(grid [][]square, start, end coord) int {
	// printGrid(grid)

	grid[start.row][start.column].distace = 0
	toVisit := []coord{start}

	for {
		if len(toVisit) == 0 {
			// printGrid(grid)
			return math.MaxUint32
		}
		node := toVisit[0]
		cuurentSquare := grid[node.row][node.column]
		// fmt.Println("Current Square: ", cuurentSquare)
		distance := cuurentSquare.distace + 1

		if node.row == end.row && node.column == end.column {
			distance := grid[node.row][node.column].distace
			return distance
		}

		nextNodes := getNextNodes(grid, node)

		for _, node := range nextNodes {
			nodeSquare := getSqaure(grid, node)
			if distance < nodeSquare.distace {
				toVisit = append(toVisit, node)
				grid[node.row][node.column].distace = distance
			}
		}

		toVisit = toVisit[1:]
	}
}

func printGrid(grid [][]square) {
	for _, row := range grid {
		for _, sq := range row {
			if sq.distace == math.MaxUint32 {
				fmt.Printf("[%v,%v]", sq.value, "__")
			} else {
				fmt.Printf("[%v,%v]", sq.value, sq.distace)
			}
		}
		fmt.Println()
	}
}
