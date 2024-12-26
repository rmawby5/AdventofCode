package day25

import (
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func comp(l []int, k []int) bool {
	match := true
	for i := range l {
		if l[i]+k[i] > 7 {
			match = false
			break
		}
	}
	return match
}
func conv(block []string) ([]int, string) {
	heights := make([]int, len(block[0]))
	var typ string
	if block[0] == "#####" {
		typ = "Lock"
	} else {
		typ = "Key"
	}

	if typ == "Lock" {
		for i := 0; i < len(block[0]); i++ {
			for j := 0; j < len(block); j++ {
				if string(block[j][i]) != "#" {
					heights[i] = j
					break
				}
			}
		}
	} else if typ == "Key" {
		for i := 0; i < len(block[0]); i++ {
			for j := 0; j < len(block); j++ {
				if string(block[j][i]) == "#" {
					heights[i] = (len(block) - j)
					break
				}
			}
		}
	}

	return heights, typ
}
func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day25/Input.txt")

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	//Insert additional processsing here
	var locks [][]int
	var keys [][]int
	var buffer []string
	for r, i := range raw {
		if len(i) != 0 {
			buffer = append(buffer, i)
		}
		if len(i) == 0 || r == len(raw)-1 {
			switch h, t := conv(buffer); t {
			case "Key":
				keys = append(keys, h)
			case "Lock":
				locks = append(locks, h)
			}
			buffer = nil
		}
	}
	//insert Puzzle solution here
	sum := 0
	for _, l := range locks {
		for _, k := range keys {
			if comp(l, k) {
				sum++
			}
		}
	}

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}
