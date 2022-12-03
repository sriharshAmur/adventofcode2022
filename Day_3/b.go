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
	sum := 0

	for fileScanner.Scan() {
		input1 := fileScanner.Text()
		fileScanner.Scan()
		input2 := fileScanner.Text()
		fileScanner.Scan()
		input3 := fileScanner.Text()

		repeatedLetter := getBadge(input1, input2, input3)
		sum += getLetterValue(repeatedLetter)
	}

	fmt.Println("My Score: ", sum)
}

func getBadge(first, second, third string) string {
	firstMap := make(map[int32]int)
	secondMap := make(map[int32]int)

	firstRune := []rune(first)
	secondRune := []rune(second)
	thirdRune := []rune(third)

	for i := 0; i < len(firstRune); i++ {
		if val, ok := firstMap[firstRune[i]]; ok {
			firstMap[firstRune[i]] = val + 1
		} else {
			firstMap[firstRune[i]] = 0
		}
	}

	for i := 0; i < len(secondRune); i++ {
		if val, ok := secondMap[secondRune[i]]; ok {
			secondMap[secondRune[i]] = val + 1
		} else {
			secondMap[secondRune[i]] = 0
		}
	}

	for i := 0; i < len(thirdRune); i++ {
		_, inFirstComp := firstMap[thirdRune[i]]
		_, inSecondComp := secondMap[thirdRune[i]]
		if inFirstComp && inSecondComp {
			return string(thirdRune[i])
		}
	}
	return ""
}
