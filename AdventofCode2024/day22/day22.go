package day22

import (
	"slices"
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

func stringSquash(buffer []int) string {
	str := ""
	for _, i := range buffer {
		str = str + strconv.Itoa(i) + ","
	}
	return str
}
func price(seed int, cache map[string]int) ([]int, map[string]int) {
	start := seed
	var change []int
	seq := ""
	var visited []string

	for i := 0; i < 2000; i += 1 {
		b := start % 10
		start = (((start * 64) ^ start) % 16777216)
		start = (((start / 32) ^ start) % 16777216)
		start = (((start * 2048) ^ start) % 16777216)
		change = append(change, start%10-b)
		if len(change) == 5 {
			change = change[1:]
			seq = stringSquash(change)
			if cache[seq] == 0 {
				cache[seq] = start % 10
				visited = append(visited, seq)
			} else if !slices.Contains(visited, seq) {
				cache[seq] = cache[seq] + start%10
				visited = append(visited, seq)
			}
		}
	}
	return change, cache
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day22/Input.txt")
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
	raw := fileparse.FileParse("day22/Input.txt")
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
	cache := make(map[string]int)
	for _, i := range input {
		_, cache = price(i, cache)
	}

	for _, ban := range cache {
		if ban > sum {
			sum = ban
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sum
}
