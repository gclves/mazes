package main

import (
	"math/rand"
	"mazes/core"
	"mazes/display"
	"mazes/generators"
	"time"
	"log"
	"os"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	grid := core.NewGrid(6, 6)
	generators.SideWinder(grid)

	targetFile := "maze.png"
	f, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("Failed to open %q for writing: %v", targetFile, err)
	}
	defer f.Close()

	displayer := display.MakePNGCreator(f, 64, 5)
	// displayer := display.MakeTerminalDisplay(os.Stdout)
	displayer.Display(grid)
}
