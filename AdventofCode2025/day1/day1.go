package day1

import "time"

func Part1() (time.Duration, time.Duration, string) {
	var p1 string
	ParseStart := time.Now()
	//parse function
	ParseTime := time.Since(ParseStart)
	startP1 := time.Now()
	//puzzle
	p1 = "Hello"

	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, string) {
	var p2 string
	ParseStart := time.Now()
	//parse function
	ParseTime := time.Since(ParseStart)
	startP2 := time.Now()
	//puzzle
	p2 = "World"

	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
