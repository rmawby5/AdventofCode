package day7

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func eqCheck(Tot int, current int, comp []int) bool {
	var success bool
	success = false //default value unless eq is proven correct
	//try multiple
	remaining := len(comp)
	total := current * comp[0]
	if remaining == 1 { //end of components
		if total != Tot {
			success = false
		} else if total == Tot {
			success = true
		}
	} else {
		success = eqCheck(Tot, total, comp[1:])
	}
	if !success { //try addition
		total = current + comp[0]
		if remaining == 1 { //end of components
			if total != Tot {
				success = false
			} else if total == Tot {
				success = true
			}
		} else {
			success = eqCheck(Tot, total, comp[1:])
		}
	}
	return success
}

func concat(a int, b int) int {
	c, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return c
}

func eqCheckExpanded(Tot int, current int, comp []int) bool {
	var success bool
	success = false //default value unless eq is proven correct
	//try multiple
	remaining := len(comp)
	total := current * comp[0]
	if remaining == 1 { //end of components
		if total != Tot {
			success = false
		} else if total == Tot {
			success = true
		}
	} else {
		success = eqCheckExpanded(Tot, total, comp[1:])
	}
	if !success { //try with addition
		total = current + comp[0]
		if remaining == 1 { //end of components
			if total != Tot {
				success = false
			} else if total == Tot {
				success = true
			}
		} else {
			success = eqCheckExpanded(Tot, total, comp[1:])
		}
	}
	if !success { //try with concat
		total = concat(current, comp[0])
		if remaining == 1 { //end of components
			if total != Tot {
				success = false
			} else if total == Tot {
				success = true
			}
		} else {
			success = eqCheckExpanded(Tot, total, comp[1:])
		}
	}
	return success
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day7/Input.txt")
	var components [][]int
	var Totals []int
	for _, i := range raw {
		var lineCompents []int
		line := strings.Split(i, ": ")
		lineTotal, _ := strconv.Atoi(line[0])
		Totals = append(Totals, lineTotal)
		for _, j := range strings.Split(line[1], " ") {
			comp, _ := strconv.Atoi(j)
			lineCompents = append(lineCompents, comp)
		}
		components = append(components, lineCompents)
	}
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	var calibratrion int
	calibratrion = 0
	for i := range Totals {
		if eqCheck(Totals[i], components[i][0], components[i][1:]) {
			calibratrion += Totals[i]
		}
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, calibratrion
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day7/Input.txt")
	var components [][]int
	var Totals []int
	for _, i := range raw {
		var lineCompents []int
		line := strings.Split(i, ": ")
		lineTotal, _ := strconv.Atoi(line[0])
		Totals = append(Totals, lineTotal)
		for _, j := range strings.Split(line[1], " ") {
			comp, _ := strconv.Atoi(j)

			lineCompents = append(lineCompents, comp)
		}
		components = append(components, lineCompents)
	}
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	var calibratrion int
	calibratrion = 0
	for i := range Totals {
		if eqCheckExpanded(Totals[i], components[i][0], components[i][1:]) {
			calibratrion += Totals[i]
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, calibratrion
}
