package day4

import (
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func BorderCheck(grid [][]string, row int, col int) int {
	rollCount := 0
	for k := (row - 1); k < (row + 2); k++ {
		for l := (col - 1); l < (col + 2); l++ {
			if grid[k][l] == "@" {
				rollCount++
			}
		}
	}
	if rollCount < 5 {
		return 1
	}
	return 0
}

func BorderCheckRemove(grid [][]string, row int, col int) int {
	rollCount := 0
	for k := (row - 1); k < (row + 2); k++ {
		for l := (col - 1); l < (col + 2); l++ {
			if grid[k][l] == "@" {
				rollCount++
			}
		}
	}
	if rollCount < 5 {
		grid[row][col] = "."
		return 1
	}
	return 0
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day4/Input.txt")
	var InputGrid [][]string
	var boundary []string
	width := len(Input[0]) + 2
	for i := 0; i < width; i++ {
		boundary = append(boundary, ".")
	}
	InputGrid = append(InputGrid, boundary)
	for _, i := range Input {
		t := []string{"."}
		t = append(t, strings.Split(i, "")...)
		t = append(t, ".")
		InputGrid = append(InputGrid, t)
	}
	InputGrid = append(InputGrid, boundary)
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP1 := time.Now()
	for i := 1; i < (len(InputGrid) - 1); i++ {
		for j := 1; j < (width - 1); j++ {
			if InputGrid[i][j] == "@" {
				p1 += BorderCheck(InputGrid, i, j)
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
	Input := fileparse.FileParse("day4/Input.txt")
	var InputGrid [][]string
	var boundary []string
	width := len(Input[0]) + 2
	for i := 0; i < width; i++ {
		boundary = append(boundary, ".")
	}
	InputGrid = append(InputGrid, boundary)
	for _, i := range Input {
		t := []string{"."}
		t = append(t, strings.Split(i, "")...)
		t = append(t, ".")
		InputGrid = append(InputGrid, t)
	}
	InputGrid = append(InputGrid, boundary)
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	currentIterCnt := -1

	for currentIterCnt != 0 {
		currentIterCnt = 0
		for i := 1; i < (len(InputGrid) - 1); i++ {
			for j := 1; j < (width - 1); j++ {
				if InputGrid[i][j] == "@" {
					currentIterCnt += BorderCheckRemove(InputGrid, i, j)

				}
			}
		}
		p2 += currentIterCnt
	}

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
