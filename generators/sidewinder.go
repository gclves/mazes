package generators

import (
	"fmt"
	"math/rand"
	"mazes/core"
)

func SideWinder(g core.Grid) {
	g.EachRow(func(row []*core.Cell, _ int) {
		run := make([]*core.Cell, 0)
		for i := 0; i < g.Columns; i++ {
			cell := row[i]
			run = append(run, cell)

			shouldBranchDown := cell.East == nil || (cell.South != nil && rand.Intn(2) == 0)
			if shouldBranchDown {
				targetIndex := rand.Intn(len(run))
				target := run[targetIndex]
				if target.South != nil {
					target.Link(target.South)
				}
				run = make([]*core.Cell, 0)
			} else {
				if cell.East != nil {
					cell.Link(cell.East)
				}
			}
		}
	})
}

func log(cell *core.Cell, target *core.Cell) {
	fmt.Printf("Linking %s to %s\n", cell.Debug(), target.Debug())
}
