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
			if changeType != i && changeType != (0-i) {
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
	Var1 := lineCopy(report)
	reportVar1 := append(Var1[:(i-1)], Var1[i:]...)
	Var2 := lineCopy(report)
	reportVar2 := append(Var2[:(i-1)], Var2[i:]...)

	return reportVar1, reportVar2
}

func dampedReportCheck(report []int) bool {
	/*
		if reportCheck(report) {

			return true
		} else {

			for i := range report {

				reportTemp := make([]int, len(report))
				count := copy(reportTemp, report)
				if count == 0 {
					panic(count)
				}
				reportVariation := append(reportTemp[:i], reportTemp[(i+1):]...)

				if reportCheck(reportVariation) {

					return true
				} else {

				}
			}
		}
		return false
	*/
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
			if changeType != i && changeType != (0-i) {
				//fmt.Println("Fail changetype ")
				if i == 2 {
					//reportTemp := lineCopy(report)
					dvar := lineCopy(report)[1:]
					if reportCheck(dvar) {
						//fmt.Println(dvar)
						//fmt.Println("Safe")
						return true
					}

				}
				//reportTemp := lineCopy(report)
				//dvar := append(reportTemp[:(i-1)], reportTemp[i:]...)
				dvar, dvar2 := lineCopyDouble(report, i)
				if reportCheck(dvar) {
					//fmt.Println(dvar)
					//fmt.Println("Safe")
					return true
				}

				//reportTemp2 := lineCopy(report)
				//dvar2 := append(reportTemp2[:(i)], reportTemp2[(i+1):]...)
				if reportCheck(dvar2) {
					//fmt.Println(dvar)
					//fmt.Println("Safe")
					return true
				}

				return false
			}

		} else {
			/*fmt.Println("Fail change size")
			reportTemp := make([]int, len(report))
			count := copy(reportTemp, report)
			if count == 0 {
				panic(count)
			}


			dvar := append(reportTemp[:(i-1)], reportTemp[i:]...)
			*/
			dvar, dvar2 := lineCopyDouble(report, i)
			if reportCheck(dvar) {
				//fmt.Println(dvar)
				//fmt.Println("Safe")
				return true
			}
			/*
				reportTemp2 := make([]int, len(report))
				count2 := copy(reportTemp2, report)
				if count2 == 0 {
					panic(count2)
				}
				dvar2 := append(reportTemp2[:(i)], reportTemp2[(i+1):]...)
			*/

			if reportCheck(dvar2) {
				//fmt.Println(dvar)
				//fmt.Println("Safe")
				return true
			}

			return false
		}
	}
	//fmt.Println("Safe")
	return true

}

func Part1() (time.Duration, time.Duration, int) {
	Startfile := time.Now()
	var input = fileparse.FileParse("day2/Input.txt")
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
