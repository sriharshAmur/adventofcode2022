package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type monkey struct {
	id              int
	items           *[]int
	test            int
	operation       string
	operationNumber int
	testTrue        int
	testFalse       int
	inspectedItems  *int
}

func (m *monkey) print() {
	fmt.Printf("Monkey %v has inspected %v items\n", m.id, *m.inspectedItems)

}

func A() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	monkeys := parseInput(fileScanner)
	// for _, m := range monkeys {
	// 	m.print()
	// }
	// fmt.Println()

	for round := 0; round < 20; round++ {

		for _, monkey := range monkeys {

			for _, item := range *monkey.items {
				worryLevel := 0
				operationNumber := 0
				if monkey.operationNumber == -1 {
					operationNumber = item
				} else {
					operationNumber = monkey.operationNumber
				}
				if monkey.operation == "*" {
					worryLevel = item * operationNumber
				} else {
					worryLevel = item + operationNumber
				}
				worryLevel /= 3
				worryLevel = int(math.Round(float64(worryLevel)))
				if (worryLevel % monkey.test) == 0 {
					*monkeys[monkey.testTrue].items = append(*monkeys[monkey.testTrue].items, worryLevel)
					// fmt.Printf("Giving item %v to monkey %v\n", worryLevel, monkey.testTrue)
					// monkeys[monkey.testTrue].print()
				} else {
					*monkeys[monkey.testFalse].items = append(*monkeys[monkey.testFalse].items, worryLevel)
					// fmt.Printf("Giving item %v to monkey %v\n", worryLevel, monkey.testFalse)
					// monkeys[monkey.testFalse].print()
				}
			}
			length := len(*monkey.items)
			*monkey.inspectedItems += length
			*monkey.items = make([]int, 0)
		}
	}
	inspectedItems := getInspectedItems(monkeys)
	sort.Ints(inspectedItems)
	length := len(inspectedItems)
	result := inspectedItems[length-1] * inspectedItems[length-2]
	fmt.Println()
	for _, m := range monkeys {
		m.print()
	}

	fmt.Println("Result: ", result)

}

func getInspectedItems(monkeys []monkey) []int {
	inspectedItems := make([]int, 0)

	for _, monkey := range monkeys {
		inspectedItems = append(inspectedItems, *monkey.inspectedItems)
	}
	return inspectedItems
}

func parseInput(fileScanner *bufio.Scanner) []monkey {

	var monkeys []monkey

	r, _ := regexp.Compile(`\d+`)
	multipleReg, _ := regexp.Compile("[*]")

	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = string(input)

		inputStrings := r.FindString(input)
		id := stringToInt(inputStrings)
		inspectedItems := 0
		monkey := monkey{id: id, inspectedItems: &inspectedItems}

		fileScanner.Scan()
		input = fileScanner.Text()
		itemsString := r.FindAllString(input, -1)
		var items []int
		for _, item := range itemsString {
			items = append(items, stringToInt(item))
		}
		monkey.items = &items

		fileScanner.Scan()
		input = fileScanner.Text()
		if r.MatchString(input) {
			operationNumber := r.FindString(input)
			monkey.operationNumber = stringToInt(operationNumber)
		} else {
			monkey.operationNumber = -1
		}
		operationString := multipleReg.MatchString(input)
		if operationString {
			monkey.operation = "*"
		} else {
			monkey.operation = "+"
		}

		fileScanner.Scan()
		input = fileScanner.Text()
		test := r.FindString(input)
		monkey.test = stringToInt(test)

		fileScanner.Scan()
		input = fileScanner.Text()
		testTrue := r.FindString(input)
		monkey.testTrue = stringToInt(testTrue)

		fileScanner.Scan()
		input = fileScanner.Text()
		testFalse := r.FindString(input)
		monkey.testFalse = stringToInt(testFalse)

		monkeys = append(monkeys, monkey)
		fileScanner.Scan()
	}
	return monkeys

}

func stringToInt(input string) int {
	number, _ := strconv.Atoi(input)
	return number
}
