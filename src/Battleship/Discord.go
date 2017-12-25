package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"fmt"
)

func discordConnect() *discordgo.Session{
	discord, err := discordgo.New("Bot Mzk0OTA5MDU0MzE0NzQxNzcz.DSLkuA.Eake1xS39S_ZBFCeh_5An_NEXlY")
	if err != nil {
		log.Fatalln("Failed to start Discord BOT:", err)
	}

	err = discord.Open()
	if err != nil {
		log.Fatalln("Failed to open Discord:", err)
	}

	discord.AddHandler(messageCreate)

	chans, err := discord.UserChannels()
	if err != nil {
		log.Println("Error getting channels:", err)
	}
	fmt.Println("Channels:")
	for _, v := range chans {
		fmt.Println(v.Name)
	}

	dChan, err := discord.Channel("265899186371821577")
	if err != nil {
		log.Println("Error getting random channel:", err)
	} else {
		fmt.Println("Channel:", dChan.Name, dChan.GuildID)
	}

	guilds, err := discord.UserGuilds(10, "", "")
	if err != nil {
		log.Println("Error getting guilds:", err)
	}
	fmt.Println("Guilds:")
	for _, v := range guilds {
		fmt.Println(v.Name, v.ID)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Discord bot is now running.")

	return discord

	err = discord.Close()
	if err != nil {
		log.Fatalln("Could not close Discord:", err)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Println("Failed to get message channel ID:", err)
		return
	}
	fmt.Printf("Message received from %s (channel: %s %v %s): %s\n", m.Author.Username, channel.Name, channel.Type, m.ChannelID, m.Content)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

contentSwitch:
	switch m.Content {
	case "ping", "Ping":
		// If the message is "ping" reply with "Pong!"
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	case "pong", "Pong":
		// If the message is "pong" reply with "Ping!"
		s.ChannelMessageSend(m.ChannelID, "Ping!")
		return
	case "<@394909054314741773>":
	default:
		// look for mentions
		isMentioned := false
		for _, v := range m.Mentions {
			if v.ID == "394909054314741773" {
				isMentioned = true
			}
		}
		if isMentioned {
			// reply to mentions with a DM
			fmt.Println("Got a Mention")
			dmChannel, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Println("Failed to create DM channel:", err)
			} else {
				s.ChannelMessageSend(dmChannel.ID, "Hai, try talking to me, rather than about me.")
			}
			break contentSwitch
		}

		// look and see if it's a DM to the bot
		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			log.Println("Failed to get message channel ID:", err)
			return
		}
		if channel.Type == discordgo.ChannelTypeDM {
			s.ChannelMessageSend(m.ChannelID, "Hi! Sorry, I'm not able to understand you yet")
			break contentSwitch
		}
	}
}
