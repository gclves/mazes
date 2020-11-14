package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"mazes/core"
	"mazes/display"
	"mazes/generators"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	grid := core.NewGrid(6, 6)
	generators.SideWinder(grid)

	writeToPNGPtr := flag.Bool("png", false, "whether to write an image")
	flag.Parse()

	positionalArgs := flag.Args()
	if len(positionalArgs) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	// XXX what about .close()?
	targetFile := positionalArgs[0]
	out := openTargetFile(targetFile)
	var displayer display.Displayer
	if *writeToPNGPtr {
		displayer = display.MakePNGDisplayer(out, 64, 5)
	} else {
		displayer = display.MakeTextDisplayer(out)
	}
	displayer.Display(grid)
}

func openTargetFile(targetFile string) io.Writer {
	if targetFile == "-" {
		return os.Stdout
	}
	f, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("Failed to open %q for writing: %v", targetFile, err)
	}
	return f
}
