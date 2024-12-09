package day9

import (
	"slices"
	"strconv"
	"strings"
	"time"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

func diskExpander(input []string) []string {
	disk := strings.Split(input[0], "")
	//Even index numbers = files, odd idex numbers = spaces
	var expandedDisk []string
	for i, j := range disk {
		if i%2 == 0 {
			char, _ := strconv.Atoi(j)
			mark := strconv.Itoa(i / 2)
			for k := 0; k < char; k++ {
				expandedDisk = append(expandedDisk, mark)
			}
		} else {
			char, _ := strconv.Atoi(j)
			for k := 0; k < char; k++ {
				expandedDisk = append(expandedDisk, ".")
			}
		}
	}
	return expandedDisk
}

func count(ls []string, c string) int {
	//slice element counting function, change input type as needed
	count := 0
	for _, i := range ls {
		if i == c {
			count++
		}
	}
	return count
}

func checksum(ls []string) int64 {
	var sum int64
	sum = 0
	for i, j := range ls {
		l, _ := strconv.Atoi(j)
		sum += int64(i * l)
	}
	return int64(sum)
}

func Part1() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day9/Input.txt")
	ParseTime := time.Since(ParseStart)

	splitDisk := diskExpander(raw)
	P1Start := time.Now()
	disksize := len(splitDisk) - 1
	dots := count(splitDisk, ".")

	for i := disksize; (disksize - dots) < i; i-- {
		if splitDisk[i] != "." { //move to start
			lastElmnt := splitDisk[i]
			firstSpace := slices.Index(splitDisk, ".")
			splitDisk[firstSpace] = lastElmnt
			splitDisk = splitDisk[:i]

		}
	}
	chksum := checksum(splitDisk)
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, chksum

}

/*
//Part 2 WIP
func Part2() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day9/Input.txt")
	disk := diskExpander(raw)
	//Insert additional processsing here
	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//inerst Puzzle solution here
	disksize := len(disk) - 2
	buffer := []string{disk[len(disk)-1]}
	for i := disksize; i > -1; i-- {
		if slices.Contains(buffer, disk[i]) {
			buffer = append(buffer, disk[i])
		}

	}

	chksum := 0

	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, int64(chksum)
}
*/
