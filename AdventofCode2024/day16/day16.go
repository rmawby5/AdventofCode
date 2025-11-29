package day16

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func HeapI(hp [][]int, n []int) [][]int {
	j := len(hp)
	for i := range hp {
		if hp[i][2] > n[2] {
			j = i
			break
		}
	}
	return slices.Insert(hp, j, n)
}

func maze(mp [][]string, coords [][]int) int {
	var nextCoords [][]int
	nextCoords = append(nextCoords, coords[0])
	endWeight := 0
	atEnd := false

	for !atEnd {
		x := 0
		y := 1
		for k := 0; k < 4; k++ {
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
			if (i[0]+y) >= 0 && (i[0]+y) < len(mp) && (i[1]+x) >= 0 && (i[1]+x) < len(mp[i[0]]) { //in bounds
				if mp[i[0]+y][i[1]+x] == "." || mp[i[0]+y][i[1]+x] == "E" {
					curDir := i[3]
					curWeight := i[2]
					nextWeight := 0
					nextDir := (y * 1) + (x * 2)
					if curDir != nextDir {
						nextWeight = curWeight + 1001
					} else {
						nextWeight = curWeight + 1
					}
					newNode := []int{i[0] + y, i[1] + x, nextWeight, nextDir}
					nextCoords = HeapI(nextCoords, newNode)
				}
			}
			tmp := x
			x = y
			y = -tmp
		}
		nextCoords = slices.Delete(nextCoords, 0, 1)
	}
	return endWeight
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day16/TestInput.txt")
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
	sumP1 := maze(mp, start)
	fmt.Println("Mapped Maze: ")
	for i := range mp {
		fmt.Println(mp[i])
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sumP1
}

func seatCounter(mp [][]string, start []int, PathLen int) string {
	counter := ""
	//atStart := false

	coords := []int{start[0], start[1], start[2], start[3]}
	//mp[coords[0]][coords[1]] = "V"

	x := 0
	y := 1
	for k := 0; k < 4; k++ {
		splitCounter := ""
		if mp[coords[0]][coords[1]] == "E" {
			//atStart = true
			return strconv.Itoa(coords[0]) + "," + strconv.Itoa(coords[1]) + "|"

		}
		if mp[coords[0]][coords[1]] == "0" {
			mp[coords[0]][coords[1]] = strconv.Itoa(coords[2])
		}

		if (coords[0]+y) >= 0 && (coords[0]+y) < len(mp) && (coords[1]+x) >= 0 && (coords[1]+x) < len(mp[coords[0]]) { //in bounds
			//counter = 1
			curDir := coords[3]
			curWeight := coords[2]
			nextWeight := 0
			nextDir := (y * 1) + (x * 2)
			if curDir != nextDir {
				nextWeight = curWeight + 1001
			} else {
				nextWeight = curWeight + 1
			}

			if mp[coords[0]+y][coords[1]+x] == "0" || mp[coords[0]+y][coords[1]+x] == "E" {

				if nextWeight <= PathLen {
					nextCoord := []int{coords[0] + y, coords[1] + x, nextWeight, nextDir}
					tmp := seatCounter(mp, nextCoord, PathLen)
					if len(tmp) != 0 {
						splitCounter = strconv.Itoa(coords[0]) + "," + strconv.Itoa(coords[1]) + "|" + tmp
					}
				}

			}
		}
		tmp := x
		x = y
		y = -tmp
		counter += splitCounter
	}
	mpMod[coords[0]][coords[1]] = strconv.Itoa(coords[2])
	//mp[coords[0]][coords[1]] = "0"
	return counter

}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day16/TestInput.txt")
	var mp [][]string
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
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
	PathLen := maze(mp, start)

	startP2 := []int{start[0][0], start[0][1], 0, 2}

	sumP2 := seatCounter(mp, startP2, PathLen)

	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sumP2
}
