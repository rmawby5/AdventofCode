package main

import (
	"os"

	"Aoc.com/AdventOfCode2024/day7"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	Parse1Time, P1Time, P1Total := day7.Part1()
	Parse2Time, P2Time, P2Total := day7.Part2()
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
