package main

import (
	"fmt"

	"Aoc.com/AdventOfCode2024/day3"
)

func main() {
	//day3.Part2()
	Parse1Time, P1Time, P1Total := day3.Part2()
	fmt.Printf("Parse Time: %v ", Parse1Time)
	fmt.Printf("Part 1 Time: %v ", P1Time)
	fmt.Printf("Part 1 Total: %d\n", P1Total)
	Parse2Time, P2Time, P2Total := day3.Part2()
	fmt.Printf("Parse Time: %v ", Parse2Time)
	fmt.Printf("Part 2 Time: %v ", P2Time)
	fmt.Printf("Part 2 Total: %d ", P2Total)

}
