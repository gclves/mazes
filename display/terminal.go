package display

import (
	"fmt"
	"io"
	"mazes/core"
)

type TerminalDisplayer struct {
	writer io.Writer
}

func MakeTerminalDisplay(writer io.Writer) TerminalDisplayer {
	return TerminalDisplayer{writer}
}

func (d TerminalDisplayer) Display(g core.Grid) {
	d.printUpperLine(g)
	d.printRows(g)
	d.printBottomLine(g)
	fmt.Fprintln(d.writer)
}

func (d TerminalDisplayer) printUpperLine(g core.Grid) {
	columns := ""
	for i := 0; i < g.Columns-1; i++ {
		columns += "───┬"
	}
	fmt.Fprintf(d.writer, "┌%s───┐\n", columns)
}

func (d TerminalDisplayer) printRows(g core.Grid) {
	g.EachRow(func(cells []*core.Cell, _ int) {
		fmt.Fprint(d.writer, "│")
		bottomRow := "├"

		for i := 0; i < len(cells); i++ {
			cell := cells[i]
			fmt.Fprint(d.writer, "   ")
			if cell.East != nil && cell.IsLinked(*cell.East) {
				fmt.Fprint(d.writer, " ")
			} else {
				fmt.Fprint(d.writer, "│")
			}
			if cell.South != nil && cell.IsLinked(*cell.South) {
				bottomRow += "   "
			} else {
				bottomRow += "───"
			}
			if cell.South != nil {
				if cell.East == nil {
					bottomRow += "┤"
				} else {
					bottomRow += "┼"
				}
			}
		}
		fmt.Fprintln(d.writer)
		if cells[0].South != nil {
			fmt.Fprintln(d.writer, bottomRow)
		}
	})
}

func (d TerminalDisplayer) printBottomLine(g core.Grid) {
	fmt.Fprint(d.writer, "└")
	for i := 0; i < g.Columns-1; i++ {
		fmt.Fprint(d.writer, "───┴")
	}
	fmt.Fprint(d.writer, "───┘")
}
