package main

type GameClient struct {
	Player *Player
	ID string
}

func (c *GameClient) Init() {
	var err error
	c.Player, err = clientGame.Connect()
	if err != nil {

	}
}