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
	sum := 0

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)
		length := len(input)
		halfLength := length / 2
		firstComp := input[0:halfLength]
		secondComp := input[halfLength:]

		repeatedLetter := getRepeatedLetter(firstComp, secondComp)
		sum += getLetterValue(repeatedLetter)
	}

	fmt.Println("My Score: ", sum)
}

func getRepeatedLetter(firstComp, secondComp string) string {
	firstMap := make(map[int32]int)
	secondMap := make(map[int32]int)

	firstRune := []rune(firstComp)
	secondRune := []rune(secondComp)
	for i := 0; i < len(firstRune); i++ {
		if val, ok := firstMap[firstRune[i]]; ok {
			firstMap[firstRune[i]] = val + 1
		} else {
			firstMap[firstRune[i]] = 0
		}
	}

	for i := 0; i < len(secondRune); i++ {
		_, inFirstComp := firstMap[secondRune[i]]
		if inFirstComp {
			return string(secondRune[i])
		}
		if val, ok := secondMap[secondRune[i]]; ok {
			secondMap[secondRune[i]] = val + 1
		} else {
			secondMap[secondRune[i]] = 0
		}
	}
	return ""
}

func getLetterValue(letter string) int {
	number := letter[0]
	if number >= 97 {
		number -= 96 // 1-25
	} else {
		number = (number - 65) + 27
	}
	return int(number)
}
