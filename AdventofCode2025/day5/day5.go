package day5

import (
	"fmt"
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

func Overlap(checked []int, source []int) ([]int, int) {
	total := 0
	if len(checked) == 0 {
		//first run only
		checked = append(checked, source...)
		return checked, total
	}

	if source[0] <= checked[1]+1 {
		//start overlaps with previous range
		if source[1] > checked[1] {
			//and new end is higher than the previous end
			checked[1] = source[1]
		}
	} else {
		//no overlap
		total = (checked[1] - checked[0] + 1)
		checked[0] = source[0]
		checked[1] = source[1]
	}
	return checked, total
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
	partial := 0
	//parse function
	ParseStart := time.Now()
	RangeText := fileparse.FileParse("day5/Range.txt")
	var RangeList [][]int

	for _, i := range RangeText {
		row := strings.Split(i, "-")
		v1, _ := strconv.Atoi(row[0])
		v2, _ := strconv.Atoi(row[1])
		rowInt := []int{v1, v2}
		RangeList = append(RangeList, rowInt)
	}
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	var mergedRange []int
	fmt.Println(len(RangeList))

	for len(RangeList) > 0 {
		LowCur := RangeList[0][0]
		idx := 0
		for j, k := range RangeList {
			if k[0] < LowCur {
				LowCur = k[0]
				idx = j
			}
		}
		mergedRange, partial = Overlap(mergedRange, RangeList[idx])
		p2 += partial
		RangeList = append(RangeList[:idx], RangeList[idx+1:]...)
	}

	p2 += (mergedRange[1] - mergedRange[0] + 1)

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
