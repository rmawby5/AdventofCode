package day3

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func getInput(Path string) string {
	input := fileparse.FileParse(Path)
	var fullStr string
	for i := range input {
		fullStr = fullStr + input[i]
	}
	return fullStr
}

func multiple(mul string) int {
	numb := strings.Split(strings.Split(strings.Split(mul, "(")[1], ")")[0], ",")
	num1, err1 := strconv.Atoi(numb[0])
	num2, err2 := strconv.Atoi(numb[1])
	if err1 != nil || err2 != nil {
		panic("Error on int conversion")
	}
	return num1 * num2
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	stg := getInput("day3/Input.txt")
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	Total := 0

	pattn := regexp.MustCompile(`mul\(\d+,\d+\)`)
	for _, i := range pattn.FindAllString(stg, -1) {
		Total += multiple(i)
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, Total
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	stg := getInput("day3/Input.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	includeMul := true
	adjustedTotal := 0

	pattn := regexp.MustCompile(`don\'t\(\)|do\(\)|mul\(\d+,\d+\)`)
	//StgMatch := pattn.FindAllString(stg, -1)

	for _, i := range pattn.FindAllString(stg, -1) {
		if i == "don't()" {
			includeMul = false
		}
		if i == "do()" {
			includeMul = true
		} else {
			if includeMul {
				adjustedTotal += multiple(i)
			}
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, adjustedTotal
}
