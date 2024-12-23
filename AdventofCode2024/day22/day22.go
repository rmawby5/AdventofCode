package day22

import (
	"fmt"
	"strconv"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func secret(seed int) int {
	start := seed
	for i := 0; i < 2000; i += 1 {
		start = (((start * 64) ^ start) % 16777216)
		start = (((start / 32) ^ start) % 16777216)
		start = (((start * 2048) ^ start) % 16777216)
	}
	return start
}

func lastDigit(in int) int {
	st := strconv.Itoa(in)
	last := st[len(st)-1:]
	fmt.Println(last)
	lI, _ := strconv.Atoi(last)
	return lI
}

func price(seed int) int {
	start := seed
	var change []int
	for i := 0; i < 2000; i += 1 {
		start = (((start * 64) ^ start) % 16777216)
		start = (((start / 32) ^ start) % 16777216)
		start = (((start * 2048) ^ start) % 16777216)
		change = append(change, lastDigit(start))
	}

	return start
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day22/TestInput.txt")
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	//Insert additional processsing here
	var input []int
	for _, i := range raw {
		j, _ := strconv.Atoi(i)
		input = append(input, j)
	}

	//insert Puzzle solution here
	sum := 0
	for _, i := range input {
		sum += secret(i)
	}

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day22/TestInput.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()

	//Insert additional processsing here
	var input []int
	for _, i := range raw {
		j, _ := strconv.Atoi(i)
		input = append(input, j)
	}

	//insert Puzzle solution here
	sum := 0
	for _, i := range input {
		sum += secret(i)
	}

	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sum
}
