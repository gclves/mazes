package core

type Grid struct {
	Rows    int
	Columns int
	grid    [][]Cell
}

type RowWalker func(row []*Cell, index int)
type CellWalker func(cell *Cell, i, j int)

func (g Grid) SetAt(row, col int, c Cell) {
	g.grid[row][col] = c
}

func (g Grid) GetAt(row, col int) Cell {
	return g.grid[row][col]
}

func (g Grid) EachRow(cb RowWalker) {
	for i := 0; i < g.Rows; i++ {
		refs := make([]*Cell, 0)
		for j := 0; j < g.Columns; j++ {
			refs = append(refs, &g.grid[i][j])
		}
		cb(refs, i)
	}
}

func (g Grid) EachCell(cb CellWalker) {
	g.EachRow(func(row []*Cell, rowIndex int) {
		for i := 0; i < g.Columns; i++ {
			cb(row[i], rowIndex, i)
		}
	})
}

func NewGrid(rows, columns int) Grid {
	g := Grid{rows, columns, makeGrid(rows, columns)}
	return g
}

func makeGrid(rows, columns int) [][]Cell {
	grid := make([][]Cell, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]Cell, columns)
		for j := 0; j < columns; j++ {
			grid[i][j] = NewCell()
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			cell := &grid[i][j]
			if i > 0 {
				cell.North = &grid[i-1][j]
			}
			if i < rows-1 {
				cell.South = &grid[i+1][j]
			}
			if j > 0 {
				cell.West = &grid[i][j-1]
			}
			if j < columns-1 {
				cell.East = &grid[i][j+1]
			}
		}
	}

	return grid
}
