package main

import (
	"fmt"
	"time"
)

func gameStart() Game {
	game := Game{}

	game.State = "setup"

	p1 := Player{
		Name:  "Bill",
		game:  &game,
		Ships: makePlayerShips(),
		Board: &GameBoard{},
	}
	//p1.Board.init()
	game.Player1 = &p1
	p2 := Player{
		Name:  "Bill",
		game:  &game,
		Ships: makePlayerShips(),
		Board: &GameBoard{},
	}
	//p2.Board.init()
	game.Player2 = &p2

	me = game.Player1

	go me.Init()

	go func() {
		for {
			select {
			case msg := <-me.Output:
				if msg == ">" {
					//me.Input <- readLine(">")
					fmt.Print(msg)
				} else {
					fmt.Println(msg)
				}
			case <-time.After(time.Millisecond * 100):
				//do nothing
			}
		}
	}()

	//fmt.Println("This is your game board:")
	//fmt.Println(render(me))
	//fmt.Println("Start placing ships!")

	//game.State = "playing"

	return game
}
