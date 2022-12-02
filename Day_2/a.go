package main

import (
	"bufio"
	"fmt"
	"os"
)

func A() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	opponentLegend := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}
	myLegend := map[string]string{"X": "Rock", "Y": "Paper", "Z": "Scissors"}
	shapeScore := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3}
	resultScore := map[string]int{"lost": 0, "draw": 3, "won": 6}
	sum := 0

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)

		opponentHand := opponentLegend[input[0:1]]
		myHand := myLegend[input[2:3]]

		sum += shapeScore[myHand]
		result := getResult(opponentHand, myHand)
		sum += resultScore[result]
	}

	fmt.Println("My Score: ", sum)
}

func getResult(opponent, me string) string {
	if opponent == me {
		return "draw"
	} else {
		var result string
		switch opponent {
		case "Rock":
			if me != "Paper" {
				result = "lost"
			} else {
				result = "won"
			}
		case "Paper":
			if me != "Scissors" {
				result = "lost"
			} else {
				result = "won"
			}
		case "Scissors":
			if me != "Rock" {
				result = "lost"
			} else {
				result = "won"
			}
		}
		return result
	}
}
