package main

type GameHost struct {
	Discord *Discord
}

func (gh *GameHost) Init() {
	gh.Discord = &Discord{}

	gh.Discord.Connect()
}