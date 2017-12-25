package main

import "github.com/bwmarrin/discordgo"

type GameHost struct {
	Discord *discordgo.Session
}

func (gh *GameHost) Init() {
	gh.Discord = discordConnect()
}