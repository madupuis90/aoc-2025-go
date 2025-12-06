package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 5322004718681 in 209.951µs
Part 2: 9876636978528 in 351.25µs
*/
func main() {
	start := time.Now()
	part1 := part1("input.txt")
	fmt.Printf("Part 1: %v in %s\n", part1, time.Since(start))

	start = time.Now()
	part2 := part2("input.txt")
	fmt.Printf("Part 2: %v in %s\n", part2, time.Since(start))
}

func part1(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	numbers := [][]int{}
	operators := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "*") {
			break
		}
		for idx, str := range strings.Fields(line) {
			num, _ := strconv.Atoi(str)
			if idx >= len(numbers) {
				numbers = append(numbers, []int{})
			}
			numbers[idx] = append(numbers[idx], num)
		}
	}
	operators = append(operators, strings.Fields(scanner.Text())...)

	for i := range numbers {
		subtotal := 0
		switch operators[i] {
		case "*":
			subtotal = reduce(numbers[i], func(acc, current int) int {
				return acc * current
			})
		case "+":
			subtotal = reduce(numbers[i], func(acc, current int) int {
				return acc + current
			})
		default:
			log.Fatal("Expected operator to be * or +")
		}
		total += subtotal
	}

	return total
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	numbersStr := [][]rune{}
	operators := []string{}
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if strings.Contains(line, "*") {
			break
		}
		for idx, str := range []rune(line) {

			if idx >= len(numbersStr) {
				numbersStr = append(numbersStr, []rune{})
			}
			numbersStr[idx] = append(numbersStr[idx], str)
		}

	}
	operators = append(operators, strings.Fields(scanner.Text())...)

	arr := []int{}
	index := 0

	numbersStr = append(numbersStr, []rune(" ")) // small hack to trigger the last calculation
	for _, runes := range numbersStr {
		numStr := strings.TrimSpace(string(runes))
		subtotal := 0
		if numStr == "" { // calculate and clear buffer
			switch operators[index] {
			case "*":
				subtotal = reduce(arr, func(acc, current int) int {
					return acc * current
				})
			case "+":
				subtotal = reduce(arr, func(acc, current int) int {
					return acc + current
				})
			default:
				log.Fatal("Expected operator to be * or +")
			}
			arr = []int{}
			index++
			total += subtotal
			continue
		}
		num, _ := strconv.Atoi(numStr)
		arr = append(arr, num)
	}

	return total
}

func reduce(arr []int, fn func(acc int, current int) int) int {
	acc := 0
	if len(arr) >= 1 {
		acc = arr[0]
	}
	for i := 1; i < len(arr); i++ {
		acc = fn(acc, arr[i])
	}
	return acc
}
