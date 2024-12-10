package day10

import (
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func trailCheck(mp [][]string, coords [][]int, elev int) int {
	trailCount := 0
	var nextCoords [][]int
	for _, i := range coords {
		if mp[i[0]][i[1]] == 9 {
			trailCount += 1
		} else {

			//check up
			if (i[0] - 1) >= 0 { //in bounds
				if mp[i[0]-1][i[1]] == (elev + 1) {

					nextCoords = append(nextCoords, []int{i[0] - 1, i[1]})
				}

			}

			//check down
			if (i[0] + 1) < len(mp) { //in bounds
				if mp[i[0]+1][i[1]] == (elev + 1) {

					nextCoords = append(nextCoords, []int{i[0] + 1, i[1]})
				}

			}

			//check left
			if (i[1] - 1) >= 0 { //in bounds
				if mp[i[0]][i[1]-1] == (elev + 1) {

					nextCoords = append(nextCoords, []int{i[0], i[1] - 1})
				}

			}

			//check right
			if (i[1] + 1) < len(mp[i[0]]) { //in bounds
				if mp[i[0]][i[1]+1] == (elev + 1) {

					nextCoords = append(nextCoords, []int{i[0], i[1] + 1})
				}

			}
		}
	}

	if len(nextCoords) == 0 && elev != 9{
		return 0
	} else if elev == 9 {
		return trailCount
	} else {
		newElev := elev + 1
		trailCount = trailCheck(mp, nextCoords, newElev)
	}

	return trailCount
}

func Part1() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day10/TestInput.txt")
	//Insert additional processsing here
	var trailMap [][]string
	for _, i := range raw {
		trailMap = append(trailMap, strings.Split(i, ""))
	}
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//inerst Puzzle solution here
	var chksum int64
	chksum = 0
	for row, i := range trailMap {
		for col, j := range i {
			if trailMap[row][col] == "0" { //checkTrails
				Startcords := [][]int
				Startcords = append(Startcords, []int{row, col})
				chksum += trailCheck(mp,Startcords,0)
			}
		}
	}

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, chksum
}

func Part1() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day10/TestInput.txt")
	//Insert additional processsing here

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//inerst Puzzle solution here
	var chksum int64
	chksum = 0

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, chksum
}
