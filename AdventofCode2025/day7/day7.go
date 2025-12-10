package day7

import (
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.Grid("day7/Input.txt", "")

	ParseTime := time.Since(ParseStart)
	//puzzle

	startP1 := time.Now()
	for i, j := range Input[0] {
		if j == "S" {
			Input[1][i] = "|"
		}
	}

	for i := 1; i < len(Input)-1; i++ {
		for j := 0; j < len(Input[i]); j++ {
			if Input[i][j] == "|" {
				switch Input[i+1][j] {
				case ".":
					Input[i+1][j] = "|"
				case "^":
					Input[i+1][j+1] = "|"
					Input[i+1][j-1] = "|"
					p1++
				}

			}
		}

	}
	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	p2 := 0
	//parse function
	ParseStart := time.Now()
	//Input := fileparse.FileParse("day9/TestInput.txt")

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
