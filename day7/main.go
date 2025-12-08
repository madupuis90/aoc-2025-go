package main

import (
	"fmt"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 1646 in 108.29µs
Part 2: 32451134474991 in 143.91µs
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

	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '^' && grid[i-1][j] == '|' { // check if splitter
				grid[i+1][j-1] = '|'
				grid[i+1][j+1] = '|'
				total += 1
			} else if grid[i][j] == '|' && i < len(grid)-1 && grid[i+1][j] == '.' { // check if beam
				grid[i+1][j] = '|'
			} else if grid[i][j] == 'S' { // check if Start
				grid[i+1][j] = '|'
				break
			}
		}
	}
	return total
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	counts := map[int]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '^' && grid[i-1][j] == '|' { // check if splitter
				grid[i+1][j-1] = '|'
				grid[i+1][j+1] = '|'
				if current, exists := counts[j]; exists {
					counts[j-1] += current
					counts[j+1] += current
				}
				counts[j] = 0
			} else if grid[i][j] == '|' && i < len(grid)-1 && grid[i+1][j] == '.' { // check if beam
				grid[i+1][j] = '|'
			} else if grid[i][j] == 'S' { // check if Start
				grid[i+1][j] = '|'
				counts[j] = 1
				break
			}
		}
	}

	for _, v := range counts {
		total += v
	}

	return total
}

// Anything below this point is basically useless. I tried to create a graph and traverse it to count, but it had way to many levels

func PrintGrid(grid [][]rune, counts map[Vector2D]int) {
	for i, r := range grid {
		for j, c := range r {
			if cc, exists := counts[Vector2D{X: i, Y: j}]; exists {
				fmt.Printf("%v", cc)
			} else {
				fmt.Printf("%c ", c)
			}
		}
		fmt.Println()
	}
}

func traverseTree(root *TreeNode) int {
	traversal := 0
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	traversal += traverseTree(root.Left)
	traversal += traverseTree(root.Right)
	return traversal
}
func createNode(grid [][]rune, current Vector2D) *TreeNode {
	var node *TreeNode
	if current.X >= len(grid)-1 {
		return nil
	} else if grid[current.X+1][current.Y] == '^' {
		node = &TreeNode{
			Value: current,
			Left:  createNode(grid, Vector2D{X: current.X + 1, Y: current.Y - 1}),
			Right: createNode(grid, Vector2D{X: current.X + 1, Y: current.Y + 1}),
		}
	} else if grid[current.X+1][current.Y] == '|' {
		node = &TreeNode{
			Value: current,
			Left:  createNode(grid, Vector2D{X: current.X + 1, Y: current.Y}),
		}
	}
	return node
}

type Vector2D = struct{ X, Y int }

type TreeNode struct {
	Value Vector2D
	Left  *TreeNode
	Right *TreeNode
}
