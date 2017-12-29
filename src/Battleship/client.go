package main

type GameClient struct {
	Player *Player
	ID string
	Name string
}

func (c *GameClient) Init() {
	var err error
	c.Player, err = clientGame.Connect()
	if err != nil {

	}
}

// Handle incoming messages for this client
func (c *GameClient) HandleMessage(text string) {
	switch text {

	}
}