package day16

import (
	"fmt"
	"slices"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func HeapI(hp [][]int, n []int) [][]int {
	//fmt.Println("")
	//fmt.Println("Current Nodes : ", hp)
	//fmt.Println("New Node : ", n)

	//hp_N := make([][]int, len(hp)+1)
	//fmt.Println(hp)
	j := len(hp)
	for i := range hp {
		if hp[i][2] > n[2] {
			j = i
			//fmt.Println("Insert index", j)
			break
		}
	}

	hp_N := slices.Insert(hp, j, n)
	//fmt.Println("New node list : ", hp_N)

	return hp_N
}

func maze(mp [][]string, coords [][]int) (int, int) {
	var nextCoords [][]int
	nextCoords = append(nextCoords, coords[0])
	endWeight := 0
	atEnd := false
	seatCount := 0

	for !atEnd {
		i := nextCoords[0]
		//check up
		if mp[i[0]][i[1]] == "E" {
			atEnd = true
			endWeight = i[2]
			break
		}
		if mp[i[0]][i[1]] == "." {
			mp[i[0]][i[1]] = "0"
		}
		if (i[0] - 1) >= 0 { //in bounds
			if mp[i[0]-1][i[1]] == "." || mp[i[0]-1][i[1]] == "E" {
				curDir := i[3]
				curWeight := i[2]
				nextWeight := 0
				if curDir == 270 || curDir == 90 {
					nextWeight = curWeight + 1001
				} else {
					nextWeight = curWeight + 1
				}
				newNode := []int{i[0] - 1, i[1], nextWeight, 0}
				nextCoords = HeapI(nextCoords, newNode)
			}

		}
		//check down
		if (i[0] + 1) < len(mp) { //in bounds
			if mp[i[0]+1][i[1]] == "." || mp[i[0]+1][i[1]] == "E" {
				curDir := i[3]
				curWeight := i[2]
				nextWeight := 0
				if curDir == 270 || curDir == 90 {
					nextWeight = curWeight + 1001
				} else {
					nextWeight = curWeight + 1
				}
				newNode := []int{i[0] + 1, i[1], nextWeight, 180}
				nextCoords = HeapI(nextCoords, newNode)
			}
		}
		//check left
		if (i[1] - 1) >= 0 { //in bounds
			if mp[i[0]][i[1]-1] == "." || mp[i[0]][i[1]-1] == "E" {
				curDir := i[3]
				curWeight := i[2]
				nextWeight := 0
				if curDir == 180 || curDir == 0 {
					nextWeight = curWeight + 1001
				} else {
					nextWeight = curWeight + 1
				}
				newNode := []int{i[0], i[1] - 1, nextWeight, 270}
				nextCoords = HeapI(nextCoords, newNode)
			}
		}
		//check right
		if (i[1] + 1) < len(mp[i[0]]) { //in bounds
			if mp[i[0]][i[1]+1] == "." || mp[i[0]][i[1]+1] == "E" {
				curDir := i[3]
				curWeight := i[2]
				nextWeight := 0
				if curDir == 0 || curDir == 180 {
					nextWeight = curWeight + 1001
				} else {
					nextWeight = curWeight + 1
				}
				newNode := []int{i[0], i[1] + 1, nextWeight, 90}
				nextCoords = HeapI(nextCoords, newNode)
			}
		}
		nextCoords = slices.Delete(nextCoords, 0, 1)
	}

	return endWeight, seatCount
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day16/Input.txt")
	var mp [][]string
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	start := make([][]int, 1)
	//Insert additional processsing here
	for row, i := range raw {
		mp = append(mp, strings.Split(i, ""))
		if strings.Contains(i, "S") {
			tmp := make([]int, 4)
			tmp[0] = row
			tmp[1] = strings.Index(i, "S")
			tmp[2] = 0
			tmp[3] = 90
			start[0] = tmp
		}
	}
	//insert Puzzle solution here
	sumP1, sumP2 := maze(mp, start)
	fmt.Println("Part 2: ", sumP2)
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sumP1
}
