package main

import (
	"fmt"
	"strings"
	"log"
	"time"
)

type GameClient struct {
	Player *Player
	Host   *GameHost
	ID     string
	Name   string
}

func (c *GameClient) HandleGameMessages(out chan DiscordMessage) chan bool {
	terminateChannel := make(chan bool)
	// listen for output from Player.Output
	out <- DiscordMessage{c.ID, "Listening for messages from the game"}
	go func() {
		for {
			select {
			case <-terminateChannel:
				return
			case msg := <-c.Player.Output:
				out <- DiscordMessage{c.ID, msg}
			//default:
			//	time.Sleep(time.Millisecond * 100)
				// do nothing
			}
		}
	}()

	return terminateChannel
}

func (c *GameClient) StopHandlingGameMessages(term chan bool) {
	term <- true
	close(term)
}

// Handle incoming messages for this client
func (c *GameClient) HandleDiscordMessage(out chan DiscordMessage, text string) {

	var helpText = `Commands:
help, ? - Prints this message.
rename - Forgets your nickname, and asks for a new one.
start, new, new game - Starts a game of Battleship!`
	switch {
	case c.Name == "": // name is blank, we must be naming the player
		c.Name = text
		out <- DiscordMessage{c.ID, fmt.Sprintf("Hi %s! You can change what you're called if you aren't playing a game by saying 'rename'.", c.Name)}
	case c.Player != nil:
		log.Println("Player is in game, processing to game:", text)
		switch text {
		case "exit":
			out <- DiscordMessage{c.ID, "Sorry, you can't just quit a game. You might hurt someone else's feelings"}
		default:
			c.Player.Input <- text
		}
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
			p := c.Host.ConnectToGame()
			if p == nil {
				out <- DiscordMessage{c.ID, "Sorry, I couldn't find a game for you. Try again soon!"}
			} else {
				c.Player = p
				c.HandleGameMessages(out)
				go c.Player.Init(c.Name)
				out <- DiscordMessage{c.ID, fmt.Sprintf("Success, you have connected to game: %s", p.game.ID)}
			}
		default:
			out <- DiscordMessage{c.ID, fmt.Sprintf("I'm sorry, %s, I don't understand what you mean by `%s`.\r\nYou can say `help` to see what you can ask me to do.", c.Name, text)}
		}
	}
}
