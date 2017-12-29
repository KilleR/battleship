package main

type GameHost struct {
	Discord *Discord
	clients []*GameClient
}

func (gh *GameHost) Init() {
	gh.Discord = &Discord{}

	gh.Discord.Connect()
}