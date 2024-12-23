package day20

import (
	"slices"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func trackCheck(mp [][]string, start []int) [][]string {
	stepCount := 0
	modMP := make([][]string, len(mp))
	_ = copy(modMP, mp)
	row := start[0]
	col := start[1]
	race := true
	modMP[row][col] = strconv.Itoa(stepCount)
	for race {
		//fmt.Println("Row:", row, " Col:", col)
		//check up
		if (row - 1) >= 0 { //in bounds
			if modMP[row-1][col] == "." {
				row = row - 1
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				continue
			} else if modMP[row-1][col] == "E" {
				race = false
				row = row - 1
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
			}
		}
		//check down
		if (row + 1) < len(modMP) { //in bounds
			if modMP[row+1][col] == "." {
				row++
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				continue
			} else if modMP[row+1][col] == "E" {
				race = false
				row++
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
			}
		}
		//check left
		if (col - 1) >= 0 { //in bounds
			if modMP[row][col-1] == "." {
				col = col - 1
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				continue
			} else if modMP[row][col-1] == "E" {
				race = false
				col = col - 1
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
			}
		}
		//check right
		if (col + 1) < len(modMP[0]) { //in bounds
			if mp[row][col+1] == "." {
				col++
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				continue
			} else if modMP[row][col+1] == "E" {
				race = false
				col++
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
			}
		}
	}
	return modMP
}

func cheat(mp [][]string, row int, col int) int {
	t := 101 //cheat threashold
	if mp[row-1][col] != "#" && mp[row+1][col] != "#" {
		m, _ := strconv.Atoi(mp[row-1][col])
		n, _ := strconv.Atoi(mp[row+1][col])
		if n-m > t || m-n > t {
			return 1
		}
	}
	if mp[row][col-1] != "#" && mp[row][col+1] != "#" {
		m, _ := strconv.Atoi(mp[row][col-1])
		n, _ := strconv.Atoi(mp[row][col+1])
		if n-m > t || m-n > t {
			return 1
		}
	}
	return 0
}

func trackCheckList(mp [][]string, start []int) [][]int {
	//same as track check, but return the path as a sorted list of coordinates
	stepCount := 0
	modMP := make([][]string, len(mp))
	_ = copy(modMP, mp)
	row := start[0]
	col := start[1]
	race := true
	modMP[row][col] = strconv.Itoa(stepCount)
	var defaultTrack [][]int
	defaultTrack = append(defaultTrack, start)
	for race {
		//check up
		if (row - 1) >= 0 { //in bounds
			if modMP[row-1][col] == "." {
				row = row - 1
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				point := []int{row, col}
				defaultTrack = append(defaultTrack, point)
				continue
			} else if modMP[row-1][col] == "E" {
				race = false
				row = row - 1
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				point := []int{row, col}
				defaultTrack = append(defaultTrack, point)
			}
		}
		//check down
		if (row + 1) < len(modMP) { //in bounds
			if modMP[row+1][col] == "." {
				row++
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				point := []int{row, col}
				defaultTrack = append(defaultTrack, point)
				continue
			} else if modMP[row+1][col] == "E" {
				race = false
				row++
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				point := []int{row, col}
				defaultTrack = append(defaultTrack, point)
			}
		}
		//check left
		if (col - 1) >= 0 { //in bounds
			if modMP[row][col-1] == "." {
				col = col - 1
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				point := []int{row, col}
				defaultTrack = append(defaultTrack, point)
				continue
			} else if modMP[row][col-1] == "E" {
				race = false
				col = col - 1
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				point := []int{row, col}
				defaultTrack = append(defaultTrack, point)
			}
		}
		//check right
		if (col + 1) < len(modMP[0]) { //in bounds
			if mp[row][col+1] == "." {
				col++
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				point := []int{row, col}
				defaultTrack = append(defaultTrack, point)
				continue
			} else if modMP[row][col+1] == "E" {
				race = false
				col++
				stepCount++
				modMP[row][col] = strconv.Itoa(stepCount)
				point := []int{row, col}
				defaultTrack = append(defaultTrack, point)
			}
		}
	}
	return defaultTrack
}

func absI(i int) int {
	if i < 0 {
		return (0 - i)
	}
	return i
}

func manhattan(p1 []int, p2 []int) int {
	y := absI(p1[0] - p2[0])
	x := absI(p1[1] - p2[1])
	return x + y
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day20/Input.txt")

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	//Insert additional processsing here
	var input [][]string
	start := []int{0, 0}
	for j, i := range raw {
		tmp := strings.Split(i, "")
		if slices.Contains(tmp, "S") {
			start[0] = j
			start[1] = slices.Index(tmp, "S")
		}
		input = append(input, tmp)
	}

	//insert Puzzle solution here
	sum := 0
	mod := trackCheck(input, start)
	for i := 1; i < len(mod)-1; i++ {
		for k := 1; k < len(mod[i])-1; k++ {
			if mod[i][k] == "#" {
				sum += cheat(mod, i, k)
			}

		}
	}

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day20/Input.txt")

	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()

	//Insert additional processsing here
	var input [][]string
	start := []int{0, 0}
	for j, i := range raw {
		tmp := strings.Split(i, "")
		if slices.Contains(tmp, "S") {
			start[0] = j
			start[1] = slices.Index(tmp, "S")
		}
		input = append(input, tmp)
	}
	//insert Puzzle solution here
	sum := 0
	mod := trackCheckList(input, start)
	cheatThreashold := 100
	cheatDuration := 20
	for i := 0; i < len(mod)-cheatThreashold; i++ {
		for j := i + cheatThreashold; j < len(mod); j++ {
			cheatLen := manhattan(mod[i], mod[j])
			if cheatLen <= cheatDuration && (j-i) >= (cheatLen+cheatThreashold) {
				sum++
			}
		}
	}

	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sum
}
