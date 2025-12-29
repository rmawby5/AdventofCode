package day8

import (
	"sort"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func CheckJunc(junctions [][]string, pair string) [][]string {
	nodes := strings.Split(pair, "/")
	n1 := nodes[0]
	n2 := nodes[1]
	n1p := false
	n2p := false
	n1i := 0
	n2i := 0

	for i, junc := range junctions {
		for _, node := range junc {
			if node == n1 {
				n1p = true
				n1i = i
			}
			if node == n2 {
				n2p = true
				n2i = i
			}
		}
	}

	if n1p && !n2p {
		junctions[n1i] = append(junctions[n1i], n2)
	}
	if !n1p && n2p {
		junctions[n2i] = append(junctions[n2i], n1)
	}
	if n1p && n2p && (n1i != n2i) {
		junctions[n1i] = append(junctions[n1i], junctions[n2i]...)
		junctions = append(junctions[:n2i], junctions[n2i+1:]...)
	}
	if !n1p && !n2p {
		junctions = append(junctions, []string{n1, n2})
	}
	return junctions
}

type NodePair struct {
	nodes string
	dist  int
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 1
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day8/Input.txt")
	var InputInt [][]int
	for _, i := range Input {
		var row []int
		for _, j := range strings.Split(i, ",") {
			c, _ := strconv.Atoi(j)
			row = append(row, c)
		}
		InputInt = append(InputInt, row)
	}

	ParseTime := time.Since(ParseStart)
	startP1 := time.Now()
	//puzzle
	var discache []NodePair
	var junctions [][]string
	for i := 0; i < (len(Input) - 1); i++ {
		for j := i + 1; j < len(Input); j++ {
			xd := InputInt[i][0] - InputInt[j][0]
			yd := InputInt[i][1] - InputInt[j][1]
			zd := InputInt[i][2] - InputInt[j][2]
			discache = append(discache, NodePair{Input[i] + "/" + Input[j], (xd * xd) + (yd * yd) + (zd * zd)})
		}
	}

	sort.Slice(discache, func(i int, j int) bool {
		return discache[i].dist < discache[j].dist
	})
	lim := 1000
	for i := 0; i < lim; i++ {
		junctions = CheckJunc(junctions, discache[i].nodes)
	}

	for jCount := 0; jCount < 3; jCount++ {
		idx := 0
		min := 0
		for j, i := range junctions {
			jSize := len(i)
			if jSize > min {
				min = jSize
				idx = j
			}
		}
		junctions = append(junctions[:idx], junctions[idx+1:]...)
		p1 = p1 * min
	}

	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	p2 := 1
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day8/Input.txt")
	var InputInt [][]int
	for _, i := range Input {
		var row []int
		for _, j := range strings.Split(i, ",") {
			c, _ := strconv.Atoi(j)
			row = append(row, c)
		}
		InputInt = append(InputInt, row)
	}

	ParseTime := time.Since(ParseStart)
	startP2 := time.Now()
	//puzzle
	var discache []NodePair
	var junctions [][]string

	for i := 0; i < (len(Input) - 1); i++ {
		for j := i + 1; j < len(Input); j++ {
			xd := InputInt[i][0] - InputInt[j][0]
			yd := InputInt[i][1] - InputInt[j][1]
			zd := InputInt[i][2] - InputInt[j][2]
			discache = append(discache, NodePair{Input[i] + "/" + Input[j], (xd * xd) + (yd * yd) + (zd * zd)})
		}
	}

	sort.Slice(discache, func(i int, j int) bool {
		return discache[i].dist < discache[j].dist
	})

	junctions = CheckJunc(junctions, discache[0].nodes)
	i := 0
	for len(junctions[0]) != 1000 {
		i++
		junctions = CheckJunc(junctions, discache[i].nodes)
	}

	FinalPair := strings.Split(discache[i].nodes, ",")
	x1, _ := strconv.Atoi(FinalPair[0])
	x2, _ := strconv.Atoi(strings.Split(FinalPair[2], "/")[1])
	p2 = x1 * x2

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
