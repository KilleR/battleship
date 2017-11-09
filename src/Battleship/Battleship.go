package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
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
func main () {

	var (
		game Game
	)

	// start shell
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Battleship!")
	fmt.Println("------------")

Shell:
	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
		// strip LF
		text = strings.Replace(text, "\n", "", -1)
		// strip CR
		text = strings.Replace(text, "\r", "", -1)

		switch text {
		case "hi":
			fmt.Println("Hello!")
		case "start":
			fmt.Println("Starting...")
			game = gameStart()
		case "stop":
			fmt.Println("Stopping...")
		case "render":
			fmt.Println(render(game.Player1.Board))
		case "exit":
			fmt.Println("Bye!")
			break Shell
		default:
			fmt.Println("Unknown command:",text)
		}

	}
}
