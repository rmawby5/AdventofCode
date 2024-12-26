package day8

import (
	"slices"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func nodePos(p1 []int, p2 []int) (string, string) {
	yDiff := p2[0] - p1[0]
	xDiff := p2[1] - p1[1]
	xLim := 49
	yLim := 49
	an1 := ""
	an2 := ""
	if p1[0]-yDiff < 0 || p1[1]-xDiff < 0 || p1[1]-xDiff > xLim {
		an1 = ""
	} else {
		an1 = strconv.Itoa(p1[0]-yDiff) + "," + strconv.Itoa(p1[1]-xDiff)
	}
	if p2[0]+yDiff > yLim || p2[1]+xDiff < 0 || p2[1]+xDiff > xLim {
		an2 = ""
	} else {
		an2 = strconv.Itoa(p2[0]+yDiff) + "," + strconv.Itoa(p2[1]+xDiff)
	}
	return an1, an2
}

func antiNodes(nodes string, anList []string) []string {
	var coords [][]int
	tmp := strings.Split(nodes, ",")[:len(strings.Split(nodes, ","))-1]
	for _, i := range tmp {
		point := strings.Split(i, "-")
		Py, _ := strconv.Atoi(point[0])
		Px, _ := strconv.Atoi(point[1])
		Pi := []int{Py, Px}
		coords = append(coords, Pi)
	}
	for i := 0; i <= len(coords)-2; i++ {
		for j := i + 1; j < len(coords); j++ {
			an1, an2 := nodePos(coords[i], coords[j])
			if !slices.Contains(anList, an1) && an1 != "" {
				anList = append(anList, an1)
			}
			if !slices.Contains(anList, an2) && an2 != "" {
				anList = append(anList, an2)
			}
		}
	}
	return anList
}

func nodePosAll(p1 []int, p2 []int, anList []string) []string {
	yDiff := p2[0] - p1[0]
	xDiff := p2[1] - p1[1]
	xLim := 49
	yLim := 49
	an := ""
	inGrid := true
	nextNode := []int{p1[0] - yDiff, p1[1] - xDiff}
	for inGrid {
		if nextNode[0] < 0 || nextNode[1] < 0 || nextNode[1] > xLim {
			inGrid = false
		} else {
			an = strconv.Itoa(nextNode[0]) + "," + strconv.Itoa(nextNode[1])
			if !slices.Contains(anList, an) && an != "" {
				anList = append(anList, an)
			}
			nextNode[0] = nextNode[0] - yDiff
			nextNode[1] = nextNode[1] - xDiff
		}
	}
	inGrid = true
	nextNode = []int{p2[0] + yDiff, p2[1] + xDiff}
	for inGrid {
		if nextNode[0] > yLim || nextNode[1] < 0 || nextNode[1] > xLim {
			inGrid = false
		} else {
			an = strconv.Itoa(nextNode[0]) + "," + strconv.Itoa(nextNode[1])
			if !slices.Contains(anList, an) && an != "" {
				anList = append(anList, an)
			}
			nextNode[0] = nextNode[0] + yDiff
			nextNode[1] = nextNode[1] + xDiff
		}
	}
	return anList
}

func antiNodesAll(nodes string, anList []string) []string {
	var coords [][]int
	tmp := strings.Split(nodes, ",")[:len(strings.Split(nodes, ","))-1]
	for _, i := range tmp {
		point := strings.Split(i, "-")
		Py, _ := strconv.Atoi(point[0])
		Px, _ := strconv.Atoi(point[1])
		Pi := []int{Py, Px}
		coords = append(coords, Pi)
	}
	for i := 0; i <= len(coords)-2; i++ {
		for j := i + 1; j < len(coords); j++ {
			anList = nodePosAll(coords[i], coords[j], anList)
		}
	}
	return anList
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day8/Input.txt")

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	//Insert additional processsing here
	input := make(map[string]string)
	for y, i := range raw {
		tmp := strings.Split(i, "")
		for x, j := range tmp {
			if j != "." {
				input[j] = input[j] + strconv.Itoa(y) + "-" + strconv.Itoa(x) + ","
			}
		}
	}
	//insert Puzzle solution here
	var nodes []string
	for _, i := range input {
		nodes = antiNodes(i, nodes)
	}
	sum := len(nodes)

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day8/Input.txt")

	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()

	//Insert additional processsing here
	var nodes []string
	input := make(map[string]string)
	for y, i := range raw {
		tmp := strings.Split(i, "")
		for x, j := range tmp {
			if j != "." {
				nodes = append(nodes, strconv.Itoa(y)+","+strconv.Itoa(x))
				input[j] = input[j] + strconv.Itoa(y) + "-" + strconv.Itoa(x) + ","
			}
		}
	}
	//insert Puzzle solution here
	for _, i := range input {
		nodes = antiNodesAll(i, nodes)
	}
	sum := len(nodes)

	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sum
}
