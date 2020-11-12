package generators

import (
	"math/rand"
	"mazes/core"
)

func BinaryTree(g core.Grid) {
	g.EachCell(func(cell *core.Cell, _, _ int) {
		neighbours := make([]*core.Cell, 0)
		if cell.East != nil {
			neighbours = append(neighbours, cell.East)
		}
		if cell.South != nil {
			neighbours = append(neighbours, cell.South)
		}
		if len(neighbours) == 0 {
			return
		}
		index := rand.Intn(len(neighbours))
		cell.Link(neighbours[index])
	})
}
