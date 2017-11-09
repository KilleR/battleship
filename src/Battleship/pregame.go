package main

func gameStart() Game{
	game := Game{}

	p1 := Player{
		Name: "Bill",
		Ships: makePlayerShips(),
		Board: &GameBoard{},
	}
	//p1.Board.init()
	game.Player1 = p1
	p2 := Player{
		Name: "Bill",
		Ships: makePlayerShips(),
		Board: &GameBoard{},
	}
	//p2.Board.init()
	game.Player2 = p2

	return game
}

func makePlayerShips() []Ship {
	ships := []Ship{
		Ship{
			Name: "Carrier",
			Length: 5,
			Hits: 0,
		},
		Ship{
			Name: "Battleship",
			Length: 4,
			Hits: 0,
		},
		Ship{
			Name: "Submarine",
			Length: 3,
			Hits: 0,
		},
		Ship{
			Name: "Cruiser",
			Length: 3,
			Hits: 0,
		},
		Ship{
			Name: "Destroyer",
			Length: 2,
			Hits: 0,
		},
	}

	return ships
}