package main

import (
	"fmt"
	"strings"
)

func main() {
	raw := fileParse.fileParse("day6/TestInput.txt")
	var row int
	var col int
	input := [][]string
	for _, i := range raw {
		input = append(input, strings.Split(i, ""))
	}

	//get startpoint
	for i, lin := range input {
		if strings.Contains(lin, "^") {
			row = i
			col = strings.Index(lin, ">")
		}
	}
	nextrow := row
	nextcol := col
	block := true
	inMap := true
	currentDirInd := 0                                         //always starts moving up
	directionMods := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //direction modifiers Up -> Right -> Down -> Left

	//pathing loop
	for inMap {
		currentDir = directionMods[currentDirInd]
		for block {
			nextrow = row + currentDir[0]
			nextcol = col + currentDir[1]
			if nextrow == -1 || nextrow == (len(input)-1) || nextcol == -1 || nextrow == (len(input[0])-1) { // if boundary (goal) is reached
				inMap = false

				if currentDirInd == 3 {
					currentDirInd = 0
				} else {
					currentDirInd += 1
				}
				block = false
			} else if input[row][col] == "#" {
				//end of line, turn right
				if currentDirInd == 3 {
					currentDirInd = 0
				} else {
					currentDirInd += 1
				}
				block = false
			} else { //continue as normal, mark block, update coords to next coords
				input[row][col] = "X"
				row = nextrow
				col = newcol
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

	fmt.Println(steps)

}
