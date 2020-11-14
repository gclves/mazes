package main

import (
	"io"
	"log"
	"math/rand"
	"mazes/core"
	"mazes/display"
	"mazes/generators"
	"os"
	"time"
)

const printPNG = false

func main() {
	rand.Seed(time.Now().UnixNano())
	grid := core.NewGrid(6, 6)
	generators.SideWinder(grid)

	targetFile := "-"

	var displayer display.Displayer
	if printPNG {
		displayer = WriteToImage(targetFile)
	} else {
		displayer = display.MakeTerminalDisplay(os.Stdout)
	}
	displayer.Display(grid)
}

func WriteToImage(targetFile string) display.Displayer {
	var out io.Writer
	if targetFile == "-" {
		out = os.Stdout
	} else {
		f, err := os.Create(targetFile)
		if err != nil {
			log.Fatalf("Failed to open %q for writing: %v", targetFile, err)
		}
		out = f
	}

	return display.MakePNGCreator(out, 64, 5)
}
