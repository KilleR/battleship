package main

import (
	"fmt"
	"time"
)

/*
A simulator for the Battleship game

Final aim:
2 players will connect from different locatinos
Each will be created a game board, and appropriate ships to place on the board

Players will place their ships on their board
No ships may occupy the same space(s)

Once both players are done placing ships, play begins with a random player

Each player enters coordinates (format [A-J][0-9] to fire at the other player's board
Hits are announced
A ship is sunk when all of its locations are hit.
Destroyed ships are announced (optional)
[Extra shots for unsunk ships] (optional)

Once a player has sunk all the op
 */
var (
	me   *Player
	clientGame *Game
	host *GameHost
)

func main() {
	// start shell
	fmt.Println("Battleship!")
	fmt.Println("------------")
	fmt.Print(">")

	host = &GameHost{}
	host.Init()
	defer host.Discord.Close()

//Shell:
	// loop reading from discord
	for {
		select {
		case text := <-host.Discord.Output:
			fmt.Println("Message from discord:",text)
		case <-time.After(time.Millisecond * 100):
			// do nothing
		}
	}

	// loop reading from STDIN
	//for {
	//	// Check for game end
	//	if clientGame != nil && clientGame.State == "playing" && me.ShipsRemaining() <= 0 {
	//		fmt.Println("You Lose!")
	//		break Shell
	//	}
	//
	//	text := readLine("")
	//
	//	switch text {
	//	case "start":
	//		fmt.Println("Starting...")
	//		clientGame = gameStart()
	//	case "exit":
	//		fmt.Println("Bye!")
	//		break Shell
	//	default:
	//		if clientGame == nil || clientGame.State == "" {
	//			fmt.Println("Unknown command:", text)
	//			fmt.Print(">")
	//		} else {
	//			me.Input <- text
	//		}
	//	}
	//
	//}
}
