package day3

import (
	"math"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func Power(bank string, pos int, pow int) int {
	bankArr := strings.Split(bank, "")
	idx := pos
	D1 := 0
	D2 := 0
	for i := pos; i < (len(bank) - pow); i++ {
		curVal, _ := strconv.Atoi(bankArr[i])
		if curVal > D1 {
			D1 = curVal
			idx = i
		}
	}
	if pow > 0 {
		D2 = Power(bank, (idx + 1), (pow - 1))
	}
	return (D1 * int(math.Pow(10, (float64(pow))))) + D2
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day3/Input.txt")
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP1 := time.Now()
	for _, i := range Input {
		p1 += Power(i, 0, 1)
	}
	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	p2 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day3/Input.txt")
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	for _, i := range Input {
		p2 += Power(i, 0, 11)
	}
	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
