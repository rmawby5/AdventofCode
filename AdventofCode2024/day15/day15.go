package day15

import (
	"slices"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func dirKey(dir string) []int {
	if dir == "<" {
		return []int{0, -1}
	} else if dir == ">" {
		return []int{0, 1}
	} else if dir == "v" {
		return []int{1, 0}
	} else if dir == "^" {
		return []int{-1, 0}
	}
	return []int{0, 0}
}

func mover(mp [][]string, start []int, seq []string) [][]string {
	modMp := make([][]string, len(mp))
	_ = copy(modMp, mp)

	curPos := make([]int, 2)
	_ = copy(curPos, start)

	//fmt.Println("Start : ", curPos)
	for _, i := range seq {
		dir := dirKey(i)
		//fmt.Println(dir)
		if modMp[curPos[0]+dir[0]][curPos[1]+dir[1]] == "." {
			modMp[curPos[0]][curPos[1]] = "."
			curPos[0] = curPos[0] + dir[0]
			curPos[1] = curPos[1] + dir[1]
			modMp[curPos[0]][curPos[1]] = "@"
		} else if modMp[curPos[0]+dir[0]][curPos[1]+dir[1]] == "O" {
			//check if boxes can be pushed
			push := true
			nextPos := []int{0, 0}
			nextPos[0] = curPos[0] + dir[0]
			nextPos[1] = curPos[1] + dir[1]
			for push {
				if modMp[nextPos[0]+dir[0]][nextPos[1]+dir[1]] == "O" { //continue checking along direction
					nextPos[0] = nextPos[0] + dir[0]
					nextPos[1] = nextPos[1] + dir[1]
				} else if modMp[nextPos[0]+dir[0]][nextPos[1]+dir[1]] == "#" { //blocks cannot be pushed, stop and continue with next move
					push = false
				} else if mp[nextPos[0]+dir[0]][nextPos[1]+dir[1]] == "." { //space found to push blocks
					modMp[nextPos[0]+dir[0]][nextPos[1]+dir[1]] = "O"
					modMp[curPos[0]+dir[0]][curPos[1]+dir[1]] = "@"
					modMp[curPos[0]][curPos[1]] = "."
					curPos[0] = curPos[0] + dir[0]
					curPos[1] = curPos[1] + dir[1]
					push = false
				}
			}
		}
	}
	return modMp
}

func gps(mp [][]string) int {
	tot := 0

	for i, row := range mp {
		for j, point := range row {
			if point == "O" {
				tot += ((i * 100) + j)
			}
		}
	}

	return tot
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	pathS := fileparse.FileParse("day15/InputPath.txt")
	mpRaw := fileparse.FileParse("day15/Input.txt")

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	var pathf string
	var mp [][]string

	var start []int
	//Insert additional processsing here
	for _, i := range pathS {
		pathf = pathf + i
	}
	path := strings.Split(pathf, "")

	for k, j := range mpRaw {
		if strings.Contains(j, "@") {
			start = append(start, k)
			tmp := strings.Split(j, "")
			start = append(start, slices.Index(tmp, "@"))
			mp = append(mp, tmp)
		} else {
			mp = append(mp, strings.Split(j, ""))
		}
	}
	//insert Puzzle solution here
	modMap := mover(mp, start, path)
	sum := gps(modMap)

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	pathS := fileparse.FileParse("day15/InputPath.txt")
	mpRaw := fileparse.FileParse("day15/Input.txt")

	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	var pathf string
	var mp [][]string

	var start []int
	//Insert additional processsing here
	for _, i := range pathS {
		pathf = pathf + i
	}
	path := strings.Split(pathf, "")

	for k, j := range mpRaw {
		if strings.Contains(j, "@") {
			start = append(start, k)
			tmp := strings.Split(j, "")
			start = append(start, slices.Index(tmp, "@"))
			mp = append(mp, tmp)
		} else {
			mp = append(mp, strings.Split(j, ""))
		}
	}
	//insert Puzzle solution here
	modMap := mover(mp, start, path)
	sum := gps(modMap)

	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sum
}
