package display

import "mazes/core"

type Displayer interface {
	Display(g core.Grid)
}

