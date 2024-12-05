package day5

import (
	"fmt"
	"strconv"
	"strings"

	fileparse "Aoc.com/AdventOfCode2024/FileParse"
)

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

func Part1() {
	//get file inputs and generate hash map
	var counter int64
	guide := fileparse.FileParse("day5/InputGuide.txt")
	pages := fileparse.FileParse("day5/Input.txt")
	pageGuide := make(map[string][]string)
	for _, row := range guide {
		firstpage := strings.Split(row, "|")[0]
		laterpage := strings.Split(row, "|")[1]
		i, ok := pageGuide[firstpage]
		if ok {
			pageGuide[firstpage] = append(i, laterpage)
		} else { //first entry of preceeding page in guide
			var val []string
			val = append(val, laterpage)
			pageGuide[firstpage] = val

		}
	}

	for _, i := range pages {
		var buffer []string
		line := strings.Split(i, ",")
		var validity bool
		for _, j := range line {
			invaildPg := pageGuide[j]
			for _, k := range invaildPg {
				if count(buffer, k) != 0 {
					//pagefile is invalid
					validity = false
					break
				}
				validity = true
			}
			//page num is valid
			if validity {
				buffer = append(buffer, j)
			} else {
				break
			}

		}
		if validity {

			midVal, _ := strconv.Atoi(line[len(line)/2])
			counter += int64(midVal)
		}

	}
	fmt.Println(counter)

}
