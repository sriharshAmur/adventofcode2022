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

	cycleCount := 0
	sum := 1
	signalStrength := 0

	for fileScanner.Scan() {
		strength := 0
		input := fileScanner.Text()
		input = string(input)

		inputStrings := strings.Split(input, " ")
		command := inputStrings[0]

		if command == "noop" {
			cycleCount += 1
			// end of cycle
			strength = updateStrength(cycleCount, sum)
			signalStrength += strength
		} else {
			for i := 0; i < 2; i++ {
				cycleCount += 1
				strength = updateStrength(cycleCount, sum)
				signalStrength += strength
			}
			// end of cycle
			number, _ := strconv.Atoi(inputStrings[1])
			sum += number
		}
	}
	fmt.Println("strength: ", signalStrength)
}

func updateStrength(cycleCount, sum int) int {
	strength := 0
	if cycleCount == 20 || (cycleCount > 20 && (cycleCount-20)%40 == 0) {
		strength = cycleCount * sum
	}
	return strength
}
