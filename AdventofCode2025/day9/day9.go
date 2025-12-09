package day9

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func ab(num int) int {
	if num < 0 {
		return (0 - num)
	}
	return num
}

func Area(p1 []int, p2 []int) int {

	return (ab(p1[0]-p2[0]) + 1) * (ab(p1[1]-p2[1]) + 1)

}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day9/Input.txt")
	var InputInt [][]int
	for _, i := range Input {
		row := strings.Split(i, ",")
		v0, _ := strconv.Atoi(row[0])
		v1, _ := strconv.Atoi(row[1])
		InputInt = append(InputInt, []int{v0, v1})
	}
	ParseTime := time.Since(ParseStart)
	//puzzle

	startP1 := time.Now()

	for i := 0; i < len(InputInt)-1; i++ {
		for j := 1; j < len(InputInt); j++ {
			curArea := Area(InputInt[i], InputInt[j])
			if curArea > p1 {
				p1 = curArea
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
