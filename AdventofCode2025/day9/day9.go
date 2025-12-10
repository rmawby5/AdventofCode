package day9

import (
	"fmt"
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

func Area(p1 []int, p2 []int) int {

	return (ab(p1[0]-p2[0]) + 1) * (ab(p1[1]-p2[1]) + 1)

}

func Intersect(c1 []int, c2 []int, p1 []int, p2 []int) bool {
	c3:= []int{c1[0],c2[1]}
	c4:=[]int{c2[0], c1[1]}

	//check intersection between c1, c3
	if  {
		fmt.Println("X P1")
		return true
	}
	//check intersection between c1, c4
	if  {
		fmt.Println("X P2")
		return true
	}
	//check intersection between c2, c3
	if  {
		fmt.Println("Y P1")
		return true
	}
	//check intersection between c2, c4
	if  {
		fmt.Println("Y p2")
		return true
	}

	return false

}

	fmt.Println(Intersect([]int{7, 1}, []int{11, 1}, []int{9, 5}, []int{2, 5}))

func IntersectValidation(p1 []int, p2 []int, pList [][]int) bool {
	//var pv bool
	//	for i := 0; i < (len(pList) - 1); i++ {
	//		if Intersect(p1, p2, pList[i], pList[i+1]) {
	//			//intersection is present, area is not valid
	//			fmt.Println("false")
	//			fmt.Println(i)
	//			fmt.Print(pList[i])
	//			fmt.Print(" : ")
	//			fmt.Println(pList[i+1])
	//
	//			return false
	//		}
	//
	//	}
	for i := 0; i < (len(pList) - 1); i++ {
		for j := i + 1; j < len(pList); j++ {
			if Intersect(p1, p2, pList[i], pList[i+1]) {
				//intersection is present, area is not valid
				fmt.Println("false")
				fmt.Println(i)
				fmt.Print(pList[i])
				fmt.Print(" : ")
				fmt.Println(pList[i+1])

				return false
			}
		}
	}

	return true
}

func BoundaryCheckList(p1 []int, p2 []int, pList [][]int) bool {
	//var pv bool
	for _, i := range pList {
		if !BoundaryCheck(p1, p2, i) {
			return false
		}

	}
	return true
}

func BoundaryCheck(p1 []int, p2 []int, p3 []int) bool {
	//checks if p3 is within a rectangle formed with p2 and p1 as diagonal points
	xdiff := p1[0] - p2[0]
	ydiff := p1[1] - p2[1]

	if xdiff <= 0 {
		//p2 right
		if (p3[0]-p2[0]) >= 0 || (p3[0]-p1[0] <= 0) {
			return true

		}
	} else {
		//p2 left
		if (p3[0]-p2[0]) <= 0 || (p3[0]-p1[0] >= 0) {
			return true
		}
	}

	if ydiff <= 0 {
		//p2 higher
		if (p3[1]-p2[1]) >= 0 || (p3[1]-p1[1] <= 0) {
			return true
		}
	} else {
		//p2 lower
		if (p3[1]-p2[1]) <= 0 || (p3[1]-p1[1] >= 0) {
			return true
		}
	}
	//fmt.Print(p3)
	//fmt.Print(" : ")
	//fmt.Println(pv)
	return false
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day9/TestInput.txt")
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
			curArea := Area(InputInt[i], InputInt[j])
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
	Input := fileparse.FileParse("day9/TestInput.txt")
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

	//	for i := 0; i < len(InputInt)-1; i++ {
	//		for j := i + 1; j < len(InputInt); j++ {
	//			curArea := Area(InputInt[i], InputInt[j])
	//			if curArea > p2 {
	//
	//				//fmt.Print(InputInt[i])
	//				//fmt.Print(" : ")
	//				//fmt.Println(InputInt[j])
	//				//fmt.Println(curArea)
	//				//fmt.Println(IntersectValidation(InputInt[i], InputInt[j], InputInt))
	//				if IntersectValidation(InputInt[i], InputInt[j], InputInt) {
	//
	//					p2 = curArea
	//				}
	//
	//			}
	//		}
	//	}
	//fmt.Println(IntersectValidation([]int{7, 1}, []int{11, 1}, InputInt))
	//fmt.Println(BoundaryCheckList([]int{9, 5}, []int{2, 3}, InputInt))

	fmt.Println(Intersect([]int{7, 1}, []int{11, 1}, []int{9, 5}, []int{2, 5}))
	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
