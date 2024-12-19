package day12

import (
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func regionMapper(mp [][]string, coords [][]int, crop string) (int, int) {
	area := 0
	perimeter := 0
	var nextCoords [][]int
	for _, i := range coords {
		if mp[i[0]][i[1]] != ("#" + crop) { //don't revist node
			area += 1
			mp[i[0]][i[1]] = ("#" + crop)
			//check up
			if (i[0] - 1) >= 0 { //in bounds
				if mp[i[0]-1][i[1]] == crop {
					nextCoords = append(nextCoords, []int{i[0] - 1, i[1]})
				} else if mp[i[0]-1][i[1]] != ("#" + crop) {
					perimeter += 1
				}
			} else {
				perimeter += 1
			}
			//check down
			if (i[0] + 1) < len(mp) { //in bounds
				if mp[i[0]+1][i[1]] == crop {
					nextCoords = append(nextCoords, []int{i[0] + 1, i[1]})
				} else if mp[i[0]+1][i[1]] != ("#" + crop) {
					perimeter += 1
				}
			} else {
				perimeter += 1
			}
			//check left
			if (i[1] - 1) >= 0 { //in bounds
				if mp[i[0]][i[1]-1] == crop {
					nextCoords = append(nextCoords, []int{i[0], i[1] - 1})
				} else if mp[i[0]][i[1]-1] != ("#" + crop) {
					perimeter += 1
				}
			} else {
				perimeter += 1
			}
			//check right
			if (i[1] + 1) < len(mp[i[0]]) { //in bounds
				if mp[i[0]][i[1]+1] == crop {
					nextCoords = append(nextCoords, []int{i[0], i[1] + 1})
				} else if mp[i[0]][i[1]+1] != ("#" + crop) {
					perimeter += 1
				}
			} else {
				perimeter += 1
			}
		}
	}

	if len(nextCoords) != 0 {
		areaEx, perimeterEX := regionMapper(mp, nextCoords, crop)
		area += areaEx
		perimeter += perimeterEX
	}
	return area, perimeter
}

func cornerCheck(u int, d int, l int, r int, ul int, dl int, ur int, dr int) int {
	//returns number of corners based on which surrounding blocks are clear
	cornerCount := 0
	//reflex corners
	if u == 1 && l == 1 {
		cornerCount += 1
	}
	if d == 1 && l == 1 {
		cornerCount += 1
	}
	if u == 1 && r == 1 {
		cornerCount += 1
	}
	if d == 1 && r == 1 {
		cornerCount += 1
	}
	//accute corners
	if d == 0 && r == 0 && dr == 1 {
		cornerCount += 1
	}
	if u == 0 && r == 0 && ur == 1 {
		cornerCount += 1
	}
	if d == 0 && l == 0 && dl == 1 {
		cornerCount += 1
	}
	if u == 0 && l == 0 && ul == 1 {
		cornerCount += 1
	}
	return cornerCount
}

func regionMapperBulk(mp [][]string, coords [][]int, crop string) (int, int) {
	area := 0
	corner := 0
	var nextCoords [][]int
	for _, i := range coords {
		if mp[i[0]][i[1]] != ("#" + crop) { //don't revist node
			area += 1
			mp[i[0]][i[1]] = ("#" + crop)
			u := 0
			d := 0
			l := 0
			r := 0
			dr := 0
			dl := 0
			ur := 0
			ul := 0
			//check up
			if (i[0] - 1) >= 0 { //in bounds
				if mp[i[0]-1][i[1]] == crop {
					nextCoords = append(nextCoords, []int{i[0] - 1, i[1]})
				} else if mp[i[0]-1][i[1]] != ("#" + crop) {
					u = 1
				}
			} else {
				u = 1
			}
			//check down
			if (i[0] + 1) < len(mp) { //in bounds
				if mp[i[0]+1][i[1]] == crop {
					nextCoords = append(nextCoords, []int{i[0] + 1, i[1]})
				} else if mp[i[0]+1][i[1]] != ("#" + crop) {
					d = 1
				}
			} else {
				d = 1
			}
			//check left
			if (i[1] - 1) >= 0 { //in bounds
				if mp[i[0]][i[1]-1] == crop {
					nextCoords = append(nextCoords, []int{i[0], i[1] - 1})
				} else if mp[i[0]][i[1]-1] != ("#" + crop) {
					l = 1
				}
			} else {
				l = 1
			}
			//check right
			if (i[1] + 1) < len(mp[i[0]]) { //in bounds
				if mp[i[0]][i[1]+1] == crop {
					nextCoords = append(nextCoords, []int{i[0], i[1] + 1})
				} else if mp[i[0]][i[1]+1] != ("#" + crop) {
					r = 1
				}
			} else {
				r = 1
			}
			//diagonals (just for corners)
			if (i[1]+1) < len(mp[i[0]]) && (i[0]+1) < len(mp) {
				if mp[i[0]+1][i[1]+1] != crop && mp[i[0]+1][i[1]+1] != ("#"+crop) {
					dr = 1
				}
			} else {
				dr = 1
			}
			if (i[1]+1) < len(mp[i[0]]) && (i[0]-1) >= 0 {
				if mp[i[0]-1][i[1]+1] != crop && mp[i[0]-1][i[1]+1] != ("#"+crop) {
					ur = 1
				}
			} else {
				ur = 1
			}
			if (i[1]-1) >= 0 && (i[0]-1) >= 0 {
				if mp[i[0]-1][i[1]-1] != crop && mp[i[0]-1][i[1]-1] != ("#"+crop) {
					ul = 1
				}
			} else {
				ul = 1
			}
			if (i[1]-1) >= 0 && (i[0]+1) < len(mp) {
				if mp[i[0]+1][i[1]-1] != crop && mp[i[0]+1][i[1]-1] != ("#"+crop) {
					dl = 1
				}
			} else {
				dl = 1
			}
			corner += cornerCheck(u, d, l, r, ul, dl, ur, dr)
		}
	}
	if len(nextCoords) != 0 {
		areaEx, cornerEX := regionMapperBulk(mp, nextCoords, crop)
		area += areaEx
		corner += cornerEX
	}
	return area, corner
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day12/Input.txt")
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//Insert additional processsing here
	var garden [][]string
	for _, i := range raw {
		garden = append(garden, strings.Split(i, ""))
	}
	//insert Puzzle solution here
	sum := 0
	for row, i := range garden {
		for col, j := range i {
			if !strings.Contains(j, "#") { //checkTrails
				var Startcords [][]int
				Startcords = append(Startcords, []int{row, col})
				a, p := regionMapper(garden, Startcords, j)
				sum += (a * p)

			}
		}
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day12/Input.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	//Insert additional processsing here
	var garden [][]string
	for _, i := range raw {
		garden = append(garden, strings.Split(i, ""))
	}
	//inerst Puzzle solution here
	sum := 0
	for row, i := range garden {
		for col, j := range i {
			if !strings.Contains(j, "#") { //checkTrails
				var Startcords [][]int
				Startcords = append(Startcords, []int{row, col})
				a, p := regionMapperBulk(garden, Startcords, j)
				sum += (a * p)
			}
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sum
}
