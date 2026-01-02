package day10

import (
	"math"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2025/FileParse"
)

func CombCheck(mac machine, butSel []string) bool {
	status := make([]string, len(mac.indicator))
	for i := range status {
		status[i] = "."
	}
	for i, v := range butSel {
		j, _ := strconv.Atoi(v)
		if j == 1 {
			but := strings.Split(strings.Trim(mac.buttons[i], "()"), ",")
			for _, k := range but {
				idx, _ := strconv.Atoi(k)
				if status[idx] == "." { //if indicator is off, set to on
					status[idx] = "#"
				} else {
					status[idx] = "." //otherwise set to off
				}
			}

		}
	}
	for m := range status {
		if mac.indicator[m] != status[m] {
			return false
		}
	}
	return true
}

func pressCount(comb []string) int {
	tot := 0
	for _, i := range comb {
		if i == "1" {
			tot++
		}
	}
	return tot
}

type machine struct {
	indicator []string
	buttons   []string
	joltage   []string
}

func Part1() (time.Duration, time.Duration, int) {
	p1 := 0
	//parse function
	ParseStart := time.Now()
	Input := fileparse.FileParse("day10/Input.txt")

	ParseTime := time.Since(ParseStart)
	//puzzle
	startP1 := time.Now()
	for _, i := range Input {
		sep := strings.Split(i, " ")
		curMc := machine{strings.Split(strings.Trim(sep[0], "[]"), ""), sep[1 : len(sep)-1], strings.Split(strings.Trim(sep[len(sep)-1], "{}"), ",")}
		length := len(curMc.buttons)
		curMin := length
		lim := int(math.Pow(2, float64(length)))
		for i := 0; i < lim; i++ { //generate combinations of 0/1s upto the number of buttons using binary counting
			var c []string
			d := strings.Split(strconv.FormatInt(int64(i), 2), "") //generate binary represenation of the current number for each combination
			for j := 0; j < length-len(d); j++ {                   //add leading 0s
				c = append(c, "0")
			}
			c = append(c, d...)              //add binary representation
			if CombCheck(curMc, c) == true { //check if the combination is valid
				t := pressCount(c)
				if t < curMin { // if valid, check if the combination has fewer presses than the current best
					curMin = t
				}
			}
		}
		p1 += curMin //add the best combination to total
	}

	P1Time := time.Since(startP1)
	return ParseTime, P1Time, p1
}

func Part2() (time.Duration, time.Duration, int) {
	p2 := 0
	//parse function
	ParseStart := time.Now()
	//Input := fileparse.FileParse("day1/Input.txt")
	ParseTime := time.Since(ParseStart)
	//puzzle
	startP2 := time.Now()
	P2Time := time.Since(startP2)
	return ParseTime, P2Time, p2
}
