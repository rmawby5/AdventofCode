package day11

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)




func splitter(full int)(int, int){
	//takes int value, returns two ints with left and right half of input's digits
	char,err = strconv.Itoa(full)
	left := char[:len(char)/2]
	right := char[len(char)/2:]
	leftI,_ := strconv.Atoi(left)	
	rightI,_ := strconv.Atoi(right)
	return leftI,rightI
}

func lenInt(i int) bool {
    if i >= 1e18 {
        return 19
    }
    x, count := 10, 1
    for x <= i {
        x *= 10
        count++
    }
	if count%2 == 0{
		return true
	}
    return false
}
func blinker(input int, blinkCount int)int{
	stones := 0
	//case 0
	if blinkCount == 25{
		stones = 1
	}else{
		if input == 0{
			stones += blinker(1,blinkCount+1)
		}else if lenInt(input){										//check function
			l,r := splitter(input)
			stones+= blinker(l, blinkCount+1)
			stones+=blinker(r,blinkCount+1)
		}else{
			stones+= blinker(input*2024, blinkCount+1)
		}
	}
	return stones
}

func Part1() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day10/TestInput.txt")
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//Insert additional processsing here
	var in []int
	for _,i:= range raw{
		in = strings.Split(i,"")
	}
	//inerst Puzzle solution here
	stoneCount := 0
	for _,k := range in {
		conv,_ := strconv.Atoi(k)
		stoneCount+= blinker(conv, )
		
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, stoneCount
}
