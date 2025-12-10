package day7

import (
	"strconv"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func TachyonTrace(manifold [][]string, row int, pos int, cache map[string]int) int {
	n := strconv.Itoa(row) + "," + strconv.Itoa(pos)
	branch := 0

	if cache[n] == 0 {
		if row != (len(manifold) - 1) {
			switch manifold[row+1][pos] {
			case ".":
				branch = TachyonTrace(manifold, (row + 1), pos, cache)
			case "^":
				branch += TachyonTrace(manifold, (row + 1), (pos + 1), cache)
				branch += TachyonTrace(manifold, (row + 1), (pos - 1), cache)
			}
		} else {
			branch = 1
		}
		cache[n] = branch
	}
	return cache[n]
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
	//cache := make(map[string]int)
	p2 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.Grid("day7/Input.txt", "")

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	//var initialPos int
	for i := range Input[0] {
		if Input[0][i] == "S" {
			//initialPos = i
			Input[1][i] = "1"
		}
	}

	for i := 1; i < len(Input)-1; i++ {
		for j := 0; j < len(Input[i]); j++ {
			if (Input[i][j] != ".") && (Input[i][j] != "^") {
				switch Input[i+1][j] {
				case ".":
					Input[i+1][j] = Input[i][j]
				case "^":
					if Input[i+1][j+1] != "." {
						//Tachyon beam already present
						Input[i+1][j+1] = StringAdd(Input[i+1][j+1], Input[i][j])
					} else {
						Input[i+1][j+1] = Input[i][j]
					}
					if Input[i+1][j-1] != "." {
						//Tachyon beam already present
						Input[i+1][j-1] = StringAdd(Input[i+1][j-1], Input[i][j])
					} else {
						Input[i+1][j-1] = Input[i][j]
					}
				default:
					Input[i+1][j] = StringAdd(Input[i+1][j], Input[i][j])
				}
			}
		}
	}

	curval := "0"
	for _, i := range Input[(len(Input) - 1)] {
		if i != "." {
			curval = StringAdd(curval, i)
		}
	}
	p2, _ = strconv.Atoi(curval)
	//p2 = TachyonTrace(Input, 1, initialPos, cache)

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
