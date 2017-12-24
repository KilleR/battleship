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
var (
	me   *Player
	game Game
)

func main() {
	// start shell
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Battleship!")
	fmt.Println("------------")

Shell:
	for {
		// Check for game end
		if game.State == "playing" && me.ShipsRemaining() <= 0 {
			fmt.Println("You Lose!")
			break Shell
		}

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
			if game.State == "" {
				fmt.Println("Game has not started, there is no board to render")
			} else {
				fmt.Println(render(game.Player1))
			}
		case "fired":
			if game.State == "" {
				fmt.Println("Game has not started, there is no board to render")
			} else {
				fmt.Println(renderFired(game.Player1))
			}
		case "ships":
			fmt.Println(me.ShipStatus())
		case "exit":
			fmt.Println("Bye!")
			break Shell
		default:
			isCoord := false
			if len(text) == 2 {
				// on an exactly 2 character input, check if it's a coordinate
				coords, err := stringToCoords(text)
				if err != nil {
					fmt.Println("Failed to parse coordinate:", err)
				} else {
					fmt.Println("coords entered:", coords)
					isCoord = true
					// since it's a coord, I need to know what I'm doing with it
					switch game.State {
					case "setup":
						// if we're in setup, I'm probalby placing ships
						read := readLine("make an input")
						fmt.Println("Read:", read)
					case "playing":
						fmt.Println("Shot fired at: " + text)
						fireResult, err := me.FireAt(coords)
						if err != nil {
							fmt.Println("Firing error:",err)
							continue Shell
						}
						fmt.Println(fireResult)
					}
				}
			}
			// if it was not a coord, print the unknown command text
			if !isCoord {
				fmt.Println("Unknown command:", text)
			}
		}

	}
}
