package main

import (
	"os"

	"Aoc.com/AdventOfCode2024/day16"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	/*
		days := map[int]func() (time.Duration, time.Duration, int){
			2:  day1.Part2,
			3:  day2.Part1,
			4:  day2.Part2,
			5:  day3.Part1,
			6:  day3.Part1,
			7:  day4.Part1,
			8:  day4.Part2,
			9:  day5.Part1,
			10: day5.Part2,
			11: day6.Part1,
			12: day6.Part2,
			13: day7.Part1,
			14: day7.Part2,
			15: day8.Part1,
			16: day8.Part2,
			17: day9.Part1,
			18: day9.Part2,
			19: day10.Part1,
			20: day10.Part2,
			21: day11.Part1,
			22: day11.Part2,
			23: day12.Part1,
			24: day12.Part2,
			25: day13.Part1,
			26: day13.Part2,
			27: day14.Part1,
			28: day14.Part2,
			29: day15.Part1,
			31: day17.Part2,
			32: day18.Part1,
			34: day19.Part1,
			35: day19.Part2,
			36: day20.Part1,
			37: day20.Part2,
			38: day22.Part1,
			39: day22.Part2,
			40: day24.Part1,
			41: day25.Part1,
		}
		timeTotalStart := time.Now()

		Parse1Time, P1Time, P1Total := day1.Part1()
	*/

	Parse1Time, P1Time, P1Total := day16.Part1()
	//Parse2Time, P2Time, P2Total := day16.Part2()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Part", "Parse Time", "Part Time", "Result"})
	t.AppendRows([]table.Row{
		{1, Parse1Time, P1Time, P1Total},
	})
	t.AppendSeparator()
	//t.AppendRow([]interface{}{2, "Total Time", TotalTime})
	t.Render()

	/*
		for i := 2; i <= 29; i++ {
			ParTime, PTime, res := days[i]()
			t.AppendSeparator()
			t.AppendRow([]interface{}{i, ParTime, PTime, res})
		}

		Parse17Time, P17Time, P17Total := day17.Part1()
		t.AppendSeparator()
		t.AppendRow([]interface{}{30, Parse17Time, P17Time, P17Total})

		for i := 31; i <= 32; i++ {
			ParTime, PTime, res := days[i]()
			t.AppendSeparator()
			t.AppendRow([]interface{}{i, ParTime, PTime, res})
		}

		Parse19Time, P19Time, P19Total := day18.Part2()
		t.AppendSeparator()
		t.AppendRow([]interface{}{33, Parse19Time, P19Time, P19Total})

		for i := 34; i <= 41; i++ {
			ParTime, PTime, res := days[i]()
			t.AppendSeparator()
			t.AppendRow([]interface{}{i, ParTime, PTime, res})
		}

		TotalTime := time.Since(timeTotalStart)
			t.AppendSeparator()
		t.AppendRow([]interface{}{42, "Total Time", TotalTime})
			t.Render()
	*/

}

//function template
/*
func Part1() (time.Duration, time.Duration, int) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day9/TestInput.txt")

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()

	//Insert additional processsing here


	//insert Puzzle solution here


	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, sum
}
*/
