package day4

import (
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func count(ls []string, c string) int {
	//slice element counting function, change input type as needed
	count := 0
	for _, i := range ls {
		if i == c {
			count++
		}
	}
	return count
}
func checkWord(crsWrd [][]string, rowI int, colI int) int {
	count := 0
	var found bool
	letters := []string{"M", "A", "S"}
	maxY := len(crsWrd) - 1
	maxX := len(crsWrd[0]) - 1
	newRowI := rowI
	newColI := colI
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, d := range directions {
		for i := range letters {
			newRowI = newRowI + d[0]
			newColI = newColI + d[1]
			if newRowI < 0 || newColI < 0 || newRowI > maxY || newColI > maxX {
				found = false
				break
			} else if crsWrd[newRowI][newColI] == letters[i] {
				found = true
			} else {
				found = false
				break
			}
		}
		if found {
			count += 1
		}
		newRowI = rowI
		newColI = colI
	}
	return count
}

func checkMas(crsWrd [][]string, rowI int, colI int) int {
	var valid int
	pair1 := []string{crsWrd[rowI-1][colI-1], crsWrd[rowI+1][colI+1]}
	pair2 := []string{crsWrd[rowI-1][colI+1], crsWrd[rowI+1][colI-1]}
	if count(pair1, "M") == 1 && count(pair1, "S") == 1 && count(pair2, "M") == 1 && count(pair2, "S") == 1 {
		valid = 1
	} else {
		valid = 0
	}
	return valid
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	input := fileparse.FileParse("day4/Input.txt")
	var crossWord [][]string
	for _, i := range input {
		crossWord = append(crossWord, strings.Split(i, ""))
	}
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	wordCount := 0
	for rowIdx, row := range crossWord {
		for colIdx, col := range row {
			if col == "X" {
				wordCount += checkWord(crossWord, rowIdx, colIdx)
			}
		}
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, wordCount
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	input := fileparse.FileParse("day4/Input.txt")
	var crossWord [][]string
	for _, i := range input {
		crossWord = append(crossWord, strings.Split(i, ""))
	}
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	XCount := 0
	for rowIdx, row := range crossWord {
		if rowIdx != 0 && rowIdx != (len(crossWord)-1) {
			for colIdx, col := range row {
				if colIdx != 0 && colIdx != (len(row)-1) {
					if col == "A" {
						XCount += checkMas(crossWord, rowIdx, colIdx)
					}
				}
			}
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, XCount
}
