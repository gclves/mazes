package core

import "testing"

func TestCellLinking(t *testing.T) {
	one := NewCell()
	two := NewCell()
	one.East = &two
	two.West = &one

	if one.IsLinked(two) {
		t.Errorf("Expected {one, two} not to be linked")
	}
	if two.IsLinked(one) {
		t.Errorf("Expected {two, one} not to be linked")
	}
	one.Link(&two)
	if !one.IsLinked(two) {
		t.Errorf("Expected one to be linked to two")
	}
	if !two.IsLinked(one) {
		t.Errorf("Expected two to be linked to one")
	}
}

