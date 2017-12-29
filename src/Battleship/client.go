package main

import (
	"fmt"
	"strings"
)

type GameClient struct {
	Player *Player
	Host   *GameHost
	ID     string
	Name   string
}

// Handle incoming messages for this client
func (c *GameClient) HandleDiscordMessage(out chan DiscordMessage, text string) {
	var err error
	var helpText = `Commands:
help, ? - Prints this message.
rename - Forgets your nickname, and asks for a new one.`
	switch {
	case c.Name == "": // name is blank, we must be naming the player
		c.Name = text
		out <- DiscordMessage{c.ID, fmt.Sprintf("Hi %s! You can change what you're called if you aren't playnig a game by saying 'rename'.", c.Name)}
	default:
		text = strings.ToLower(text)
		switch text {
		case "help", "?":
			out <- DiscordMessage{c.ID, helpText}
		case "rename":
			c.Name = ""
			out <- DiscordMessage{c.ID, "OK! What do you want to be called instead?"}
		case "start", "new", "new game":
			out <- DiscordMessage{c.ID, "Searching for a game..."}
			g:=c.Host.Games.New()
			if g == nil {
				out <- DiscordMessage{c.ID, "Sorry, I couldn't find a game for you. Try again soon!"}
			} else {
				fmt.Println(c)
				fmt.Println(g)
				c.Player, err = g.Connect()
				if err != nil {
					out <- DiscordMessage{c.ID, fmt.Sprintf("Sorry, something went wrong connecting to the game: %s", err.Error())}
				} else {
					out <- DiscordMessage{c.ID, fmt.Sprintf("Success, you have connected to game: %s", g.ID)}
				}
			}

		default:
			out <- DiscordMessage{c.ID, fmt.Sprintf("I'm sorry, %s, I don't understand what you mean by `%s`.\r\nYou can say `help` to see what you can ask me to do.", c.Name, text)}
		}
	}
}
