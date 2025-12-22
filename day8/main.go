package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 538191549061 in 17.9063ms
Part 2: 34612812972206 in 1.6353173s
*/
func main() {
	start := time.Now()
	part1 := part1("sample.txt")
	fmt.Printf("Part 1: %v in %s\n", part1, time.Since(start))

	start = time.Now()
	part2 := part2("input.txt")
	fmt.Printf("Part 2: %v in %s\n", part2, time.Since(start))
}

func part1(f string) int {
	var maxPair int
	if f == "input.txt" {
		maxPair = 1000
	} else if f == "sample.txt" {
		maxPair = 10
	}

	scanner := util.CreateScannerFromFile(f)
	total := 0

	boxes := []Junction{}
	for scanner.Scan() {
		v := NewJunction(scanner.Text())
		boxes = append(boxes, v)
	}

	h := &CircuitHeap{}
	heap.Init(h)
	for i := 0; i <= len(boxes)-1; i++ {
		for j := i + 1; j <= len(boxes)-1; j++ {
			d := calcDistance(boxes[i], boxes[j])
			heap.Push(h, Circuit{Distance: d, Left: boxes[i], Right: boxes[j]})
		}
	}

	circuits := [][]Junction{}
	connections := maxPair

	for connections != 0 {
		c := heap.Pop(h).(Circuit)
		leftIndex := -1
		rightIndex := -1

		for i := range circuits {

			if slices.Contains(circuits[i], c.Left) {
				leftIndex = i
			}
			if slices.Contains(circuits[i], c.Right) {
				rightIndex = i
			}

			if leftIndex >= 0 && rightIndex >= 0 {
				break
			}
		}
		if (leftIndex >= 0 || rightIndex >= 0) && leftIndex == rightIndex {
			continue // same circuit already
		} else if leftIndex >= 0 && rightIndex == -1 {
			circuits[leftIndex] = append(circuits[leftIndex], c.Right)
			connections--
		} else if leftIndex == -1 && rightIndex >= 0 {
			circuits[rightIndex] = append(circuits[rightIndex], c.Left)
			connections--
		} else if leftIndex >= 0 && rightIndex >= 0 {
			circuits[leftIndex] = append(circuits[leftIndex], circuits[rightIndex]...) // merge circuits
			circuits = slices.Delete(circuits, rightIndex, rightIndex+1)
			connections--
		} else {
			circuits = append(circuits, []Junction{c.Left, c.Right}) // create new circuit
			connections--
		}
	}

	subtotal := 1
	slices.SortFunc(circuits, func(a, b []Junction) int {
		return len(b) - len(a) // descending order
	})
	top3 := circuits[:3]
	for _, group := range top3 {
		subtotal = len(group) * subtotal
	}
	total = subtotal

	return total
}

func part2(f string) int {
	// scanner := util.CreateScannerFromFile(f)
	total := 0

	return total
}

func calcDistance(v1, v2 Junction) float64 {
	return math.Sqrt(math.Pow(float64(v1.X-v2.X), 2) + math.Pow(float64(v1.Y-v2.Y), 2) + math.Pow(float64(v1.Z-v2.Z), 2))
}
func NewJunction(s string) Junction {
	parts := strings.Split(s, ",")
	n1, _ := strconv.Atoi(parts[0])
	n2, _ := strconv.Atoi(parts[1])
	n3, _ := strconv.Atoi(parts[2])

	return Junction{X: n1, Y: n2, Z: n3}
}

type Junction = struct {
	X, Y, Z int
}

type Circuit = struct {
	Left     Junction
	Right    Junction
	Distance float64
}

// An CircuitHeap is a min-heap of ints.
type CircuitHeap []Circuit

func (h CircuitHeap) Len() int           { return len(h) }
func (h CircuitHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h CircuitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CircuitHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Circuit))
}

func (h *CircuitHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *CircuitHeap) addValue(x any, size int) {
	if h.Len() >= size {
		h.Pop()
	}
	h.Push(x)
}
