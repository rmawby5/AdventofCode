package day9

import (
	"fmt"
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

func getSpace(dsk [][]int, block []int, rightlim int) (int, int) {
	reqSpace := block[1]
	spaceBlockIdx := -1
	spaceRemaining := 0
	for i := 0; i < rightlim; i++ {
		if dsk[i][0] == -1 && dsk[i][1] >= reqSpace {
			spaceBlockIdx = i
			spaceRemaining = dsk[i][1] - reqSpace
			break
		}
	}
	return spaceBlockIdx, spaceRemaining
}

func Part1() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day9/Input.txt")
	ParseTime := time.Since(ParseStart)

	P1Start := time.Now()
	splitDisk := diskExpander(raw)
	disksize := len(splitDisk) - 1
	fmt.Println(disksize)
	dots := count(splitDisk, ".")
	firstSpace := 0
	for i := disksize; (disksize - dots) < i; i-- {
		if splitDisk[i] != "." { //move to start
			lastElmnt := splitDisk[i]
			openSpace := 0
			for l := firstSpace; l < i; l++ {
				if splitDisk[l] == "." {
					openSpace = l
					break
				}
			}
			firstSpace = openSpace
			splitDisk[openSpace] = lastElmnt
			splitDisk = splitDisk[:i]
		}
	}
	chksum := checksum(splitDisk)
	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, chksum
}

func Part2() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day9/Input.txt")
	ParseTime := time.Since(ParseStart)
	P2Start := time.Now()
	var disk [][]int
	id := 0
	blockSize := 0
	for i, j := range strings.Split(raw[0], "") {
		if i%2 == 0 {
			blockSize, _ = strconv.Atoi(j)
			id = i / 2
		} else {
			blockSize, _ = strconv.Atoi(j)
			id = -1
		}
		block := []int{id, blockSize}
		disk = append(disk, block)
	}
	//create compact disk
	disksize := len(disk) - 1
	for i := disksize; i > 0; i-- {
		j := disk[i]
		if j[0] != -1 { //if block is file and no free space, attempt to move
			spaceIdx, spaceRem := getSpace(disk, j, i)
			if spaceIdx != -1 { //freespace is found
				disk[spaceIdx] = j
				disk[i] = []int{-1, j[1]}
				if spaceRem != 0 {
					newSpaceBlock := []int{-1, spaceRem}
					disk = slices.Insert(disk, spaceIdx+1, newSpaceBlock)
				}
			}
		}
	}
	//calc checksum for compacted disk
	var chksum int64
	chksum = 0
	expIdx := 0
	for _, j := range disk {
		if j[0] != -1 { //ignore free space blocks
			chksum += int64(j[0] * ((j[1] * (expIdx - 1)) + j[1]*(j[1]+1)/2)) //calulate the checksum for the specific block based on the expande		d Eq indx, block size and file ID
			expIdx += j[1]
		} else {
			expIdx += j[1]
		}
	}
	P2Time := time.Since(P2Start)
	return ParseTime, P2Time, chksum
}
