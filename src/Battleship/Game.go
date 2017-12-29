package main

import (
	"fmt"
	"log"
	"time"
)

type Game struct {
	State        string // holds game state 'setup', 'playing' ...
	PlayerTurn   *Player
	PlayersReady int
	Player1      *Player
	Player2      *Player
}

func (g *Game) Connect() (*Player, error) {
	// @TODO: possible race-condition here, multiple players could be allocated the same game, theoretically, but vanishingly small
	if g.Player1.ID == "" {
		return g.Player1, nil
	}
	if g.Player2.ID == "" {
		return g.Player2, nil
	}
	return nil, errors.New("no spaces left in the game")
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

func (g *Game) Fire(p *Player, coords [2]int) (bool, error) {
	o := g.GetOpponent(p)
	hitShip, err := o.GetShotAt(coords)
	if err != nil {
		return false, err
	}
	if hitShip == nil {
		return false, nil
	} else {
		// check to see if the ship was destroyed
		if hitShip.isDestroyed() {
			p.Output <- fmt.Sprintf("You destroyed their %s", hitShip.Name)
			// check to see if ALL their ships are destroyed
			allDestroyed := true
			for _,s := range o.Ships {
				if !s.isDestroyed() {
					allDestroyed = false
				}
			}
			if allDestroyed {
				g.State = "finished"
				// send the ship stats to both players
				p.Output <- "End of game details"
				o.Output <- "End of game details"
				p.Output <- fmt.Sprintf("%s's ships:", p.Name)
				o.Output <- fmt.Sprintf("%s's ships:", p.Name)
				for _,s := range p.Ships {
					if s.isDestroyed() {
						p.Output <- fmt.Sprintf("%s :- SUNK!", s.Name)
						o.Output <- fmt.Sprintf("%s :- SUNK!", s.Name)
					} else {
						p.Output <- fmt.Sprintf("%s :- %d hits", s.Name, s.Hits)
						o.Output <- fmt.Sprintf("%s :- %d hits", s.Name, s.Hits)
					}
				}
				p.Output <- fmt.Sprintf("%s's ships:", o.Name)
				o.Output <- fmt.Sprintf("%s's ships:", o.Name)
				for _,s := range o.Ships {
					if s.isDestroyed() {
						p.Output <- fmt.Sprintf("%s :- SUNK!", s.Name)
						o.Output <- fmt.Sprintf("%s :- SUNK!", s.Name)
					} else {
						p.Output <- fmt.Sprintf("%s :- %d hits", s.Name, s.Hits)
						o.Output <- fmt.Sprintf("%s :- %d hits", s.Name, s.Hits)
					}
				}

				p.Win()
				o.Lose()
			}
		}
		return true, nil
	}
}

func (g *Game) EndPlayerTurn(p *Player) {
	// change the active player turn
	g.PlayerTurn = g.GetOpponent(p)
	g.PlayerTurn.BeginTurn()
}

func (g *Game) MakeAIPlayer(p *Player) string {
	opp := g.GetOpponent(p)
	if opp.ID != "" {
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