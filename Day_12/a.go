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

	distances := map[string]int{"0-0": 0}
	// toVisit := []string{"0-0"}
	visits := map[string]int{"0-0": 0}
	grid, _, end := parseInput(fileScanner)

	count := 0

	for calculateMapLength(visits) != 0 {

		// get which node to explore next
		currentVisits := deepCopyMap(visits)
		for nodeKey, nodeVisited := range currentVisits {
			if nodeVisited == 1 {
				continue
			}
			// fmt.Println("Current Node", nodeKey)
			currentDistance := distances[nodeKey]
			currentNodeCoords := getCoords(nodeKey)
			// get the nodes which are visitable
			nextNode := getNextNode(grid, currentNodeCoords)
			for _, node := range nextNode {
				distance := currentDistance + 1
				nodeKey := getKey(node)
				if nodeDistance, exists := distances[nodeKey]; !exists {
					distances[nodeKey] = distance
					// first time visiting this node, so add it
					visits[nodeKey] = 0
					// toVisit = append(toVisit, nodeKey)
				} else {
					if distance < nodeDistance {
						distances[nodeKey] = distance
					}
				}
			}
			// remove the current node from toVisit
			// fmt.Println("Removing ", nodeKey)
			visits[nodeKey] = 1
			// toVisit = append(toVisit[:index], toVisit[index+1:]...)
		}
		count += 1
		// fmt.Println("To Visits: ", visits)
		// fmt.Println(distances)
		// fmt.Println()
		// if count == 4 {
		// 	break
		// }
	}

	fmt.Println("Length to destination", distances[getKey(end)])
}

func deepCopyMap(maps map[string]int) map[string]int {
	cp := make(map[string]int)
	for k, v := range maps {
		cp[k] = v
	}
	return cp
}

func calculateMapLength(visits map[string]int) int {
	length := 0
	for _, v := range visits {
		if v == 0 {
			length += 1
		}
	}
	return length
}

func getCoords(node string) []int {
	input := strings.Split(node, "-")
	x, _ := strconv.Atoi(input[0])
	y, _ := strconv.Atoi(input[1])
	return []int{x, y}
}

func getKey(node []int) string {
	xCoord := strconv.Itoa(node[0])
	yCoord := strconv.Itoa(node[1])
	str := xCoord + "-" + yCoord
	return str
}

func getNextNode(grid [][]string, current []int) [][]int {
	var nodes [][]int
	directions := []string{"U", "L", "D", "R"}
	currentLetter := grid[current[0]][current[1]]
	// currentLetter = getRealValue(currentLetter)

	for _, dir := range directions {
		node := updatePosition(current, dir)
		if !validPosition(grid, node) {
			continue
		}
		nodeLetter := grid[node[0]][node[1]]
		// nodeLetter = getRealValue(nodeLetter)
		if !possibleMove(currentLetter, nodeLetter) {
			continue
		}
		nodes = append(nodes, node)
	}
	return nodes

}

// func getRealValue(letter string) string {
// 	if letter == "S" {
// 		return "a"
// 	} else if letter == "E" {
// 		return "z"
// 	}
// 	return letter
// }

func possibleMove(current, node string) bool {
	currentNumber := current[0]
	nodeNumber := node[0]
	return nodeNumber <= currentNumber+1
}

func validPosition(grid [][]string, node []int) bool {
	rowLength := len(grid)
	colLength := len(grid[0])
	return node[0] >= 0 && node[1] >= 0 && node[0] <= rowLength-1 && node[1] <= colLength-1
}

func updatePosition(current []int, direction string) []int {
	position := append([]int{}, current...)
	switch direction {
	case "R":
		position[1] = position[1] + 1
	case "L":
		position[1] = position[1] - 1
	case "U":
		position[0] = position[0] - 1
	case "D":
		position[0] = position[0] + 1
	}
	// fmt.Println("Position: ", position)
	return position
}

func parseInput(fileScanner *bufio.Scanner) ([][]string, []int, []int) {

	var grid [][]string
	rowIndex := 0

	var start []int
	var end []int

	for fileScanner.Scan() {
		var row []string
		input := fileScanner.Text()
		input = string(input)

		letters := strings.Split(input, "")
		for index, letter := range letters {
			if letter == "S" {
				start = append(start, rowIndex)
				start = append(start, index)
				letter = "a"
			} else if letter == "E" {
				end = append(end, rowIndex)
				end = append(end, index)
				letter = "z"
			}
			row = append(row, letter)
		}
		rowIndex += 1
		grid = append(grid, row)
	}
	return grid, start, end
}
