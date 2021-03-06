package main

import (
	"fmt"
	"log"
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

func main() {
	var host *GameHost

	// start shell
	fmt.Println("Battleship!")
	fmt.Println("------------")

	host = &GameHost{}
	host.Init()
	defer host.Discord.Close()

	//Shell:
	// loop reading from discord
	for {
		msg := <-host.Discord.Recv
			log.Printf("Message from discord (ch: %s): %s\n", msg.ClientID, msg.Content)
			// check if the client is known
			gc := host.Clients.Get(msg.ClientID)
			if gc == nil {
				host.Discord.Send <- DiscordMessage{msg.ClientID, "I don't know you, give me a moment..."}
				host.Clients.Set(msg.ClientID, &GameClient{Host: host, ID: msg.ClientID})
				host.Discord.Send <- DiscordMessage{msg.ClientID, "Ok... what is your name?"}
			} else {
				gc.HandleDiscordMessage(host.Discord.Send, msg.Content)
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
