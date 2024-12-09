package main

import (
	"os"

	"Aoc.com/AdventOfCode2024/day9"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {

	Parse1Time, P1Time, P1Total := day9.Part1()
	//Parse2Time, P2Time, P2Total := day9.Part2()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Part", "Parse Time", "Part Time", "Result"})
	t.AppendRows([]table.Row{
		{1, Parse1Time, P1Time, P1Total},
	})
	t.AppendSeparator()
	//t.AppendRow([]interface{}{2, Parse2Time, P2Time, P2Total})
	t.Render()

}

//function template
/*
func Part1() (time.Duration, time.Duration, int64) {
	ParseStart := time.Now()
	raw := fileparse.FileParse("day9/Input.txt")
	//Insert additional processsing here

	ParseTime := time.Since(ParseStart)
	P1Start := time.Now()
	//inerst Puzzle solution here


	P1Time := time.Since(P1Start)
	return ParseTime, P1Time, chksum
}
*/
