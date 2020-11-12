package core

import "fmt"

type Cell struct {
	id int
	North *Cell
	South *Cell
	East  *Cell
	West  *Cell
	links []int
}

var count = 0

func NewCell() Cell {
	count += 1
	return Cell{count, nil, nil, nil, nil, make([]int, 0)}
}

func (c *Cell) Link(other *Cell) {
	c.links = append(c.links, other.id)
	other.links = append(other.links, c.id)
}

func (c Cell) IsLinked(other Cell) bool {
	for _, link := range c.links {
		if link == other.id {
			return true
		}
	}
	return false
}

func (c Cell) Debug() string {
	return fmt.Sprintf("%d", c.id)
}

