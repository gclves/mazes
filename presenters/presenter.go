package presenters

import "mazes/core"

type Presenter interface {
	Display(g core.Grid) error
}
