package main

import (
	"strconv"
	"errors"
)

type Player struct {
	Name  string
	Board *GameBoard
	Ships []*Ship
	ShotsFired [10][10]*bool
	ShotsReceived [10][10]*bool
	Input <-chan string
	Output chan<- string
}

// Startup script for a new Player
func (p *Player) Init() {
	p.Input = make(chan string)
	p.Output = make(chan string)

	// listen for commands
	go func() {
		for {
			msg := <-p.Input
			println(p.Name + ": RECEIVED: "+msg)
		}
	}()
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
