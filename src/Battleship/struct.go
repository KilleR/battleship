package main

import (
	"time"
	"math/rand"
	"log"
)

type Game struct {
	State        string // holds game state 'setup', 'playing' ...
	PlayerTurn   *Player
	PlayersReady int
	Player1      *Player
	Player2      *Player
}

func (g *Game) Connect() *Player {
	return g.Player1
}

func (g *Game) ReadyCheck(p *Player) {
	g.PlayersReady++
	// if both players are ready, select an active player
	if g.PlayersReady >= 2 {
		if rand.Intn(2) == 0 {
			g.PlayerTurn = g.Player1
		} else {
			g.PlayerTurn = g.Player2
		}
		// broadcast to both players who is going first
		if g.PlayerTurn == g.Player1 {
			g.Player1.Output <- "You are going first"
			g.Player2.Output <- "You are going second"
		} else {
			g.Player2.Output <- "You are going first"
			g.Player1.Output <- "You are going second"
		}
		g.PlayerTurn.BeginTurn()
		g.State = "playing"
	}
	// if one player is ready, but no player is on the other side, prompt to create an AI
	if g.GetOpponent(p).ID == "" {
		p.Output <- "Nobody is connected, make a (really dumb) AI oponent with 'aime'"
	} else {
		g.GetOpponent(p).Output <- "Your opponent is READY"
	}
}

func (g *Game) EndPlayerTurn(p *Player) {
	// change the active player turn
	g.PlayerTurn = g.GetOpponent(p)
	g.PlayerTurn.BeginTurn()
}

func (g *Game) MakeAIPlayer(p *Player) string {
	opp := g.GetOpponent(p)
	if (opp.ID != "") {
		return "There is already a player connected"
	} else {
		opp.IsAI = true
		opp.ID = "BOT"

		//start consuming the AI's output
		go func() {
			for {
				select {
				case msg := <-opp.Output:
					log.Println(msg)
				case <-time.After(time.Millisecond * 100):
					// do nothing
				}
			}
		}()

		opp.Init()
		p.Output <- "AI ready"
		//g.ReadyCheck(g.Player2)
		return "Made your opponent an AI"
	}
}

func (g *Game) GetOpponent(p *Player) *Player {
	if p == g.Player1 {
		return g.Player2
	} else {
		return g.Player1
	}
}

type GameBoard struct {
	Grid [10][10]*Ship
}

func (gb GameBoard) GetShipAt(coords [2]int) *Ship {
	return gb.Grid[coords[0]][coords[1]]
}

type Ship struct {
	Name        string
	Length      int
	Hits        int    // number of hits against it
	Location    string // top left coordinate
	Orientation string // 'h' or 'v' for horizontal or vertical
}

func (s *Ship) isDestroyed() bool {
	return s.Hits >= s.Length
}
