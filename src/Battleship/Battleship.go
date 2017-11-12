package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"regexp"
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
 	me *Player
 	game Game
 	coordRex *regexp.Regexp
 )

func main () {
	coordRex = regexp.MustCompile(`([a-jA-J])([0-9])`)

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
			isCoord := false
			if len(text) == 2 {
				// on an exactly 2 character input, check if it's a coordinate
				coords := coordRex.FindAllStringSubmatch(text, -1)
				if len(coords) > 0 {
					isCoord = true
					// since it's a coord, I need to know what I'm doing with it
					if game.State == "setup" {
						// if we're in setup, I'm probalby placing ships
						read := readLine("make an input")
						fmt.Println("Read:",read)
					}
				}
			}
			// if it was not a coord, print the unknown command text
			if !isCoord {
				fmt.Println("Unknown command:",text)
			}
		}

	}
}
