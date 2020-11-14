package display

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"mazes/core"
	"os"
)

type PNGCreator struct {
	fileName string
	cellSize int
	wallThickness int
}

func MakePNGCreator(fileName string, cellSize, wallThickness int) PNGCreator {
	return PNGCreator{fileName, cellSize, wallThickness}
}

func (c PNGCreator) Display(g core.Grid) {
	img := c.makeImage(g)
	f, err := os.Create(c.fileName)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func (c PNGCreator) makeImage(g core.Grid) *image.RGBA {
	width := g.Columns*c.cellSize + (1 * c.wallThickness)
	height := g.Rows*c.cellSize + (1 * c.wallThickness)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bg := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)

	g.EachCell(func(cell *core.Cell, row, col int) {
		if row == 0 {
			x1 := col * c.cellSize
			y1 := 0
			x2 := (col + 1) * c.cellSize
			y2 := c.wallThickness
			drawLine(img, x1, y1, x2, y2)
		}

		if col == 0 {
			x1 := 0
			y1 := row * c.cellSize
			x2 := c.wallThickness
			y2 := (row+1)*c.cellSize + c.wallThickness
			drawLine(img, x1, y1, x2, y2)
		}

		if cell.South == nil || !cell.IsLinked(*cell.South) {
			x1 := col * c.cellSize
			y1 := (row + 1) * c.cellSize
			x2 := (col+1)*c.cellSize + c.wallThickness
			y2 := (row+1)*c.cellSize + c.wallThickness
			drawLine(img, x1, y1, x2, y2)
		}

		if cell.East == nil || !cell.IsLinked(*cell.East) {
			x1 := (col + 1) * c.cellSize
			y1 := row * c.cellSize
			x2 := (col+1)*c.cellSize + c.wallThickness
			y2 := (row + 1) * c.cellSize
			drawLine(img, x1, y1, x2, y2)
		}
	})

	return img
}

func drawLine(img draw.Image, x1, y1, x2, y2 int) {
	fg := color.RGBA{0, 0, 255, 255}
	wall := image.Rect(x1, y1, x2, y2)
	draw.Draw(img, wall, &image.Uniform{fg}, image.ZP, draw.Src)
}
