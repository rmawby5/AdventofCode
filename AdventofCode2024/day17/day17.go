package day17

import (
	"math"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func dv(reg int, op int) int {
	return reg / powInt(2, op)
}

func solver(regA int, regB int, regC int, prog []string) string {
	ops := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": regA,
		"5": regB,
		"6": regC,
	}
	output := ""
	i := 0
	for i < len(prog) {
		if prog[i] == "0" {
			ops["4"] = dv(ops["4"], ops[prog[i+1]])
		} else if prog[i] == "1" {
			op, _ := strconv.Atoi(prog[i+1])
			ops["5"] = (ops["5"] ^ op)
		} else if prog[i] == "2" {
			ops["5"] = (ops[prog[i+1]] % 8)
		} else if prog[i] == "4" {
			ops["5"] = (ops["5"] ^ ops["6"])
		} else if prog[i] == "5" {
			output = output + strconv.Itoa(ops[prog[i+1]]%8) + ","
		} else if prog[i] == "6" {
			ops["5"] = dv(ops["4"], ops[prog[i+1]])
		} else if prog[i] == "7" {
			ops["6"] = dv(ops["4"], ops[prog[i+1]])
		}
		if prog[i] == "3" {
			if ops["4"] != 0 {
				i, _ = strconv.Atoi(prog[i+1])
			} else {
				i += 2
			}
		} else {
			i += 2
		}
	}
	return output
}

func match(prog []string, out string) (bool, bool) {
	outS := strings.Split(out[:len(out)-1], ",")
	match := true
	lenMatch := false

	for i := range outS {
		if prog[len(prog)-1-i] != outS[len(outS)-1-i] {
			match = false
			break
		}
	}
	if len(outS) == len(prog) {
		lenMatch = true
	}
	return match, lenMatch
}

func Part1() (time.Duration, time.Duration, string) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day17/Input.txt")

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	//Insert additional processsing here
	var input [][]string
	for _, i := range raw {
		input = append(input, strings.Split(i, ": "))
	}
	regA, _ := strconv.Atoi(input[0][1])
	regB, _ := strconv.Atoi(input[1][1])
	regC, _ := strconv.Atoi(input[2][1])
	prog := strings.Split(input[3][1], ",")
	//insert Puzzle solution here
	output := solver(regA, regB, regC, prog)
	output = output[:len(output)-1]

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, output
}

func Part2() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day17/Input.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()

	//Insert additional processsing here
	var input [][]string
	for _, i := range raw {
		input = append(input, strings.Split(i, ": "))
	}
	regB, _ := strconv.Atoi(input[1][1])
	regC, _ := strconv.Atoi(input[2][1])
	prog := strings.Split(input[3][1], ",")

	//insert Puzzle solution here
	regA := 1
	output := ""
	progMatch := false

	for !progMatch {
		output = solver(regA, regB, regC, prog)
		curMatch, progM := match(prog, output)
		if curMatch {
			if progM {
				progMatch = true
			} else {
				regA = (regA * 8)
			}
		} else {
			regA++
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, regA
}
