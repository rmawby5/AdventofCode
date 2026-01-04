package day9

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func ab(num int) int {
	if num < 0 {
		return (0 - num)
	}
	return num
}

func Intersect(p1 []int, p2 []int, c1 []int, c2 []int) bool {
	//p=area definition, c = line for intersect check
	//checks if line created by c1/c2 intersects with the borders of an area defined by diagonal points p1/p2
	if c1[0] == c2[0] {
		// verticle line
		if (c1[0] < p1[0] && c1[0] > p2[0]) || (c1[0] > p1[0] && c1[0] < p2[0]) {
			switch h := p1[1] < p2[1]; h {
			case true:
				//p2 higher
				if (c1[1] > p2[1] && c2[1] > p2[1]) || (c1[1] < p1[1] && c2[1] < p1[1]) {
					//both points are either lower or higher than the area, hence do not intersect
					return false
				}
			case false:
				//p1 higher
				if (c1[1] > p1[1] && c2[1] > p1[1]) || (c1[1] < p2[1] && c2[1] < p2[1]) {
					//both points are either lower or higher than the area, hence do not intersect
					return false
				}
			}
			//line intersects boundary at some point
			return true
		}
	} else if c1[1] == c2[1] {
		//horzontal line
		if (c1[1] < p1[1] && c1[1] > p2[1]) || (c1[1] > p1[1] && c1[1] < p2[1]) {
			switch h := p1[0] < p2[0]; h {
			case true:
				//p2 is higher
				if (c1[0] > p2[0] && c2[0] > p2[0]) || (c1[0] < p1[0] && c2[0] < p1[0]) {
					//both points are either lower or higher than the area, hence do not intersect
					return false
				}
			case false:
				//p1 is higher
				if (c1[0] > p1[0] && c2[0] > p1[0]) || (c1[0] < p2[0] && c2[0] < p2[0]) {
					//both points are either lower or higher than the area, hence do not intersect
					return false
				}
			}
			//line intersects boundary at some point
			return true
		}
	}
	return false
}

func IntersectValidation(p1 []int, p2 []int, pList [][]int) bool {
	//gets line segments from a list of points and returns true is no line segments intersect with the boundy of area defined by p1/p2
	for i := 0; i < (len(pList) - 1); i++ {
		if Intersect(p1, p2, pList[i], pList[i+1]) {
			//there is a line intersection, area is not valid
			return false
		}
	}
	return true
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day9/Input.txt")
	var InputInt [][]int
	for _, i := range Input {
		row := strings.Split(i, ",")
		v0, _ := strconv.Atoi(row[0])
		v1, _ := strconv.Atoi(row[1])
		InputInt = append(InputInt, []int{v0, v1})
	}
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP1 := time.Now()
	for i := 0; i < len(InputInt)-1; i++ {
		for j := i + 1; j < len(InputInt); j++ {
			curArea := (ab(InputInt[i][0]-InputInt[j][0]) + 1) * (ab(InputInt[i][1]-InputInt[j][1]) + 1)
			if curArea > p1 {
				p1 = curArea
			}
		}
	}
	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	p2 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day9/Input.txt")
	var InputInt [][]int
	for _, i := range Input {
		row := strings.Split(i, ",")
		v0, _ := strconv.Atoi(row[0])
		v1, _ := strconv.Atoi(row[1])
		InputInt = append(InputInt, []int{v0, v1})
	}
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	for i := 0; i < len(InputInt)-1; i++ {
		for j := i + 1; j < len(InputInt); j++ {
			curArea := (ab(InputInt[i][0]-InputInt[j][0]) + 1) * (ab(InputInt[i][1]-InputInt[j][1]) + 1)
			if curArea > p2 {
				if IntersectValidation(InputInt[i], InputInt[j], InputInt) {
					p2 = curArea
				}
			}
		}
	}
	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
