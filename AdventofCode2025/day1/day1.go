package day1

import (
	"strconv"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func Part1() (time.Duration, time.Duration, int) {
	var direction string
	var weight int
	position := 50
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day1/Input.txt")
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP1 := time.Now()
	for _, i := range Input {
		direction = string(i[0])
		weight, _ = strconv.Atoi(string(i[1:]))
		switch direction {
		case "L":
			position = (position - weight) % 100
			if position < 0 {
				position = position + 100
			}
		case "R":
			position = (position + weight) % 100
		}
		if position == 0 {
			p1++
		}
	}
	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	var direction string
	var weight int
	position := 50
	p2 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day1/Input.txt")
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	for _, i := range Input {
		direction = string(i[0])
		weight, _ = strconv.Atoi(string(i[1:]))
		switch direction {
		case "L":
			p2 = p2 - (position-weight)/100
			position = (position - weight) % 100
			if (position < 0) && (position != (0 - (weight % 100))) { //moving left from a point other than zero
				p2++
			}
			if position == 0 { //land on 0
				p2++
			}
			if position < 0 { //reset to positive vale
				position = position + 100
			}

		case "R":
			p2 = p2 + ((position + weight) / 100)
			position = (position + weight) % 100
		}
	}

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
