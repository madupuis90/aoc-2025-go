package main

import (
	"fmt"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 538191549061 in 17.9063ms
Part 2: 34612812972206 in 1.6353173s
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

	for scanner.Scan() {
		bank := scanner.Text()

		batA := 0
		batB := 0

		for idx, r := range bank {
			n := runeToInt(r)

			if n > batA && idx != len(bank)-1 {
				batA = n
				batB = 0
			} else if n > batB {
				batB = n
			}
		}
		total += batA*10 + batB
	}
	return total
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	batteries := [12]int{}
	for scanner.Scan() {
		bank := scanner.Text()

		for idx, r := range bank {
			n := runeToInt(r)
			for i, bat := range batteries {
				if n > bat && idx != len(bank)-(i+1) {
					bat = n
					for j := i; j <= 12; j++ {
						batteries[j] = 0
					}
				}
			}
		}
	}
	return total
}

func runeToInt(r rune) int {
	return int(r - '0')
}
