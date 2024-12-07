package day6

import (
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day6/Input.txt")
	var row int
	var col int
	var input [][]string
	for _, i := range raw {
		input = append(input, strings.Split(i, ""))
	}

	//get startpoint
	for i, lin := range raw {
		if strings.Contains(lin, "^") {
			row = i
			col = strings.Index(lin, "^")
		}
	}
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	nextrow := row
	nextcol := col
	block := true
	inMap := true
	currentDirInd := 0                                         //always starts moving up
	directionMods := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //direction modifiers Up -> Right -> Down -> Left

	//pathing loop
	for inMap {
		currentDir := directionMods[currentDirInd]

		block = true
		for block {
			nextrow = row + currentDir[0]
			nextcol = col + currentDir[1]

			if nextrow == -1 || nextrow == (len(input)) || nextcol == -1 || nextcol == (len(input[0])) { // if boundary (goal) is reached
				input[row][col] = "X"
				inMap = false

				break
			} else if input[nextrow][nextcol] == "#" {

				if currentDirInd == 3 {
					currentDirInd = 0
				} else {
					currentDirInd += 1
				}

				block = false
			} else { //continue as normal, mark block, update coords to next coords
				input[row][col] = "X"
				row = nextrow
				col = nextcol
			}
		}

	}

	steps := 0
	for _, i := range input {
		for _, j := range i {
			if j == "X" {
				steps += 1
			}
		}
	}
	P1Time := time.Since(P1Start)

	return ParseTime, P1Time, steps

}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day6/Input.txt")
	var row int
	var col int
	var startRow int
	var startCol int
	var input [][]string
	for _, i := range raw {
		input = append(input, strings.Split(i, ""))
	}

	//get startpoint
	for i, lin := range raw {
		if strings.Contains(lin, "^") {
			startRow = i
			startCol = strings.Index(lin, "^")
		}
	}

	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()

	var blockLocations [][]int

	row = startRow
	col = startCol
	nextrow := row
	nextcol := col
	block := true
	inMap := true
	currentDirInd := 0                                         //always starts moving up
	directionMods := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //direction modifiers Up -> Right -> Down -> Left

	//pathing loop
	for inMap {
		currentDir := directionMods[currentDirInd]

		block = true
		for block {
			nextrow = row + currentDir[0]
			nextcol = col + currentDir[1]

			if nextrow == -1 || nextrow == (len(input)) || nextcol == -1 || nextcol == (len(input[0])) { // if boundary (goal) is reached
				if input[row][col] != "X" {
					input[row][col] = "X"
					blockLocations = append(blockLocations, []int{row, col})
				}
				inMap = false

				break
			} else if input[nextrow][nextcol] == "#" {
				//hit block, turn right
				//fmt.Println("Wibble")
				if currentDirInd == 3 {
					currentDirInd = 0
				} else {
					currentDirInd += 1
				}

				block = false
			} else { //continue as normal, mark block, update coords to next coords
				if input[row][col] != "X" {
					input[row][col] = "X"
					blockLocations = append(blockLocations, []int{row, col})
				}
				row = nextrow
				col = nextcol
			}
		}

	}

	/// for brute force = for each coord in slice "block location", try the loop above with addition of loop detection (look two blocks ahead rather than 1.)
	loopCount := 0
	for _, c := range blockLocations {
		block := true
		inMap := true
		currentDirInd = 0
		var inputMod [][]string
		for _, i := range raw {
			inputMod = append(inputMod, strings.Split(i, ""))
		}
		//reset start location
		row = startRow
		col = startCol
		//place extra block
		inputMod[c[0]][c[1]] = "#"

		for inMap {
			currentDir := directionMods[currentDirInd]

			block = true
			for block {
				nextrow = row + currentDir[0]
				nextcol = col + currentDir[1]

				if nextrow == -1 || nextrow == (len(inputMod)) || nextcol == -1 || nextcol == (len(inputMod[0])) { // if boundary (goal) is reached
					inMap = false

					break
				} else if inputMod[nextrow][nextcol] == "#" {

					if inputMod[row][col] == "*" {

						loopCount += 1
						inMap = false
						break
					}

					if currentDirInd == 3 {
						currentDirInd = 0
					} else {
						currentDirInd += 1
					}

					next2row := row + directionMods[currentDirInd][0]
					next2col := col + directionMods[currentDirInd][1]
					if inputMod[next2row][next2col] != "#" {
						inputMod[row][col] = "*"
					}

					block = false
				} else { //continue as normal, mark current block, and move to the next block.
					row = nextrow
					col = nextcol
				}
			}

		}

	}

	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, loopCount

}
