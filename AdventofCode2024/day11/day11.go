package day11

import (
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func splitter(full int) (int, int) {
	//takes int value, returns two ints with left and right half of input's digits
	char := strconv.Itoa(full)
	left := char[:len(char)/2]
	right := char[len(char)/2:]
	leftI, _ := strconv.Atoi(left)
	rightI, _ := strconv.Atoi(right)
	return leftI, rightI
}

func lenInt(i int) bool {
	if i >= 1e18 {
		return false
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count%2 == 0

}

func blinker(input int, blinkCount int, blinkLimit int, cache map[string]int) int {
	//stones := 0
	//case 0
	n := strconv.Itoa(input) + "," + strconv.Itoa(blinkCount)
	if cache[n] == 0 {
		if blinkCount == blinkLimit {
			cache[n] = 1
		} else {
			if input == 0 {
				cache[n] = blinker(1, blinkCount+1, blinkLimit, cache)
			} else if lenInt(input) { //check function
				l, r := splitter(input)
				cache[n] = blinker(l, blinkCount+1, blinkLimit, cache) + blinker(r, blinkCount+1, blinkLimit, cache)
				//stones += blinker(r, blinkCount+1,cache)
			} else {
				cache[n] = blinker(input*2024, blinkCount+1, blinkLimit, cache)
			}
		}
	}

	return cache[n]
}

func Part1() (time.Duration, time.Duration, int) {
	cache := make(map[string]int)
	ParseStart := time.Now()
	raw := fileparse.FileParse("day11/Input.txt")
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//Insert additional processsing here
	var in []int
	for _, i := range raw {
		tmp := strings.Split(i, " ")
		for _, j := range tmp {
			p, _ := strconv.Atoi(j)
			in = append(in, p)
		}
	}
	//inerst Puzzle solution here
	stoneCount := 0
	//fmt.Println(in)
	for _, k := range in {
		stoneCount += blinker(k, 0, 25, cache)

	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, stoneCount
}

func Part2() (time.Duration, time.Duration, int) {
	cache := make(map[string]int)
	ParseStart := time.Now()
	raw := fileparse.FileParse("day11/Input.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	//Insert additional processsing here
	var in []int
	for _, i := range raw {
		tmp := strings.Split(i, " ")
		for _, j := range tmp {
			p, _ := strconv.Atoi(j)
			in = append(in, p)
		}
	}
	//inerst Puzzle solution here
	stoneCount := 0
	//fmt.Println(in)
	for _, k := range in {
		stoneCount += blinker(k, 0, 75, cache)

	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, stoneCount
}
