package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"mazes/core"
	"mazes/generators"
	"mazes/presenters"
	"os"
	"time"
)

type CLIArgs struct {
	TargetFile string
	WriteToPNG bool
	GridSize   int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	args := parseArgs()

	grid := core.NewGrid(args.GridSize, args.GridSize)
	generators.SideWinder(grid)

	// XXX what about .close()?
	out := openTargetFile(args.TargetFile)
	presenter := getPresenter(args, out)
	presenter.Display(grid)
}

func getPresenter(args CLIArgs, out io.Writer) presenters.Presenter {
	if args.WriteToPNG {
		return presenters.MakePNGDisplayer(out, 64, 5)
	} else {
		return presenters.MakeTextDisplayer(out)
	}
}

func parseArgs() CLIArgs {
	var args CLIArgs
	flag.BoolVar(&args.WriteToPNG, "png", false, "whether to write an image")
	// XXX: this must be positive
	flag.IntVar(&args.GridSize, "grid-size", 6, "the size of the grid")
	flag.Parse()

	positionalArgs := flag.Args()
	if len(positionalArgs) == 1 {
		args.TargetFile = positionalArgs[0]
	} else if len(positionalArgs) == 0 {
		args.TargetFile = "-"
	} else {
		flag.Usage()
		os.Exit(1)
	}

	return args
}

// maybe this should return an error
func openTargetFile(path string) io.Writer {
	if path == "-" {
		return os.Stdout
	}
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Failed to open %q for writing: %v", path, err)
	}
	return f
}
