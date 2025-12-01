package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample.txt")
	want := 3

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestPmod(t *testing.T) {
	result1 := pmod(1, 3)
	result2 := pmod(-1, 3)

	want1 := 1
	want2 := 2

	if result1 != want1 {
		t.Fatalf(`Wanted %v, but got %v`, want1, result1)
	}

	if result2 != want2 {
		t.Fatalf(`Wanted %v, but got %v`, want2, result2)
	}

}

func TestSamplePart2(t *testing.T) {
	result := part2("sample.txt")
	want := 6

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}
