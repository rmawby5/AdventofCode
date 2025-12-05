package day5

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func freshCheck(rangeList [][]int, Item string) bool {
	id, _ := strconv.Atoi(Item)
	for _, i := range rangeList {
		if (id >= i[0]) && (id <= i[1]) {
			return true
		}
	}
	return false
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day5/Input.txt")
	RangeText := fileparse.FileParse("day5/Range.txt")
	var RangeList [][]int

	for _, i := range RangeText {
		row := strings.Split(i, "-")
		v1, _ := strconv.Atoi(row[0])
		v2, _ := strconv.Atoi(row[1])
		rowInt := []int{v1, v2}
		RangeList = append(RangeList, rowInt)
		//fmt.Println(RangeList[l])

	}

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP1 := time.Now()
	for _, i := range Input {
		if freshCheck(RangeList, i) {
			p1++
		}

	}

	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	p2 := 0
	//parse function
	ParseStart := time.Now()
	//Input := fileparse.FileParse("day4/Input.txt")

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
