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

type Player struct {
	Name  string
	Board *GameBoard
	Ships []*Ship
	ShotsFired [10][10]*bool
	ShotsReceived [10][10]*bool
}

func (p Player) ShipStatus() string {
	output := ""

	for _, s := range p.Ships {
		output += s.Name + " (" + s.Location + " "
		switch s.Orientation {
		case "h":
			output += "across"
		case "v":
			output += "down"
		default:
			output += "unknown"
		}
		output += "): " + strconv.Itoa(s.Hits) + "hits"
		if s.isDestroyed() {
			output += " (sunk)"
		}
		output += "\n"
	}

	return output
}

func (p *Player) FireAt(coords [2]int) (string, error) {
	output := ""

	if p.ShotsFired[coords[0]][coords[1]] != nil{
		coordString, err := coordsToString(coords)
		if err != nil {
			return output, errors.New("You have already fired at: INVALID: "+err.Error());
		}
		return output, errors.New("You have already fired at: "+coordString)
	}

	res, err := p.GetShotAt(coords)
	if err != nil {
		return output, errors.New("Target error: "+err.Error())
	}
	if res == "Miss!" {
		b := false
		p.ShotsFired[coords[0]][coords[1]] = &b
	} else {
		b := true
		p.ShotsFired[coords[0]][coords[1]] = &b
	}

	output += "Result: "+res
	return output, nil
}

func (p *Player) GetShotAt(coords [2]int) (string, error) {
	output := ""

	// check for already received shot at this coord
	if p.ShotsReceived[coords[0]][coords[1]] != nil {
		coordString, err := coordsToString(coords)
		if err != nil {
			return output, errors.New("You have already fired at: INVALID: "+err.Error());
		}
		return output, errors.New("You have already fired at: "+coordString)
	}

	shipAtCoord := me.Board.GetShipAt(coords)
	if shipAtCoord != nil {
		b := true
		p.ShotsReceived[coords[0]][coords[1]] = &b // set received shot at this coord
		output += "Hit!"
		shipAtCoord.Hits++
		if shipAtCoord.isDestroyed() {
			output += "\nYou sank the " + shipAtCoord.Name + "!"
		}
	} else {
		b := false
		p.ShotsReceived[coords[0]][coords[1]] = &b // set received shot at this coord
		output += "Miss!"
	}

	return output, nil
}

func (p *Player) ShipsRemaining() int {
	remaining := 0

	for _,s := range p.Ships {
		if !s.isDestroyed() {
			remaining++
		}
	}

	return remaining
}

type GameBoard struct {
	Grid [10][10]*Ship
}

func (gb GameBoard) GetShipAt(coords [2]int) *Ship {
	return gb.Grid[coords[0]][coords[1]]
}

//func (b *GameBoard) init() {
//	b.Grid = make([][]*Ship, 10)
//	for i := range b.Grid {
//		b.Grid[i] = make([]*Ship, 10)
//	}
//}

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
