package main

import (
	"fmt"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 1604 in 248.271µs
Part 2: 9397 in 4.304167ms
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

	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	for x := range grid {
		for y, r := range grid[x] {

			if r == '@' {
				c := countAdjacentSymbols(grid, x, y)
				if c < 4 {
					total += 1
				}
			}

		}
	}

	return total
}

func countAdjacentSymbols(grid [][]rune, x, y int) int {
	count := 0
	lenX, lenY := len(grid), len(grid[0])

	for i := -1; i <= 1; i++ {
		newX := x + i
		if newX < 0 || newX > lenX-1 {
			continue
		}
		for j := -1; j <= 1; j++ {
			newY := y + j
			if i == 0 && j == 0 {
				continue
			}
			if newY < 0 || newY > lenY-1 {
				continue
			}
			if grid[newX][newY] == '@' {
				count++
			}
		}
	}

	return count
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	oneMoreTime := true
	for oneMoreTime {
		toBeChanged := map[[2]int]Void{}
		for x := range grid {
			for y, r := range grid[x] {
				if r == '@' {
					c := countAdjacentSymbols(grid, x, y)
					if c < 4 {
						total += 1
						toBeChanged[[2]int{x, y}] = Void{}
					}
				}
			}
		}

		if len(toBeChanged) == 0 {
			oneMoreTime = false
		}
		for k := range toBeChanged {
			grid[k[0]][k[1]] = '.'
		}
	}

	return total
}

type Void = struct{}
