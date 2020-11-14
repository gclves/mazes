package main

import (
	"math/rand"
	"mazes/core"
	"mazes/display"
	"mazes/generators"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	grid := core.NewGrid(6, 6)
	generators.SideWinder(grid)
	display.TextDisplay(grid)
}
