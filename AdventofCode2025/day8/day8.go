package day8

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func Distance(p1 string, p2 string) int {

	p1c := strings.Split(p1, ",")
	p2c := strings.Split(p2, ",")
	//fmt.Println(p1c)
	//fmt.Println(p2c)
	//x
	p1x, _ := strconv.Atoi(p1c[0])
	p2x, _ := strconv.Atoi(p2c[0])
	xd := p1x - p2x
	//fmt.Println(xd)
	//y
	p1y, _ := strconv.Atoi(p1c[1])
	p2y, _ := strconv.Atoi(p2c[1])
	yd := p1y - p2y
	//fmt.Println(yd)
	//z
	p1z, _ := strconv.Atoi(p1c[2])
	p2z, _ := strconv.Atoi(p2c[2])
	zd := p1z - p2z
	//fmt.Println(zd)

	return ((xd * xd) + (yd * yd) + (zd * zd))

}

func CheckJunc(junc [][]int, node int) int {
	for k, i := range junc {
		for j := range i {
			if node == j {
				//node is already present in a junction, return index of junction
				return k
			}
		}
	}
	return -1
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day8/TestInput.txt")

	ParseTime := time.Since(ParseStart)
	//puzzle

	var discache [][]string
	lim := 10
	for i := 0; i < (len(Input) - 1); i++ {
		for j := i + 1; j < len(Input); j++ {

			dis := []string{Input[i], Input[j], strconv.Itoa(Distance(Input[i], Input[j]))}
			discache = append(discache, dis)
		}
	}

	startP1 := time.Now()

	//fmt.Println(len(discache))
	//fmt.Println(Distance("162,817,812", "431,825,988"))
	for i := 0; i < lim; i++ {
		min := 99999999999
		//var cn []string
		var idx int
		for j, k := range discache {
			d, _ := strconv.Atoi(k[2])
			if d < min {
				//cn = k
				min = d
				idx = j

			}

		}
		discache = append(discache[:idx], discache[idx+1:]...)
		//fmt.Println(cn)
		//fmt.Println(len(discache))
	}

	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	p2 := 0
	//parse function
	ParseStart := time.Now()
	//Input := fileparse.FileParse("day9/TestInput.txt")

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
