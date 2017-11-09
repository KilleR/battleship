package main

type Game struct {
	Player1 Player
	Player2 Player
}

type Player struct {
	Name string
	Board *GameBoard
	Ships []Ship
}

type GameBoard struct {
	Grid [10][10]*Ship
}

//func (b *GameBoard) init() {
//	b.Grid = make([][]*Ship, 10)
//	for i := range b.Grid {
//		b.Grid[i] = make([]*Ship, 10)
//	}
//}

type Ship struct {
	Name string
	Length int
	Hits int // number of hits against it
	Location string // top left coordinate
	Orientation string // 'h' or 'v' for horizontal or vertical
}

func (s *Ship) isDestroyed() bool {
	return s.Hits >= s.Length
}
