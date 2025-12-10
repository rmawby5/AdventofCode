package day7

import (
	"strconv"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func TachyonTrace(manifold [][]string, row int, pos int) int {

	branch := 0

	if row != (len(manifold) - 1) {
		switch manifold[row+1][pos] {
		case ".":
			branch = TachyonTrace(manifold, (row + 1), pos)
		case "^":
			branch += TachyonTrace(manifold, (row + 1), (pos + 1))
			branch += TachyonTrace(manifold, (row + 1), (pos - 1))
		}

	} else {
		branch = 1
	}

	return branch
}

func StringAdd(v1 string, v2 string) string {
	v1i, _ := strconv.Atoi(v1)
	v2i, _ := strconv.Atoi(v2)
	return strconv.Itoa(v1i + v2i)

}

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
	Input := fileparse.Grid("day7/TestInput.txt", "")

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	var initialPos int
	for i := range Input[0] {
		if Input[0][i] == "S" {
			initialPos = i
		}
	}

	p2 = TachyonTrace(Input, 1, initialPos)

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
