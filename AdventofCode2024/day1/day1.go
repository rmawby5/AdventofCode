package day1

import (
	"sort"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

// Seperate out indivual lines and characters
func linesplit(filelines []string) ([]int, []int) {
	var ls1 []int
	var ls2 []int
	var CurrentLine string
	var CurrentLineSplit []string
	var n1 int
	var n2 int
	var err1 error
	var err2 error
	for i := range filelines {
		CurrentLine = filelines[i]
		CurrentLineSplit = strings.Split(CurrentLine, "   ")
		n1, err1 = strconv.Atoi(CurrentLineSplit[0])
		n2, err2 = strconv.Atoi(CurrentLineSplit[1])
		if err1 != nil {
			// ... handle error
			panic(err1)
		}
		if err2 != nil {
			// ... handle error
			panic(err2)
		}
		ls1 = append(ls1, n1)
		ls2 = append(ls2, n2)
	}
	sort.Ints(ls1)
	sort.Ints(ls2)
	return ls1, ls2
}

func count(ls []int, c int) int {
	count := 0
	for _, i := range ls {
		if i == c {
			count++
		}
	}
	return count
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	var input = fileparse.FileParse("day1/Input.txt")
	ParseTime := time.Since(ParseStart)

	startP1 := time.Now()
	ls1, ls2 := linesplit(input)

	var Totaldistance int
	Totaldistance = 0
	for i := range ls1 {
		var diff = ls1[i] - ls2[i]
		if diff < 0 {
			Totaldistance = Totaldistance - diff
		} else {
			Totaldistance = Totaldistance + diff
		}
	}
	P1Time := time.Since(startP1)
	return ParseTime, P1Time, Totaldistance
}

func Part2() (time.Duration, time.Duration, int) {

	ParseStart := time.Now()
	var input = fileparse.FileParse("day1/Input.txt")
	ParseTime := time.Since(ParseStart)
	startP2 := time.Now()
	ls1, ls2 := linesplit(input)

	var Similarity int
	for i := range ls1 {
		Similarity = Similarity + (ls1[i] * count(ls2, ls1[i]))
	}
	P2Time := time.Since(startP2)

	return ParseTime, P2Time, Similarity
}
