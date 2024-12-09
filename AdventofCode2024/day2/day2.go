package day2

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func reportCheck(report []int) bool {
	changeType := 0
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff > -4 && diff < 4 && diff != 0 { //check for jump size
			if diff > 0 {
				changeType++
			} else {
				changeType--
			}
			if changeType != i && changeType != (0-i) { //check jump direction
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func lineCopy(report []int) []int {
	//returns a copy of the report
	line := make([]int, len(report))
	count := copy(line, report)
	if count == 0 {
		panic(count)
	}
	return line
}

func lineCopyDouble(report []int, i int) ([]int, []int) {
	//Creates two copies of the report
	Var1 := lineCopy(report)
	Var2 := lineCopy(report)
	reportVar1 := append(Var1[:(i-1)], Var1[i:]...)
	reportVar2 := append(Var2[:i], Var2[(i+1):]...)
	return reportVar1, reportVar2
}

func dampedReportCheck(report []int) bool {
	//Non brute-force solution
	changeType := 0
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff > -4 && diff < 4 && diff != 0 { //check for jump size
			if diff > 0 {
				changeType++
			} else {
				changeType--
			}
			if changeType != i && changeType != (0-i) { //check jump direction

				if i == 2 {
					reportTemp := lineCopy(report)
					dvar := reportTemp[1:]
					if reportCheck(dvar) {
						return true
					}
				}
				dvar, dvar2 := lineCopyDouble(report, i)
				if reportCheck(dvar) {
					return true
				}

				if reportCheck(dvar2) {
					return true
				}
				return false
			}
		} else {
			dvar, dvar2 := lineCopyDouble(report, i)
			if reportCheck(dvar) {
				return true
			}
			if reportCheck(dvar2) {
				return true
			}
			return false
		}
	}
	return true
}

func Part1() (time.Duration, time.Duration, int) {
	Startfile := time.Now()
	var input = fileparse.FileParse("day2/TestInput.txt")
	fileparsetime := time.Since(Startfile)
	P1Start := time.Now()
	safeCount := 0
	var lineFormated []string
	for i := range input {
		var lineCon []int
		lineFormated = strings.Split(input[i], " ")
		for j := range lineFormated {
			num, err := strconv.Atoi(lineFormated[j])
			if err != nil {
				panic(err)
			}
			lineCon = append(lineCon, num)
		}
		if reportCheck(lineCon) {
			safeCount++
		}
	}
	P1Time := time.Since(P1Start)
	return fileparsetime, P1Time, safeCount

}

func Part2() (time.Duration, time.Duration, int) {
	Startfile := time.Now()
	var input = fileparse.FileParse("day2/Input.txt")
	fileparsetime := time.Since(Startfile)
	P2Start := time.Now()
	dampedSafeCount := 0
	var lineFormated []string
	for i := range input {
		var lineCon []int
		lineFormated = strings.Split(input[i], " ")
		for j := range lineFormated {
			num, err := strconv.Atoi(lineFormated[j])
			if err != nil {
				panic(err)
			}
			lineCon = append(lineCon, num)
		}
		if dampedReportCheck(lineCon) {
			dampedSafeCount++
		}
	}
	P2Time := time.Since(P2Start)
	return fileparsetime, P2Time, dampedSafeCount
}
