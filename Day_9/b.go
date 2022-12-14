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

	headPosition := []int{0, 0} // x, y
	tailPositions := [][]int{}
	initialTail := []int{0, 0}
	tailPositions = append(tailPositions, initialTail)

	tailVisits := make(map[string]int)
	temp := 0
	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)

		inputStrings := strings.Split(input, " ")
		direction := inputStrings[0]
		moveCountString := inputStrings[1]
		moveCount, _ := strconv.Atoi(moveCountString)

		for move := 0; move < moveCount; move++ {
			headPosition = updatePosition(headPosition, direction)
			tailLength := len(tailPositions)

			for i := 0; i < tailLength; i++ {
				tailPosition := tailPositions[i]
				if i == 0 {
					closeBy, newTailPosition := checkAndUpdate(headPosition, tailPosition)
					if !closeBy {
						tailPositions[i] = newTailPosition
					}
				} else {
					prevTailPosition := tailPositions[i-1]
					closeBy, newTailPosition := checkAndUpdate(prevTailPosition, tailPosition)
					if !closeBy {
						tailPositions[i] = newTailPosition
					}
				}
			}

			// check if more tails are needed
			lastTail := tailPositions[tailLength-1]
			if tailLength < 9 {
				if lastTail[0] != 0 || lastTail[1] != 0 {
					tailPositions = append(tailPositions, []int{0, 0})
				}
			} else {
				xCoord := strconv.Itoa(lastTail[0])
				yCoord := strconv.Itoa(lastTail[1])
				str := xCoord + "-" + yCoord
				tailVisits[str] = 1
			}
		}
		temp += 1
	}
	fmt.Println("visits: ", len(tailVisits))
}

func checkAndUpdate(headPosition, tailPosition []int) (bool, []int) {
	xHead, yHead := headPosition[0], headPosition[1]
	xTail, yTail := tailPosition[0], tailPosition[1]

	xDifference := xHead - xTail
	yDifference := yHead - yTail

	// check for diagonal
	if (xDifference == 1 && yDifference == 2) || (xDifference == 2 && yDifference == 1) {
		tailPosition = updatePosition(tailPosition, "R")
		tailPosition = updatePosition(tailPosition, "U")
	} else if (xDifference == 1 && yDifference == -2) || (xDifference == 2 && yDifference == -1) {
		tailPosition = updatePosition(tailPosition, "R")
		tailPosition = updatePosition(tailPosition, "D")
	} else if (xDifference == -1 && yDifference == 2) || (xDifference == -2 && yDifference == 1) {
		tailPosition = updatePosition(tailPosition, "L")
		tailPosition = updatePosition(tailPosition, "U")
	} else if (xDifference == -1 && yDifference == -2) || (xDifference == -2 && yDifference == -1) {
		tailPosition = updatePosition(tailPosition, "L")
		tailPosition = updatePosition(tailPosition, "D")
	} else if xDifference > 1 || xDifference < -1 || yDifference > 1 || yDifference < -1 {
		if xDifference > 0 {
			// move right
			tailPosition = updatePosition(tailPosition, "R")
		} else if xDifference < 0 {
			// move left
			tailPosition = updatePosition(tailPosition, "L")
		}
		if yDifference > 0 {
			// move up
			tailPosition = updatePosition(tailPosition, "U")
		} else if yDifference < 0 {
			// move down
			tailPosition = updatePosition(tailPosition, "D")
		}
	} else {
		return true, tailPosition
	}
	return false, tailPosition
}
