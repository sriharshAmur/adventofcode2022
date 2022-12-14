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

	headPosition := []int{0, 0} // x, y
	tailPosition := []int{0, 0}

	tailVisits := map[string]int{"0-0": 1}

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)

		inputStrings := strings.Split(input, " ")
		direction := inputStrings[0]
		moveCountString := inputStrings[1]
		moveCount, _ := strconv.Atoi(moveCountString)

		for i := 0; i < moveCount; i++ {
			prevHeadPosition := append([]int(nil), headPosition...)
			headPosition = updatePosition(headPosition, direction)
			closeBy := checkCloseBy(headPosition, tailPosition)

			if !closeBy {
				tailPosition = prevHeadPosition
				xCoord := strconv.Itoa(tailPosition[0])
				yCoord := strconv.Itoa(tailPosition[1])
				str := xCoord + "-" + yCoord
				tailVisits[str] = 1
			}
		}
	}
	fmt.Println("visits: ", len(tailVisits))
}

func updatePosition(position []int, direction string) []int {
	switch direction {
	case "R":
		position[0] = position[0] + 1
	case "L":
		position[0] = position[0] - 1
	case "U":
		position[1] = position[1] + 1
	case "D":
		position[1] = position[1] - 1
	}
	return position
}

func checkCloseBy(headPosition, tailPosition []int) bool {
	xHead, yHead := headPosition[0], headPosition[1]
	xTail, yTail := tailPosition[0], tailPosition[1]

	xDifference := xHead - xTail
	yDifference := yHead - yTail

	if xDifference > 1 || xDifference < -1 || yDifference > 1 || yDifference < -1 {
		return false
	} else {
		return true
	}

}
