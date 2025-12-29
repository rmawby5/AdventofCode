package day12

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func StrMultiply(v1 string, v2 string) int {
	i1, _ := strconv.Atoi(v1)
	i2, _ := strconv.Atoi(v2)
	return (i1 * i2)
}
func Part1() (time.Duration, time.Duration, int) {
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day12/Input.txt")

	ParseTime := time.Since(ParseStart)
	//puzzle
	presentSize := []string{"7", "5", "7", "6", "7", "7"}
	p1 := 0
	startP1 := time.Now()
	for _, i := range Input {
		rowp := strings.Split(i, ": ")
		gr := strings.Split(rowp[0], "x")
		GridSize := StrMultiply(gr[0], gr[1])
		reqSize := 0
		presents := strings.Split(rowp[1], " ")
		for j := range presents {
			reqSize += StrMultiply(presentSize[j], presents[j])
		}
		if reqSize < GridSize {
			fmt.Print(GridSize)
			fmt.Print(" : ")
			fmt.Println(GridSize - reqSize)
			p1++
		}
	}

	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	//parse function
	ParseStart := time.Now()
	ParseTime := time.Since(ParseStart)
	//puzzle
	p2 := 0
	startP2 := time.Now()
	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
