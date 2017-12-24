package main

import (
	"fmt"
	"math/rand"
)

func gameStart() Game {
	game := Game{}

	game.State = "setup"

	p1 := Player{
		Name:  "Bill",
		Ships: makePlayerShips(),
		Board: &GameBoard{},
	}
	//p1.Board.init()
	game.Player1 = &p1
	p2 := Player{
		Name:  "Bill",
		Ships: makePlayerShips(),
		Board: &GameBoard{},
	}
	//p2.Board.init()
	game.Player2 = &p2

	me = game.Player1
	name := readLine("Enter your name:")
	fmt.Println("Hello,", name)
	me.Name = name

	fmt.Println("This is your game board:")
	fmt.Println(render(me))
	fmt.Println("Start placing ships!")

	randomize := false
	// shipPlacementLoop:
	for _, ship := range me.Ships {
		fmt.Println("Place your", ship.Name, "(", ship.Length, "long)!")
	validCoordLoop:
		for { // loop until a valid coord is entered
			var coordString string
			var err error
			if randomize {
				x := rand.Intn(10)
				y := rand.Intn(10)

				coordString, err = coordsToString([2]int{x, y})
				if err != nil {
					fmt.Println("Eror randomizing input coordinate")
					randomize = false
					continue validCoordLoop
				}
			} else {

				coordString = readLine("Enter top-left coordinate:")
				switch coordString {
				case "rand", "random", "randomize":
					randomize = true
					continue validCoordLoop
				}
			}

			if len(coordString) == 2 {
				coords, err := stringToCoords(coordString)
				if err != nil {
					fmt.Println("Failed to parse coordinate:", err)
				} else {
				validDirectionLoop:
					for { // loop until a valid direction is entered
						var direction string
						fmt.Println("Coords:", coords)
						if randomize {
							dNum := rand.Intn(2)
							switch dNum {
							case 0:
								direction = "a"
							case 1:
								direction = "d"
							}
						} else {
							direction = readLine("Place your " + ship.Name + " down 'd' or across 'a' from " + coordString + " (x to change coord)")
						}
						switch direction {
						case "d", "D":
							fmt.Println("You chose Down!")
							// check for validity
							if coords[1]+ship.Length > 10 {
								// if the ship would go over the end of the grid, fail
								fmt.Println("Ship is too long to fit there")
								continue validCoordLoop
							}
							for i := 0; i < ship.Length; i++ {
								validPlacement := true
								if (me.Board.Grid[coords[0]][coords[1]+i] != nil) {
									cString, err := coordsToString([2]int{coords[0], coords[1] + i})
									if err != nil {
										fmt.Println("Something went wrong with your ship placement, try again:", err)
										continue validCoordLoop
									}
									fmt.Println("There is already a " + me.Board.Grid[coords[0]][coords[1]+i].Name + " at " +
										cString + ", try another location.")
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
								me.Board.Grid[coords[0]][coords[1]+i] = ship
							}
							break validDirectionLoop
						case "a", "A":
							fmt.Println("You chose Across!")
							// check for validity
							if coords[0]+ship.Length > 10 {
								// if the ship would go over the end of the grid, fail
								fmt.Println("Ship is too long to fit there")
								continue validCoordLoop
							}
							for i := 0; i < ship.Length; i++ {
								validPlacement := true
								if (me.Board.Grid[coords[0]+i][coords[1]] != nil) {
									cString, err := coordsToString([2]int{coords[0], coords[1] + i})
									if err != nil {
										fmt.Println("Something went wrong with your ship placement, try again:", err)
										continue validCoordLoop
									}
									fmt.Println("There is already a " + me.Board.Grid[coords[0]+i][coords[1]].Name + " at " +
										cString + ", try another location.")
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
								me.Board.Grid[coords[0]+i][coords[1]] = ship
							}
							break validDirectionLoop
						case "x", "X":
							fmt.Println("Ok, try a different one...")
							continue validCoordLoop // re-select coord
						default:
							fmt.Println("Invalid input ")
						}
					}
					fmt.Println(render(me))
					break validCoordLoop
				}
			}
			fmt.Println(coordString + " is not valid, please enter a valid coordinate (ex 'A2'):")
		}
	}

	game.State = "playing"

	return game
}

func makePlayerShips() []*Ship {
	ships := []*Ship{
		{
			Name:   "Carrier",
			Length: 5,
			Hits:   0,
		},
		{
			Name:   "Battleship",
			Length: 4,
			Hits:   0,
		},
		{
			Name:   "Submarine",
			Length: 3,
			Hits:   0,
		},
		{
			Name:   "Cruiser",
			Length: 3,
			Hits:   0,
		},
		{
			Name:   "Destroyer",
			Length: 2,
			Hits:   0,
		},
	}

	return ships
}
