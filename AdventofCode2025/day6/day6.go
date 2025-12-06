package day6

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func Sep(row string) []string {
	var rowSc []string
	var num string
	row = row + " "
	for i := 0; i < len(row); i++ {
		if row[i] != ' ' {
			num = num + string(row[i])
		}
		if row[i] == ' ' && len(num) != 0 {
			rowSc = append(rowSc, num)
			num = ""
		}
	}
	return rowSc
}

func Calc(comp []string) int {
	op := comp[len(comp)-1]
	var total int
	switch op {
	case "+":
		for i := 0; i < (len(comp) - 1); i++ {
			j, _ := strconv.Atoi(comp[i])
			total += j
		}
	case "*":
		total = 1
		for i := 0; i < (len(comp) - 1); i++ {
			j, _ := strconv.Atoi(comp[i])
			total = total * j
		}
	}
	return total
}

func CalcMod(comp []string, op byte) int {
	size := len(comp[0])
	total := 0
	switch op {
	case '+':
		for i := 0; i < size; i++ {
			var curValST string
			for j := 0; j < len(comp); j++ {
				curValST = curValST + string(comp[j][i])
			}
			curVal, _ := strconv.Atoi(strings.Trim(curValST, " "))
			total += curVal
		}
	case '*':
		total = 1
		for i := 0; i < size; i++ {
			var curValST string
			for j := 0; j < len(comp); j++ {
				curValST = curValST + string(comp[j][i])
			}
			curVal, _ := strconv.Atoi(strings.Trim(curValST, " "))
			total = total * curVal
		}
	}
	return total
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day6/Input.txt")
	var InputGrid [][]string
	for _, i := range Input {
		InputGrid = append(InputGrid, Sep(i))

	}
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP1 := time.Now()
	for i := range InputGrid[0] {
		var comp []string
		for j := range InputGrid {
			comp = append(comp, InputGrid[j][i])
		}
		p1 += Calc(comp)
	}
	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	p2 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day6/Input.txt")

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	cs := 0
	ce := 0
	opRow := len(Input) - 1
	for i := 1; i < len(Input[opRow]); i++ {
		if Input[opRow][i] != ' ' {
			ce = i - 1
			//puzzle processing
			var comp []string
			for j := 0; j < opRow; j++ {
				comp = append(comp, Input[j][cs:ce])
			}
			p2 += CalcMod(comp, Input[opRow][cs])
			cs = i
		}
	}

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
