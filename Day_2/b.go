package main

import (
	"bufio"
	"fmt"
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

	opponentLegend := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}
	scoreLegend := map[string]int{"X": 0, "Y": 3, "Z": 6}
	shapeScore := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3}
	sum := 0

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)

		opponentHand := opponentLegend[input[0:1]]
		result := scoreLegend[input[2:3]]

		sum += result
		myHand := getMyHand(opponentHand, result)
		sum += shapeScore[myHand]
	}

	fmt.Println("My Score: ", sum)
}

func getMyHand(opponent string, result int) string {
	if result == 3 {
		return opponent
	}
	var myHand string
	switch opponent {
	case "Rock":
		if result == 0 {
			myHand = "Scissors"
		} else {
			myHand = "Paper"
		}
	case "Paper":
		if result == 0 {
			myHand = "Rock"
		} else {
			myHand = "Scissors"
		}
	case "Scissors":
		if result == 0 {
			myHand = "Paper"
		} else {
			myHand = "Rock"
		}
	}
	return myHand
}
