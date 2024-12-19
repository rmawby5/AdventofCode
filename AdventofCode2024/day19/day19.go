package day19

import (
	"slices"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func isValidArrangement(arrangement string, towels []string, bufLen int) bool {
	//checks if an arrangement is valid
	valid := false
	buffer := ""
	char := 0
	for len(buffer) < bufLen {
		buffer = buffer + string(arrangement[char])
		if slices.Contains(towels, buffer) {
			if len(buffer) == len(arrangement) {
				return true
			}
			valid = isValidArrangement(arrangement[len(buffer):], towels, bufLen)
			if valid {
				return valid
			}
		} else if len(buffer) == len(arrangement) {
			break
		}
		char++
	}
	return valid
}

func numValidArrangement(arrangement string, towels []string, bufLen int, cache map[string]int) int {
	//checks the total number of valid arrangements for a given order
	buffer := ""
	char := 0
	if cache[arrangement] == 0 {
		numVar := 0
		for len(buffer) < bufLen {
			buffer = buffer + string(arrangement[char])
			if slices.Contains(towels, buffer) {
				if len(buffer) == len(arrangement) {
					numVar += 1
					break
				}
				numVar += numValidArrangement(arrangement[len(buffer):], towels, bufLen, cache)
			} else if len(buffer) == len(arrangement) {
				break
			}
			char++
		}
		cache[arrangement] = numVar
	}
	return cache[arrangement]
}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	input := fileparse.FileParse("day19/Input.txt")
	towelRaw := fileparse.FileParse("day19/Towels.txt")
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//Insert additional processsing here
	towels := strings.Split(towelRaw[0], ", ")
	maxLen := 0 //maximum number of stripes on a given towel
	for _, i := range towels {
		if len(i) > maxLen {
			maxLen = len(i)
		}
	}
	//insert Puzzle solution here
	sum := 0
	for _, i := range input {
		if isValidArrangement(i, towels, maxLen) {
			sum++
		}
	}
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}

func Part2() (time.Duration, time.Duration, int) {
	cache := make(map[string]int)
	ParseStart := time.Now()
	input := fileparse.FileParse("day19/Input.txt")
	towelRaw := fileparse.FileParse("day19/Towels.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	//Insert additional processsing here
	towels := strings.Split(towelRaw[0], ", ")
	maxLen := 0 //maximum number of stripes on a given towel
	for _, i := range towels {
		if len(i) > maxLen {
			maxLen = len(i)
		}
	}
	//insert Puzzle solution here
	sum := 0
	for _, i := range input {
		sum += numValidArrangement(i, towels, maxLen, cache)
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, sum
}
