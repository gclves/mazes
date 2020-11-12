package terminal

import (
	"fmt"
	"mazes/core"
)

func TextDisplay(g core.Grid) {
	printUpperLine(g)
	printRows(g)
	printBottomLine(g)
	fmt.Println()
}

func printUpperLine(g core.Grid) {
	columns := ""
	for i := 0; i < g.Columns-1; i++ {
		columns += "───┬"
	}
	fmt.Printf("┌%s───┐\n", columns)
}

func printRows(g core.Grid) {
	g.EachRow(func(cells []*core.Cell, _ int) {
		fmt.Print("│")
		bottomRow := "├"

		for i := 0; i < len(cells); i++ {
			cell := cells[i]
			fmt.Print("   ")
			if cell.East != nil && cell.IsLinked(*cell.East) {
				fmt.Print(" ")
			} else {
				fmt.Print("│")
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
		fmt.Println()
		if cells[0].South != nil {
			fmt.Println(bottomRow)
		}
	})
}

func printBottomLine(g core.Grid) {
	fmt.Print("└")
	for i := 0; i < g.Columns-1; i++ {
		fmt.Print("───┴")
	}
	fmt.Print("───┘")
}
