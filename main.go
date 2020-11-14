package main

import (
	"math/rand"
	"mazes/core"
	"mazes/display"
	"mazes/generators"
	"time"
	"os"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	grid := core.NewGrid(6, 6)
	generators.SideWinder(grid)
	// displayer := display.MakePNGCreator("maze.png", 64, 5)
	displayer := display.MakeTerminalDisplay(os.Stdout)
	displayer.Display(grid)
}
