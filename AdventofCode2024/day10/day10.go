package day10

import (
	"slices"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func trailScoring(mp [][]int, coords [][]int, elev int) int {
	trailCount := 0
	var nextCoords [][]int
	var dupes []string
	for _, i := range coords {
		stgCord := strconv.Itoa(i[0]) + strconv.Itoa(i[1])
		if mp[i[0]][i[1]] == 9 && !slices.Contains(dupes, stgCord) {
			//add to list of confirmed destinations
			dupes = append(dupes, stgCord)
			trailCount += 1
		} else if mp[i[0]][i[1]] != 99 {
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
	if len(nextCoords) == 0 && elev != 9 {
		return 0
	} else if elev == 9 {
		return trailCount
	} else {
		newElev := elev + 1
		trailCount = trailScoring(mp, nextCoords, newElev)
	}
	return trailCount
}

func trailRating(mp [][]int, coords [][]int, elev int) int {
	trailCount := 0
	var nextCoords [][]int
	//var dupes []string
	for _, i := range coords {
		//stgCord := strconv.Itoa(i[0]) + strconv.Itoa(i[1])
		if mp[i[0]][i[1]] == 9 {
			//add to list of confirmed destinations
			//dupes = append(dupes, stgCord)
			trailCount += 1
		} else if mp[i[0]][i[1]] != 99 {
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
	if len(nextCoords) == 0 && elev != 9 {
		return 0
	} else if elev == 9 {
		return trailCount
	} else {
		newElev := elev + 1
		trailCount = trailRating(mp, nextCoords, newElev)
	}
	return trailCount
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day10/Input.txt")
	//Insert additional processsing here
	var trailMap [][]int
	for _, i := range raw {
		temp := strings.Split(i, "")
		row := make([]int, 0, len(temp))
		for _, raw := range temp {
			v, _ := strconv.Atoi(raw)
			row = append(row, v)
		}
		//fmt.Println(row)
		trailMap = append(trailMap, row)
	}
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//insert Puzzle solution here

	chksum := 0
	for row, i := range trailMap {
		for col, j := range i {
			if trailMap[row][col] == 0 { //checkTrails
				var Startcords [][]int
				Startcords = append(Startcords, []int{row, col})
				chksum += trailScoring(trailMap, Startcords, j)
			}
		}
	}

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, chksum
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day10/Input.txt")
	//Insert additional processsing here
	var trailMap [][]int
	for _, i := range raw {
		temp := strings.Split(i, "")
		row := make([]int, 0, len(temp))
		for _, raw := range temp {
			v, _ := strconv.Atoi(raw)
			row = append(row, v)
		}
		//fmt.Println(row)
		trailMap = append(trailMap, row)
	}
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	//inerst Puzzle solution here
	chksum := 0
	for row, i := range trailMap {
		for col, j := range i {
			if trailMap[row][col] == 0 { //checkTrails
				var Startcords [][]int
				Startcords = append(Startcords, []int{row, col})
				chksum += trailRating(trailMap, Startcords, j)
			}
		}
	}

	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, chksum
}
