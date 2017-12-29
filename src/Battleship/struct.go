package main

import (
)

type GameBoard struct {
	Grid [10][10]*Ship
}

func (gb GameBoard) GetShipAt(coords [2]int) *Ship {
	return gb.Grid[coords[0]][coords[1]]
}

type Ship struct {
	Name        string
	Length      int
	Hits        int    // number of hits against it
	Location    string // top left coordinate
	Orientation string // 'h' or 'v' for horizontal or vertical
}

func (s *Ship) isDestroyed() bool {
	return s.Hits >= s.Length
}
