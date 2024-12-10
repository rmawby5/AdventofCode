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


//Part 2 WIP
func Part2() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day9/Input.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	//Insert additional processsing here
	var disk [][]int
	for i, j := range strings.Split(raw[0]) {
		if i%2 == 0 {
			blockSize, _ := strconv.Atoi(j)
			id := i / 2
		} else {
			blockSize, _ := strconv.Atoi(j)
			id := -1
		}
		block := []int{id, blockSize}
		disk = append(disk, block)
	}

	//inerst Puzzle solution here
	//create compacted disk
	disksize := len(disk) - 1
	for i, j := disksize; i > 0; i-- {
		if j[0] != -1 { //if block is file and no free space, attempt to move
			spaceIdx, spaceRem := getSpace(disk, j, i)
			if spaceIdx != -1 { //freespace is found
				tempDisk := disk[:spaceIdx]
				tempDisk = append(tempDisk, j)
				if spaceRem != 0 {
					newSpaceBlock := []int{-1, spaceRem}
					tempDisk = append(tempDisk, newSpaceBlock)
				}
				disk[i] := []int{-1, j[1]}
				tempDisk = append(tempDisk, disk[(spaceIdx+1):]...)
				tmp := copy(disk, tempDisk)
			}
		}
	}

	//calc checksum for compacted disk
	var chksum int64
	chksum = 0
	expIxd := 0
	for i, j := range disk {
		if j[0] != -1 { //ignore free space blocks
			chksum += int64(j[0] * ((j[1] * expIdx) + j[1]*(j[1]+1)*0.5)) //calulate the checksum for the specific block based on the expanded Eq indx, block size and file ID
			expIdx += j[1]
		}
	}

	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, chksum
}
