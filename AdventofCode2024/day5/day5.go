package day5

import (
	"slices"
	"strconv"
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

func firstInd(ls []string, c []string) (bool, int) {
	//returns the first index
	var ind []int
	var lowestId int
	var validity bool
	for _, i := range c {
		if slices.Contains(ls, i) {
			ind = append(ind, slices.Index(ls, i))
		}
	}
	if len(ind) != 0 {
		lowestId = slices.Min(ind)
		validity = false
	} else {
		lowestId = -1
		validity = true
	}
	return validity, lowestId
}

func mapper() (map[string][]string, []string) {
	guide := fileparse.FileParse("day5/InputGuide.txt")
	pages := fileparse.FileParse("day5/Input.txt")
	pageGuide := make(map[string][]string)
	for _, row := range guide {
		firstpage := strings.Split(row, "|")[0]
		laterpage := strings.Split(row, "|")[1]
		i, ok := pageGuide[firstpage]
		if ok {
			pageGuide[firstpage] = append(i, laterpage)
		} else {
			var val []string
			val = append(val, laterpage)
			pageGuide[firstpage] = val
		}
	}
	return pageGuide, pages
}

func makeValid(line []string, pageGuide map[string][]string) []string {
	var buffer []string
	var isValid bool
	var firstIdx int
	for _, j := range line {
		invaildPg := pageGuide[j]
		isValid, firstIdx = firstInd(buffer, invaildPg)
		if isValid {
			buffer = append(buffer, j)
		} else {
			buffer = slices.Insert(buffer, firstIdx, j)
		}
	}
	return buffer
}

func lineValidation(i string, pageGuide map[string][]string) bool {
	var isValid bool
	var buffer []string
	line := strings.Split(i, ",")
	for _, j := range line {
		invaildPg := pageGuide[j]
		for _, k := range invaildPg {
			if count(buffer, k) != 0 {
				//pagefile is invalid
				isValid = false
				break
			}
			isValid = true
		}
		if isValid {
			buffer = append(buffer, j)
		} else {
			break
		}
	}
	return isValid
}

func Part1() (time.Duration, time.Duration, int) {
	//get file inputs and generate hash map
	ParseStart := time.Now()
	var counter int
	pageGuide, pages := mapper()
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	var validity bool
	for _, i := range pages {
		line := strings.Split(i, ",")
		validity = lineValidation(i, pageGuide)
		if validity {
			midVal, _ := strconv.Atoi(line[len(line)/2])
			counter += midVal
		}
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, counter
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	var counter int
	pageGuide, pages := mapper()
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	var validity bool
	for _, i := range pages {
		line := strings.Split(i, ",")
		validity = lineValidation(i, pageGuide)
		if !validity {
			validline := makeValid(line, pageGuide)
			midVal, _ := strconv.Atoi(validline[len(validline)/2])
			counter += midVal
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, counter
}
