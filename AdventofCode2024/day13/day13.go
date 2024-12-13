package day13

import (
	"math"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func eqSolver(a float64, b float64, c float64, d float64, Xt float64, Yt float64) (float64, float64) {
	f := (Xt - ((a * Yt) / b)) / (c - ((a * d) / b))
	e := (Xt - (c * f)) / a
	return e, f
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day13/Input.txt")
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//Insert additional processsing here
	var input [][]string
	for _, i := range raw {
		input = append(input, strings.Split(i, " "))
	}
	//insert Puzzle solution here
	sum := 0
	for j := 0; j < len(input); j += 4 {

		a, _ := strconv.ParseFloat(input[j][2][2:len(input[j][2])-1], 64)
		b, _ := strconv.ParseFloat(input[j][3][2:], 64)
		c, _ := strconv.ParseFloat(input[j+1][2][2:len(input[j+1][2])-1], 64)
		d, _ := strconv.ParseFloat(input[j+1][3][2:], 64)
		Xt, _ := strconv.ParseFloat(input[j+2][1][2:len(input[j+2][1])-1], 64)
		Yt, _ := strconv.ParseFloat(input[j+2][2][2:], 64)
		e, f := eqSolver(a, b, c, d, Xt, Yt)
		if math.Abs(math.Round(e)-e) <= 0.0001 && math.Abs(math.Round(f)-f) <= 0.0001 {
			sum += int((math.Round(e) * 3) + (math.Round(f) * 1))
		}
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day13/Input.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	//Insert additional processsing here
	var input [][]string
	for _, i := range raw {
		input = append(input, strings.Split(i, " "))
	}
	//insert Puzzle solution here
	sum := 0
	for j := 0; j < len(input); j += 4 {
		a, _ := strconv.ParseFloat(input[j][2][2:len(input[j][2])-1], 64)
		b, _ := strconv.ParseFloat(input[j][3][2:], 64)
		c, _ := strconv.ParseFloat(input[j+1][2][2:len(input[j+1][2])-1], 64)
		d, _ := strconv.ParseFloat(input[j+1][3][2:], 64)
		Xt, _ := strconv.ParseFloat(input[j+2][1][2:len(input[j+2][1])-1], 64)
		Yt, _ := strconv.ParseFloat(input[j+2][2][2:], 64)
		e, f := eqSolver(a, b, c, d, Xt+10000000000000, Yt+10000000000000)
		if math.Abs(math.Round(e)-e) <= 0.0001 && math.Abs(math.Round(f)-f) <= 0.0001 {
			sum += int((math.Round(e) * 3) + (math.Round(f) * 1))
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sum
}
