package day2

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func Rep(v string) bool {
	spl := len(v) / 2
	v1s := v[:spl]
	v2s := v[spl:]
	return v1s == v2s
}

func Splitter(v int) bool {
	var res bool
	res = false
	vs := strconv.Itoa(v)
	if len(vs)%2 == 0 {
		res = Rep(vs)
	}
	return res
}

func Ran(ln string) (int, int) {
	lims := strings.Split(ln, "-")
	low, _ := strconv.Atoi(lims[0])
	up, _ := strconv.Atoi(lims[1])
	return low, up

}

func Chunk(st string, sp int) []string {
	var va []string
	for i := 0; (i + sp) <= len(st); i = i + sp {
		va = append(va, st[i:(i+sp)])
	}
	return va
}

func RepAll(va []string) bool {
	for _, i := range va {
		if va[0] != i {
			return false
		}
	}
	return true
}

func SplitLevels(v int) bool {
	vs := strconv.Itoa(v)
	mx := len(vs) / 2
	for k := 1; k <= mx; k++ {
		if len(vs)%k == 0 {
			va := Chunk(vs, k)
			if RepAll(va) {
				return true
			}
		}
	}
	return false
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := strings.Split(fileparse.FileParse("day2/Input.txt")[0], ",")
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP1 := time.Now()

	for _, i := range Input {
		L, U := Ran(i)
		for j := L; j < (U + 1); j++ {
			if Splitter(j) {
				p1 = p1 + j
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
	Input := strings.Split(fileparse.FileParse("day2/Input.txt")[0], ",")

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	for _, i := range Input {
		L, U := Ran(i)
		for j := L; j < (U + 1); j++ {
			if SplitLevels(j) {
				p2 = p2 + j
			}
		}
	}
	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
