package main

import (
)

func gameStart() *Game {
	game := Game{}.Init()

	//game.State = "setup"
	//
	//p1 := Player{
	//	game:  &game,
	//	Ships: makePlayerShips(),
	//	Board: &GameBoard{},
	//}
	////p1.Board.init()
	//game.Player1 = &p1
	//p2 := Player{
	//	game:  &game,
	//	Ships: makePlayerShips(),
	//	Board: &GameBoard{},
	//}
	////p2.Board.init()
	//game.Player2 = &p2
	//
	//var err error
	//me, err = game.Connect()
	//if err != nil {
	//	fmt.Println("Failed to connect:",err)
	//	return nil
	//}
	//
	//go me.Init()
	//
	//go func() {
	//	for {
	//		select {
	//		case msg := <-me.Output:
	//			if msg == ">" {
	//				//me.Input <- readLine(">")
	//				fmt.Print(msg)
	//			} else {
	//				fmt.Println(msg)
	//			}
	//		case <-time.After(time.Millisecond * 100):
	//			//do nothing
	//		}
	//	}
	//}()

	return &game
}
