package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 529 in 837.822µs
Part 2: 344260049617193 in 157.19µs
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

	freshIngredients := map[int]int{}

	// fresh ingredients
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		low, _ := strconv.Atoi(split[0])
		up, _ := strconv.Atoi(split[1])

		// check if already exist in map; skip if already inside interval
		if v, found := freshIngredients[low]; found {
			if v > up {
				continue
			}
		}
		freshIngredients[low] = up
	}

	// available ingredients
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)

		for low, high := range freshIngredients {
			if n >= low && n <= high {
				total += 1
				break // need to break because we didn't properly merge all intervals
			}
		}
	}

	return total
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	freshIngredients := map[int]int{}

	// fresh
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		low, _ := strconv.Atoi(split[0])
		up, _ := strconv.Atoi(split[1])

		// add new interval if current doesn't fit into an existing one
		exist, newLow, newUp := existingInterval(freshIngredients, low, up)
		if !exist {
			freshIngredients[low] = up
		}
		// rebalance intervals (fixes the need to order intervals)
		for exist && (newLow != low || newUp != up) {
			exist, newLow, newUp = existingInterval(freshIngredients, newLow, newUp)
		}
	}

	for low, up := range freshIngredients {
		total += up - low + 1
	}

	return total
}

func existingInterval(freshIngredients map[int]int, low, up int) (bool, int, int) {
	// check if interval overlaps existing interval
	for oldLow, oldUp := range freshIngredients {
		//skip yourself
		if low == oldLow && up == oldUp {
			continue
		} else if low < oldLow && up >= oldLow && up <= oldUp { // overlap left-side
			delete(freshIngredients, oldLow)
			freshIngredients[low] = oldUp
			return true, low, oldUp
		} else if low > oldLow && low <= oldUp && up > oldUp { // overlap right-side
			freshIngredients[oldLow] = up
			return true, oldLow, up
		} else if low >= oldLow && up <= oldUp { // overlap inside
			return true, oldLow, oldUp
		} else if low <= oldLow && up >= oldUp { // overlap outside
			delete(freshIngredients, oldLow)
			freshIngredients[low] = up
			return true, low, up
		}
	}
	return false, 0, 0
}
