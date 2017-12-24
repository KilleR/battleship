package main

import (
	"strconv"
	"errors"
	"fmt"
	"math/rand"
)

type Player struct {
	Name          string
	IsAI          bool
	ID            string
	game          *Game // back reference to the game the player is in, for checking opponent and whether it's Player's turn
	Board         *GameBoard
	Ships         []*Ship
	ShotsFired    [10][10]*bool
	ShotsReceived [10][10]*bool
	Input         chan string
	Output        chan string
}

// Startup script for a new Player
func (p *Player) Init() {
	p.Input = make(chan string)
	p.Output = make(chan string)

	if p.IsAI {
		p.Name = "BOT"
	} else {
		name := p.ReadLine("Enter your name:")
		p.Name = name
		p.ID = name //@TODO: add real ID generation here
		p.Output <- "Hello, " + name
	}

	p.DoShipPlacement()
	p.game.ReadyCheck(p)
	p.Output <- "Done placing ships!"

	// listen for commands
	go func() {
	CmdLoop:
		for {
			msg := p.ReadLine("")
			//println(p.Name + ": RECEIVED: " + msg)
			switch msg {
			case "hi":
				p.Output <- fmt.Sprint("Hello!")
			case "aime":
				p.Output <- p.game.MakeAIPlayer(p)
			case "turn":
				if p.game.PlayerTurn == nil {
					p.Output <- "Nobody"
				} else {
					p.Output <- p.game.PlayerTurn.Name
				}
			case "players":
				if p.game.Player1 != nil {
					p.Output <- "Player 1: "+p.game.Player1.Name
				}
				if p.game.Player2 != nil {
					p.Output <- "Player 2: "+p.game.Player2.Name
				}
			case "render":
				if p.game.State == "" {
					p.Output <- fmt.Sprint("Game has not started, there is no board to render")
				} else {
					p.Output <- fmt.Sprint(render(p))
				}
			case "fired":
				if p.game.State == "" {
					p.Output <- fmt.Sprint("Game has not started, there is no board to render")
				} else {
					p.Output <- fmt.Sprint(renderFired(p))
				}
			case "ships":
				p.Output <- fmt.Sprint(p.ShipStatus())
			case "pass":
				// check if it's this player's turn
				if p.game.PlayerTurn == nil {
					p.Output <- fmt.Sprint("Waiting for the other player")
					continue CmdLoop
				} else if p.game.PlayerTurn != p {
					p.Output <- fmt.Sprint("It is not your turn")
					continue CmdLoop
				}
				p.PassTurn()
			default:
				isCoord := false
				if len(msg) == 2 {
					// on an exactly 2 character input, check if it's a coordinate
					coords, err := stringToCoords(msg)
					if err != nil {
						p.Output <- fmt.Sprint("Failed to parse coordinate:", err)
					} else {
						// check if it's this player's turn
						if p.game.PlayerTurn == nil {
							p.Output <- fmt.Sprint("Waiting for the other player")
							continue CmdLoop
						} else if p.game.PlayerTurn != p {
							p.Output <- fmt.Sprint("It is not your turn")
							continue CmdLoop
						}
						p.Output <- fmt.Sprintf("coords entered: %v", coords)
						isCoord = true
						// since it's a coord, I need to know what I'm doing with it
						switch p.game.State {
						case "playing":
							p.Output <- fmt.Sprintf("Shot fired at: %s", msg)
							fireResult, err := p.FireAt(coords)
							if err != nil {
								p.Output <- fmt.Sprintf("Firing error: %s", err)
								continue CmdLoop
							}
							p.Output <- fmt.Sprint(fireResult)
							p.PassTurn() // end turn if we get a "successful" shot
						}
					}
				}
				// if it was not a coord, print the unknown command text
				if !isCoord {
					p.Output <- fmt.Sprintf("Unknown command: %s", msg)
				}
			}

		}

	}()
}

// Big method functions
func (p *Player) ReadLine(prompt string) string {
	if prompt != "" {
		p.Output <- prompt
	}
	p.Output <- ">"

	return <-p.Input
}

func (p *Player) DoShipPlacement() {
	randomize := false
	// if we're an AI player, we always pick random
	if p.IsAI {
		randomize = true;
	}
	// shipPlacementLoop:
	for _, ship := range p.Ships {
		p.Output <- fmt.Sprintf("Place your %s (%d long)!", ship.Name, ship.Length)
	validCoordLoop:
		for { // loop until a valid coord is entered
			var coordString string
			var err error
			if randomize {
				x := rand.Intn(10)
				y := rand.Intn(10)

				coordString, err = coordsToString([2]int{x, y})
				if err != nil {
					p.Output <- fmt.Sprint("Error randomizing input coordinate")
					randomize = false
					continue validCoordLoop
				}
			} else {

				coordString = p.ReadLine("Enter top-left coordinate:")
				switch coordString {
				case "rand", "random", "randomize":
					randomize = true
					continue validCoordLoop
				}
			}

			if len(coordString) == 2 {
				coords, err := stringToCoords(coordString)
				if err != nil {
					p.Output <- fmt.Sprintf("Failed to parse coordinate: %s", err)
				} else {
				validDirectionLoop:
					for { // loop until a valid direction is entered
						var direction string
						p.Output <- fmt.Sprint("Coords:", coords)
						if randomize {
							dNum := rand.Intn(2)
							switch dNum {
							case 0:
								direction = "a"
							case 1:
								direction = "d"
							}
						} else {
							direction = p.ReadLine("Place your " + ship.Name + " down 'd' or across 'a' from " + coordString + " (x to change coord)")
						}
						switch direction {
						case "d", "D":
							p.Output <- fmt.Sprint("You chose Down!")
							// check for validity
							if coords[1]+ship.Length > 10 {
								// if the ship would go over the end of the grid, fail
								p.Output <- fmt.Sprint("Ship is too long to fit there")
								continue validCoordLoop
							}
							for i := 0; i < ship.Length; i++ {
								validPlacement := true
								if (p.Board.Grid[coords[0]][coords[1]+i] != nil) {
									cString, err := coordsToString([2]int{coords[0], coords[1] + i})
									if err != nil {
										p.Output <- fmt.Sprintf("Something went wrong with your ship placement, try again: %s", err)
										continue validCoordLoop
									}
									p.Output <- fmt.Sprintf("There is already a %s at %s, try another location.",
										p.Board.Grid[coords[0]][coords[1]+i].Name, cString)
									validPlacement = false
								}
								if !validPlacement {
									continue validCoordLoop // re-select coord
								}
							}
							// put ship down
							ship.Orientation = "v"
							ship.Location = coordString
							for i := 0; i < ship.Length; i++ {
								p.Board.Grid[coords[0]][coords[1]+i] = ship
							}
							break validDirectionLoop
						case "a", "A":
							p.Output <- fmt.Sprint("You chose Across!")
							// check for validity
							if coords[0]+ship.Length > 10 {
								// if the ship would go over the end of the grid, fail
								p.Output <- fmt.Sprint("Ship is too long to fit there")
								continue validCoordLoop
							}
							for i := 0; i < ship.Length; i++ {
								validPlacement := true
								if (p.Board.Grid[coords[0]+i][coords[1]] != nil) {
									cString, err := coordsToString([2]int{coords[0], coords[1] + i})
									if err != nil {
										p.Output <- fmt.Sprintf("Something went wrong with your ship placement, try again: %s", err)
										continue validCoordLoop
									}
									p.Output <- fmt.Sprintf("There is already a %s at %s, try another location.",
										p.Board.Grid[coords[0]+i][coords[1]].Name, cString)
									validPlacement = false
								}
								if !validPlacement {
									continue validCoordLoop // re-select coord
								}
							}
							// put ship down
							ship.Orientation = "h"
							ship.Location = coordString
							for i := 0; i < ship.Length; i++ {
								p.Board.Grid[coords[0]+i][coords[1]] = ship
							}
							break validDirectionLoop
						case "x", "X":
							p.Output <- fmt.Sprint("Ok, try a different one...")
							continue validCoordLoop // re-select coord
						default:
							p.Output <- fmt.Sprint("Invalid input ")
						}
					}
					p.Output <- fmt.Sprint(render(p))
					break validCoordLoop
				}
			}
			p.Output <- fmt.Sprintf("%s is not valid, please enter a valid coordinate (ex 'A2'):", coordString)
		}
	}

}

// Getter/utility functions
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

func (p *Player) BeginTurn() {
	if (p.IsAI) {
		p.PassTurn()
	}
	p.Output <- "It is your turn"
}

func (p *Player) PassTurn() {
	p.game.EndPlayerTurn(p)
}

func (p *Player) FireAt(coords [2]int) (string, error) {
	output := ""

	if p.ShotsFired[coords[0]][coords[1]] != nil {
		coordString, err := coordsToString(coords)
		if err != nil {
			return output, errors.New("You have already fired at: INVALID: " + err.Error());
		}
		return output, errors.New("You have already fired at: " + coordString)
	}

	res, err := p.game.GetOpponent(p).GetShotAt(coords)
	if err != nil {
		return output, errors.New("Target error: " + err.Error())
	}
	if res == "Miss!" {
		b := false
		p.ShotsFired[coords[0]][coords[1]] = &b
	} else {
		b := true
		p.ShotsFired[coords[0]][coords[1]] = &b
	}

	output += "Result: " + res
	return output, nil
}

func (p *Player) GetShotAt(coords [2]int) (string, error) {
	output := ""

	// check for already received shot at this coord
	if p.ShotsReceived[coords[0]][coords[1]] != nil {
		coordString, err := coordsToString(coords)
		if err != nil {
			return output, errors.New("You have already fired at: INVALID: " + err.Error());
		}
		return output, errors.New("You have already fired at: " + coordString)
	}

	shipAtCoord := p.Board.GetShipAt(coords)
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

	for _, s := range p.Ships {
		if !s.isDestroyed() {
			remaining++
		}
	}

	return remaining
}
