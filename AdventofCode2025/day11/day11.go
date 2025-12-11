package day11

import (
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func Trace(node string, wire map[string]string) int {
	path := 0
	if wire[node] == "out" {
		return 1
	}
	child := strings.Split(wire[node], " ")
	for _, i := range child {
		path += Trace(i, wire)
	}
	return path
}

func Bts(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func TraceDacFft(node string, wire map[string]string, dac bool, fft bool, cache map[string]int) int {
	n := node + "," + Bts(dac) + "," + Bts(fft)
	if _, ok := cache[n]; !ok {
		path := 0
		if wire[node] == "out" {
			if dac && fft {
				return 1
			} else {
				return 0
			}
		}
		child := strings.Split(wire[node], " ")
		for _, i := range child {
			path += TraceDacFft(i, wire, (node == "dac") || dac, (node == "fft") || fft, cache)
		}
		cache[n] = path
	}
	return cache[n]
}

func Part1() (time.Duration, time.Duration, int) {
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day11/Input.txt")
	mp := make(map[string]string)
	ParseTime := time.Since(ParseStart)
	//puzzle
	p1 := 0
	startP1 := time.Now()
	for _, i := range Input {
		row := strings.Split(i, ": ")
		mp[row[0]] = row[1]
	}
	cache := make(map[string]int)
	p1 = TraceDacFft("you", mp, true, true, cache)
	//p1 = Trace("you", mp)
	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day11/Input.txt")
	mp := make(map[string]string)
	ParseTime := time.Since(ParseStart)
	//puzzle
	p2 := 0
	startP2 := time.Now()
	for _, i := range Input {
		row := strings.Split(i, ": ")
		mp[row[0]] = row[1]
	}
	cache := make(map[string]int)

	p2 = TraceDacFft("svr", mp, false, false, cache)
	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
