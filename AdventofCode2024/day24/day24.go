package day24

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func chip(vals map[string]int, gates [][]string) int {

	var procesedGates []int
	sum := 0
	for len(procesedGates) != len(gates) {

		for i, v := range gates {
			if slices.Contains(procesedGates, i) {
				continue //skip line if already processed gate
			}
			w1, ok1 := vals[v[0]]
			w2, ok2 := vals[v[2]]
			if !ok1 || !ok2 {
				continue //skip line if either input has not been procesed yet
			}
			switch op := v[1]; op {
			case "AND":
				vals[v[4]] = (w1 & w2)
			case "OR":
				vals[v[4]] = (w1 | w2)
			case "XOR":
				vals[v[4]] = (w1 ^ w2)
			}

			if v[4][:1] == "z" {
				k, _ := strconv.Atoi(v[4][1:])
				//fmt.Println(k)
				sum += (int(math.Pow(2, float64(k))) * vals[v[4]])
			}

			procesedGates = append(procesedGates, i)
		}
	}

	return sum

}

func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day24/Input.txt")
	val := fileparse.FileParse("day24/InputValues.txt")

	values := make(map[string]int)
	var input [][]string

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	//Insert additional processsing here
	for _, i := range raw {
		input = append(input, strings.Split(i, " "))
	}
	for _, i := range val {
		tmp := strings.Split(i, ": ")
		levl, _ := strconv.Atoi(tmp[1])
		values[tmp[0]] = levl
	}

	//insert Puzzle solution here
	sum := chip(values, input)

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}
