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
		coord := readLine("Enter top-left coordinate:")
		validCoordLoop:
		for { // loop until a valid coord is entered
			if len(coord) == 2 {
				coords := coordRex.FindAllStringSubmatch(coord, -1)
				if len(coords) > 0 {
					break validCoordLoop
				}
			}
			coord = readLine("Please entere a valid coordinate (ex 'A2'):")
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
