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
		//fmt.Println(input[i])
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
	//return num1*num2
	return num1 * num2

}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	stg := getInput("day3/Input.txt")
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//fmt.Println(stg)
	pattn := regexp.MustCompile(`mul\(\d+,\d+\)`)
	StgMatch := pattn.FindAllString(stg, -1)
	//fmt.Println(StgMatch)
	Total := 0
	for i := range StgMatch {
		Total += multiple(StgMatch[i])
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, Total

}
