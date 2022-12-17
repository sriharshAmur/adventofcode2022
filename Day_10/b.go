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

	cycleCount := 0
	sum := 1
	crtPosition := 0
	var pixels []string

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)
		inputStrings := strings.Split(input, " ")
		command := inputStrings[0]

		if command == "noop" {
			cycleCount += 1
			pixel := getPixel(crtPosition, sum)
			pixels = append(pixels, pixel)
			// end of cycle
			crtPosition += 1

		} else {
			for i := 0; i < 2; i++ {
				cycleCount += 1

				pixel := getPixel(crtPosition, sum)
				pixels = append(pixels, pixel)
				crtPosition += 1
			}
			// end of cycle
			number, _ := strconv.Atoi(inputStrings[1])
			sum += number
		}

	}

	for i, pixel := range pixels {
		fmt.Print(pixel)
		if i%40 == 39 {
			fmt.Println()
		}
	}
}

func getPixel(crtPosition, register int) string {
	result := "."
	position := crtPosition % 40
	if position == register-1 || position == register || position == register+1 {
		result = "#"
	}
	return result
}
