package day14

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func stepper(botStart []string, w int, h int, steps int) []int {
	/*
		Takes bot start and calculates the final position based
		on the grid size and vector. Wrap around grid.
	*/
	col, _ := strconv.Atoi(botStart[0])
	row, _ := strconv.Atoi(botStart[1])
	colV, _ := strconv.Atoi(botStart[2])
	rowV, _ := strconv.Atoi(botStart[3])

	var finalPos []int

	colOffset := ((steps * colV) % w)
	rowOffset := ((steps * rowV) % h)

	if (row + rowOffset) > (h - 1) {
		finalPos = append(finalPos, (rowOffset - ((h - 1) - row) - 1))
	} else if (row + rowOffset) < 0 {
		finalPos = append(finalPos, ((h - 1) + (row + rowOffset + 1)))
	} else {
		finalPos = append(finalPos, (row + rowOffset))
	}

	if (col + colOffset) > (w - 1) {
		finalPos = append(finalPos, (colOffset - ((w - 1) - col) - 1))
	} else if (col + colOffset) < 0 {
		finalPos = append(finalPos, ((w - 1) + (col + colOffset + 1)))
	} else {
		finalPos = append(finalPos, (col + colOffset))
	}

	return finalPos
}

func stepperS(botStart []string, w int, h int, steps int) string {
	/*
		Takes bot start and calculates the final position based
		on the grid size and vector. Wrap around grid.
		Saves coords as string to use in p2 for count function
	*/
	col, _ := strconv.Atoi(botStart[0])
	row, _ := strconv.Atoi(botStart[1])
	colV, _ := strconv.Atoi(botStart[2])
	rowV, _ := strconv.Atoi(botStart[3])

	var finalPos string

	colOffset := ((steps * colV) % w)
	rowOffset := ((steps * rowV) % h)

	if (row + rowOffset) > (h - 1) {
		finalPos = strconv.Itoa((rowOffset - ((h - 1) - row) - 1))
	} else if (row + rowOffset) < 0 {
		finalPos = strconv.Itoa(((h - 1) + (row + rowOffset + 1)))
	} else {
		finalPos = strconv.Itoa((row + rowOffset))
	}

	if (col + colOffset) > (w - 1) {
		finalPos = finalPos + "," + strconv.Itoa((colOffset - ((w - 1) - col) - 1))
	} else if (col + colOffset) < 0 {
		finalPos = finalPos + "," + strconv.Itoa(((w - 1) + (col + colOffset + 1)))
	} else {
		finalPos = finalPos + "," + strconv.Itoa((col + colOffset))
	}
	return finalPos
}
func adder(bots [][]int, w int, h int) int {
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	midRow := h / 2
	midCol := w / 2

	for _, i := range bots {
		if i[0] != midRow && i[1] != midCol {
			if i[0] < midRow && i[1] < midCol {
				q1++
			}
			if i[0] > midRow && i[1] < midCol {
				q3++
			}
			if i[0] < midRow && i[1] > midCol {
				q2++
			}
			if i[0] > midRow && i[1] > midCol {
				q4++
			}
		}
	}
	return (q1 * q2 * q3 * q4)
}

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

func uniqueCheck(bots []string) bool {
	midCount := 0
	for _, i := range bots {
		if count(bots, i) > 1 {
			midCount++
		}
	}
	return midCount < 2
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day14/Input.txt")

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	//Insert additional processsing here
	var input [][]string
	for _, i := range raw {
		var line []string
		line = append(line, strings.Split(strings.Split(strings.Split(i, " ")[0], ",")[0], "=")[1])
		line = append(line, strings.Split(strings.Split(i, " ")[0], ",")[1])
		line = append(line, strings.Split(strings.Split(strings.Split(i, " ")[1], ",")[0], "=")[1])
		line = append(line, strings.Split(strings.Split(i, " ")[1], ",")[1])
		input = append(input, line)
	}
	//insert Puzzle solution here
	width := 101
	height := 103
	steps := 100

	var finalCoords [][]int
	for _, i := range input {
		finalCoords = append(finalCoords, stepper(i, width, height, steps))
	}
	sum := adder(finalCoords, width, height)

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day14/Input.txt")

	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()

	//Insert additional processsing here
	var input [][]string
	for _, i := range raw {
		var line []string
		line = append(line, strings.Split(strings.Split(strings.Split(i, " ")[0], ",")[0], "=")[1])
		line = append(line, strings.Split(strings.Split(i, " ")[0], ",")[1])
		line = append(line, strings.Split(strings.Split(strings.Split(i, " ")[1], ",")[0], "=")[1])
		line = append(line, strings.Split(strings.Split(i, " ")[1], ",")[1])
		input = append(input, line)
	}

	//insert Puzzle solution here
	width := 101
	height := 103
	frame := 0
	for j := 7000; j < 10403; j++ {
		var finalCoords []string
		for _, b := range input {
			finalCoords = append(finalCoords, stepperS(b, width, height, j))
		}
		if uniqueCheck(finalCoords) {
			frame = j
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, frame
}
