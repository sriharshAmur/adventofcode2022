package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type pair struct {
	left  []interface{}
	right []interface{}
}

func A() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File not opened.")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	pairs := parseInputs(fileScanner)
	for i, pair := range pairs {
		// fmt.Println("Pair: ", i+1)
		ordered := comparePair(pair.left, pair.right)
		if ordered == 1 {
			// fmt.Println(i+1, " is in the right order")
			sum += (i + 1)
		}
		// fmt.Println()
		// break
	}
	fmt.Println("Sum is ", sum)
}

// Rules
// left list can be smaller than the right list
// if both are ints then the left <= right
// if comparing a list and an int, then convert the int to a list and then start comparison

func comparePair(left, right []interface{}) int {
	for i := 0; i < len(left); i++ {
		if i == len(right) {
			// right ran out
			// fmt.Println("Right Ran out")
			return -1
		}

		l := left[i]
		r := right[i]

		if bothFloats(l, r) {
			// if left < right => return 1 (in order)
			// if left > right => return -1 (not in order)
			// if left == right, continue to next element
			lInt := l.(float64)
			rInt := r.(float64)
			// fmt.Println("Comparing ", lInt, rInt)
			if lInt < rInt {
				return 1
			} else if lInt > rInt {
				return -1
			} else {
				continue
			}
		} else if bothSlices(l, r) {
			// both are slices
			lArr := l.([]interface{})
			rArr := r.([]interface{})
			result := comparePair(lArr, rArr)
			// fmt.Println(lArr, rArr, " both are slices ", result)
			if result == 0 {
				// both elements are the same
				continue
			} else {
				// either in order or not
				return result
			}
		} else if isFloat(l) || isFloat(r) {
			// one of them is a float, convert them into a list
			// fmt.Println("One of them is a flaot")
			lArr := []interface{}{}
			rArr := []interface{}{}
			if isFloat(l) {
				lArr = append(lArr, l)
				rArr = r.([]interface{})
			} else {
				rArr = append(rArr, r)
				lArr = l.([]interface{})
			}
			result := comparePair(lArr, rArr)
			// fmt.Println(lArr, rArr, " both are slices ", result)
			if result == 0 {
				// both elements are not in order
				continue
			} else {
				return result
			}
		}
	}

	if len(left) < len(right) {
		// left has greater number of elements
		return 1
	}

	return 0
}

func isSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func isFloat(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Float64
}

func bothSlices(v, w interface{}) bool {
	return isSlice(v) && isSlice(w)
}

func bothFloats(v, w interface{}) bool {
	return isFloat(v) && isFloat(w)
}

func parseInputs(fileScanner *bufio.Scanner) []pair {
	pairs := []pair{}
	for fileScanner.Scan() {
		leftInput := string(fileScanner.Text())
		fileScanner.Scan()
		rightInput := string(fileScanner.Text())
		fileScanner.Scan()
		fileScanner.Text()

		leftPacket := []interface{}{}
		if err := json.Unmarshal([]byte(leftInput), &leftPacket); err != nil {
			fmt.Println(err)
		}

		rightPacket := []interface{}{}
		if err := json.Unmarshal([]byte(rightInput), &rightPacket); err != nil {
			fmt.Println(err)
		}
		newPair := pair{left: leftPacket, right: rightPacket}
		pairs = append(pairs, newPair)
		// break
	}
	// fmt.Println("Pairs: ", pairs)
	return pairs
}
