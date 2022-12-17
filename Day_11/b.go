package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func B() {
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

	mod := 1
	for _, m := range monkeys {
		mod *= m.test
	}

	for round := 0; round < 10000; round++ {

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

				// reduce the stress level somehow
				worryLevel = worryLevel % mod
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
