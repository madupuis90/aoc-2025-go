package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/aoc-2025-go/util"
)

/*
Part 1: 18700015741 in 54.9µs
Part 2: 20077272987 in 23.783843ms
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

	if scanner.Scan() {
		line := scanner.Text()
		for _, r := range strings.Split(line, ",") {
			s := strings.Split(r, "-")
			lowerStr := s[0]
			upperStr := s[1]

			//Step 1: clean ranges

			// find even range lower (98 --> 100)
			if len(lowerStr)%2 != 0 {
				lowerStr = "1" + strings.Repeat("0", len(lowerStr))
			}
			// find even range upper (105 --> 99)
			if len(upperStr)%2 != 0 {
				upperStr = strings.Repeat("9", len(upperStr)-1)
			}
			// skip if new lower is greater than upper (could happen if both range gets trimmed)
			lower, _ := strconv.Atoi(lowerStr)
			upper, _ := strconv.Atoi(upperStr)
			if lower > upper {
				continue
			}

			// Step 2: find possible candidates by splitting strings in half
			lowerHalfStr := lowerStr[:len(lowerStr)/2]
			upperHalfStr := upperStr[:len(upperStr)/2]
			lowerHalf, _ := strconv.Atoi(lowerHalfStr)
			upperHalf, _ := strconv.Atoi(upperHalfStr)

			candidates := []int{}

			for i := lowerHalf; i <= upperHalf; i++ {
				candidateStr := strings.Repeat(strconv.Itoa(i), 2)
				candidate, _ := strconv.Atoi(candidateStr)
				// Validate that the candidates are in the initial range (candidates for range 0-50 would be ["11", "22", "33", "44", "55"])
				if candidate >= lower && candidate <= upper {
					candidates = append(candidates, candidate)
				}
			}

			// Step 3: Add everyting
			for _, c := range candidates {
				total += c
			}
		}
	}

	return total
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	if scanner.Scan() {
		line := scanner.Text()
		for _, r := range strings.Split(line, ",") {
			s := strings.Split(r, "-")
			lowerStr := s[0]
			upperStr := s[1]
			lower, _ := strconv.Atoi(lowerStr)
			upper, _ := strconv.Atoi(upperStr)

			// Step 1: find possible candidates by splitting strings in half
			lowerHalfStr := lowerStr[:(len(lowerStr))/2]
			upperHalfStr := upperStr[:(len(upperStr)+1)/2]
			lowerHalf, _ := strconv.Atoi(lowerHalfStr)
			upperHalf, _ := strconv.Atoi(upperHalfStr)

			candidatesMap := map[int]Void{}

			// Step 2: Iterate from the left half to the right half (e.g. 90-115 --> iterate from 9 to 11)
			for i := lowerHalf; i <= upperHalf; i++ {
				iStr := strconv.Itoa(i)
				// Iterate over all pattern of i (e.g. 21212 -> 2 -> 21 -> 212 -> 2121 -> 21212)
				for j := 1; j <= len(iStr); j++ {
					pattern := iStr[0:j]

					// here I just repeat the pattern an arbitrary number of time that is bigger than the boundries (e.g. 212 -> 212212212212212212)
					candidateStr := strings.Repeat(pattern, len(upperStr))

					// then I trim the excess so it is the same size as lower
					if len(lowerStr)%len(pattern) == 0 && len(lowerStr) != 1 { // single digit is not considered repeated
						candidateStr1 := candidateStr[:len(lowerStr)]
						candidate, _ := strconv.Atoi(candidateStr1)
						if candidate >= lower && candidate <= upper { // add to the map if it is in range
							candidatesMap[candidate] = Void{}
						}
						if candidate == 1 {
							fmt.Println(1)
						}
					}

					// and also trim the excess so it is the same size as upper
					if len(upperStr)%len(pattern) == 0 {
						candidateStr2 := candidateStr[:len(upperStr)]
						candidate, _ := strconv.Atoi(candidateStr2)
						if candidate >= lower && candidate <= upper { // add to the map if it is in range
							candidatesMap[candidate] = Void{}
						}
						if candidate == 1 {
							fmt.Println(1)
						}
					}
				}

			}

			for k := range candidatesMap {
				total += k
			}
		}

	}
	return total
}

type Void = struct{}
