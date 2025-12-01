package main

import (
	"fmt"
	"strconv"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 992 in 124.991µs
Part 2: 5635 in 116.93µs
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
	total := 0
	dial := 50
	scanner := util.CreateScannerFromFile(f)

	for scanner.Scan() {

		line := scanner.Text()
		direction := line[0:1]
		num, _ := strconv.Atoi(line[1:])

		if direction == "L" {
			dial = dial - num
			dial = pmod(dial, 100)
		}
		if direction == "R" {
			dial = dial + num
			dial = pmod(dial, 100)
		}

		if dial == 0 {
			total += 1
		}
	}
	return total
}

func part2(f string) int {
	total := 0
	dial := 50
	scanner := util.CreateScannerFromFile(f)

	for scanner.Scan() {

		line := scanner.Text()
		direction := line[0:1]
		num, _ := strconv.Atoi(line[1:])

		total += num / 100 // count full turns right away
		num = pmod(num, 100)

		if direction == "L" {
			diff := dial - num
			if dial != 0 && diff < 0 { // count if we pass 0, don't double count landing on zeroes
				total += 1
			}
			dial = pmod(diff, 100)

		}
		if direction == "R" {
			diff := dial + num
			if dial != 0 && diff > 100 { // count if we pass 0, don't double count landing on zeroes
				total += 1
			}
			dial = pmod(diff, 100)

		}

		if dial == 0 {
			total += 1
		}

	}
	return total
}

// Positive modulo, returns non negative solution to x % d
func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}
