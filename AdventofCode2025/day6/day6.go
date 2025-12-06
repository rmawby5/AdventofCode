package day6

import (
	"strconv"
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
			//fmt.Println(num)
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
	//fmt.Println(total)
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
		//fmt.Println(comp[len(comp)-1])
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
	var InputGrid [][]string

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
