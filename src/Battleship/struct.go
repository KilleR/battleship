package main

import (
	"strconv"
	"errors"
)

type Game struct {
	State   string // holds game state 'setup', 'playing' ...
	Player1 *Player
	Player2 *Player
}

func (g *Game) Connect() *Player {
	return g.Player1
}

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
