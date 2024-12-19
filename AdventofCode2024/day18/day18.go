package day18

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func path(queue [][]int, mp [][]string, step int, goal string) int {
	pathLength := 0
	var newQueue [][]int
	newmp := make([][]string, len(mp))
	_ = copy(newmp, mp)

	for _, i := range queue {
		if strconv.Itoa(i[0])+","+strconv.Itoa(i[1]) == goal {
			pathLength = step
			break
		} else {
			if (i[0] - 1) >= 0 { //in bounds
				if newmp[i[0]-1][i[1]] == "." {
					newQueue = append(newQueue, []int{i[0] - 1, i[1]})
					newmp[i[0]-1][i[1]] = strconv.Itoa(step + 1)
				}
			}
			if (i[0] + 1) < len(mp) { //in bounds
				if newmp[i[0]+1][i[1]] == "." {
					newQueue = append(newQueue, []int{i[0] + 1, i[1]})
					newmp[i[0]+1][i[1]] = strconv.Itoa(step + 1)
				}
			}
			if (i[1] - 1) >= 0 { //in bounds
				if newmp[i[0]][i[1]-1] == "." {
					newQueue = append(newQueue, []int{i[0], i[1] - 1})
					newmp[i[0]][i[1]-1] = strconv.Itoa(step + 1)
				}
			}
			if (i[1] + 1) < len(mp[i[0]]) { //in bounds
				if newmp[i[0]][i[1]+1] == "." {
					newQueue = append(newQueue, []int{i[0], i[1] + 1})
					newmp[i[0]][i[1]+1] = strconv.Itoa(step + 1)
				}
			}
		}
	}
	if len(newQueue) == 0 && pathLength == 0 {
		return -1
	} else if pathLength != 0 {
		return pathLength
	} else {
		pathLength = path(newQueue, newmp, (step + 1), goal)
	}
	return pathLength
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day18/Input.txt")

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	var byteBlocks [][]int

	//Insert additional processsing here
	for i := 0; i < 1024; i++ {
		tmp := strings.Split(raw[i], ",")
		r, _ := strconv.Atoi(tmp[1])
		c, _ := strconv.Atoi(tmp[0])
		tmpI := []int{r, c}
		byteBlocks = append(byteBlocks, tmpI)
	}
	row := make([]string, 71)
	mp := make([][]string, 71)
	for i := range row {
		row[i] = "."
	}
	for i := range mp {
		tmp := make([]string, len(row))
		_ = copy(tmp, row)
		mp[i] = tmp
	}
	for _, i := range byteBlocks {
		mp[i[0]][i[1]] = "#"
	}
	//insert Puzzle solution here
	var queue [][]int
	queue = append(queue, []int{0, 0})
	goal := "70,70"
	mp[0][0] = "0"
	sum := path(queue, mp, 0, goal)
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, string) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day18/Input.txt")

	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()

	var failByte string
	for j := 1024; j <= len(raw); j++ {
		var byteBlocks [][]int
		row := make([]string, 71)
		mp := make([][]string, 71)
		for i := range row {
			row[i] = "."
		}
		for i := range mp {
			tmp := make([]string, len(row))
			_ = copy(tmp, row)
			mp[i] = tmp
		}
		var queue [][]int
		queue = append(queue, []int{0, 0})
		goal := "70,70"
		mp[0][0] = "0"
		for i := 0; i < j; i++ {
			tmp := strings.Split(raw[i], ",")
			r, _ := strconv.Atoi(tmp[1])
			c, _ := strconv.Atoi(tmp[0])
			tmpI := []int{r, c}
			byteBlocks = append(byteBlocks, tmpI)
		}
		for _, i := range byteBlocks {
			mp[i[0]][i[1]] = "#"
		}
		if path(queue, mp, 0, goal) == -1 {
			failByte = strconv.Itoa(byteBlocks[len(byteBlocks)-1][1]) + "," + strconv.Itoa(byteBlocks[len(byteBlocks)-1][0])
			break
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, failByte
}
