package main

import (
	"fmt"
	"math"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 16842 in 47.58µs
Part 2: 167523425665348 in 191.65µs
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

	for scanner.Scan() {
		bank := scanner.Text()
		batteries := [12]int{}
		lenBat := len(batteries)
		lenBank := len(bank)

		for idx, r := range bank {
			n := runeToInt(r)
			var start int
			if start = lenBat - (lenBank - idx); start < 0 {
				start = 0
			}
			for i := start; i < lenBat; i++ {
				if n > batteries[i] {
					batteries[i] = n
					for j := i + 1; j < lenBat; j++ {
						batteries[j] = 0
					}
					break
				}
			}

		}
		for i, c := range batteries {
			mul := lenBat - i - 1
			total += c * int(math.Pow10(mul))
		}
	}
	return total
}

func runeToInt(r rune) int {
	return int(r - '0')
}
