package main

import "fmt"

func gameStart() Game {
	game := Game{}

	game.State = "setup"

	p1 := Player{
		Name:  "Bill",
		Ships: makePlayerShips(),
		Board: &GameBoard{},
	}
	//p1.Board.init()
	game.Player1 = p1
	p2 := Player{
		Name:  "Bill",
		Ships: makePlayerShips(),
		Board: &GameBoard{},
	}
	//p2.Board.init()
	game.Player2 = p2

	me = &game.Player1
	name := readLine("Enter your name:")
	fmt.Println("Hello,", name)
	me.Name = name

	fmt.Println("This is your game board:")
	render(me.Board)
	fmt.Println("Start placing ships!")

	// shipPlacementLoop:
	for _, ship := range me.Ships {
		fmt.Println("Place your",ship.Name,"(",ship.Length,"long)!")
		validCoordLoop:
		for { // loop until a valid coord is entered
			coordString := readLine("Enter top-left coordinate:")
			if len(coordString) == 2 {
				coords, err := stringToCoords(coordString)
				if err != nil {
					fmt.Println("Failed to parse coordinate:", err)
				} else {
					validDirectionLoop:
					for { // loop until a valid direction is entered
						fmt.Println("Coords:", coords)
						direction := readLine("Place your "+ship.Name+" down 'd' or across 'a' from "+coordString+" (x to change coord)")
						switch direction {
						case "d", "D":
							fmt.Println("You chose Down!")
							// check for validity
							for i := 0; i<ship.Length;i++ {
								validPlacement := true
								if(me.Board.Grid[coords[0]][coords[1]+i] != nil) {
									cString, err:= coordsToString([]int{coords[0],coords[1]+i})
									if err != nil {
										fmt.Println("Something went wrong with your ship placement, try again:",err)
										continue validCoordLoop
									}
									fmt.Println("There is already a "+me.Board.Grid[coords[0]][coords[1]+i].Name+" at "+
										cString+", try another location.")
									validPlacement = false
								}
								if !validPlacement {
									continue validCoordLoop // re-select coord
								}
							}
							// put ship down
							for i := 0; i<ship.Length;i++ {
								me.Board.Grid[coords[0]][coords[1]+i] = &ship
							}
							break validDirectionLoop
						case "a", "A":
							fmt.Println("You chose Across!")
							// check for validity
							for i := 0; i<ship.Length;i++ {
								validPlacement := true
								if(me.Board.Grid[coords[0]+i][coords[1]] != nil) {
									cString, err:= coordsToString([]int{coords[0],coords[1]+i})
									if err != nil {
										fmt.Println("Something went wrong with your ship placement, try again:",err)
										continue validCoordLoop
									}
									fmt.Println("There is already a "+me.Board.Grid[coords[0]+i][coords[1]].Name+" at "+
										cString+", try another location.")
									validPlacement = false
								}
								if !validPlacement {
									continue validCoordLoop // re-select coord
								}
							}
							// put ship down
							for i := 0; i<ship.Length;i++ {
								me.Board.Grid[coords[0]+i][coords[1]] = &ship
							}
							break validDirectionLoop
						case "x", "X":
							fmt.Println("Ok, try a different one...")
							continue validCoordLoop // re-select coord
						default:
							fmt.Println("Invalid input ")
						}
					}
					fmt.Println(render(me.Board))
					break validCoordLoop
				}
			}
			fmt.Println(coordString+" is not valid, please enter a valid coordinate (ex 'A2'):")
		}
	}

	return game
}

func makePlayerShips() []Ship {
	ships := []Ship{
		Ship{
			Name:   "Carrier",
			Length: 5,
			Hits:   0,
		},
		Ship{
			Name:   "Battleship",
			Length: 4,
			Hits:   0,
		},
		Ship{
			Name:   "Submarine",
			Length: 3,
			Hits:   0,
		},
		Ship{
			Name:   "Cruiser",
			Length: 3,
			Hits:   0,
		},
		Ship{
			Name:   "Destroyer",
			Length: 2,
			Hits:   0,
		},
	}

	return ships
}
