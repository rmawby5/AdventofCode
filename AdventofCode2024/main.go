package main

import (
	"os"

	"Aoc.com/AdventOfCode2024/day5"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	day5.Part2()

	Parse1Time, P1Time, P1Total := day5.Part1()
	Parse2Time, P2Time, P2Total := day5.Part2()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Part", "Parse Time", "Part Time", "Result"})
	t.AppendRows([]table.Row{
		{1, Parse1Time, P1Time, P1Total},
	})
	t.AppendSeparator()
	t.AppendRow([]interface{}{2, Parse2Time, P2Time, P2Total})
	t.Render()

}
