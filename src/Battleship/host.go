package main

import (
	"sync"
	"log"
)

type GameHost struct {
	Discord *Discord
	Clients GameHostClients
	Games   GameHostGames
}

func (gh *GameHost) Init() {
	gh.Discord = &Discord{}

	gh.Discord.Connect()

	gh.Clients = gh.Clients.Init()
	gh.Games = gh.Games.Init()
}

func (gh *GameHost) ConnectToGame() *Player {
	p := gh.Games.Connect()
	return p
}

// Data store for clients
type GameHostClients struct {
	sync.RWMutex
	data map[string]*GameClient
}

func (ghc GameHostClients) Init() GameHostClients {
	ghc.data = make(map[string]*GameClient)

	return ghc
}

func (ghc GameHostClients) Get(key string) *GameClient {
	ghc.RLock()
	defer ghc.RUnlock()

	gc := ghc.data[key] // doesn't need `gc,ok :=` idiom, will return <nil> if value is not set
	return gc
}

func (ghc GameHostClients) Set(key string, value *GameClient) {
	ghc.Lock()
	defer ghc.Unlock()

	ghc.data[key] = value
}

// data store for running games
type GameHostGames struct {
	sync.RWMutex
	data map[string]*Game
}

func (ghg GameHostGames) Init() GameHostGames {
	ghg.data = make(map[string]*Game)

	return ghg
}

func (ghr GameHostGames) Get(key string) *Game {
	ghr.RLock()
	defer ghr.RUnlock()

	gc := ghr.data[key] // doesn't need `gc,ok :=` idiom, will return <nil> if value is not set
	return gc
}

func (ghg GameHostGames) Set(key string, value *Game) {
	ghg.Lock()
	defer ghg.Unlock()

	ghg.data[key] = value
}

func (ghg GameHostGames) UnSet(key string, value *Game) {
	ghg.Lock()
	defer ghg.Unlock()

	delete(ghg.data, key)
}

func (ghg GameHostGames) New() *Game {
	// Loop a max of 10 times trying to create a new game ID until a unique one is geenrated
	// This should realistically never fail, the seed is large enough
	for i := 0; i < 10; i++ {
		key := RandString(16)
		host := ghg.Get(key)
		if host == nil {
			g := Game{ID:key}.Init()
			ghg.Set(key, &g)
			return &g
		}
	}
	return nil
}

func (ghg GameHostGames) ConnectToActive() *Player {
	// go through all of the active games and try to connect
	ghg.RLock()
	defer ghg.RUnlock()

	for _,v := range ghg.data {
		p, err := v.Connect()
		if err == nil {
			return p
		}
	}

	return nil
}

func (ghg GameHostGames) Connect() *Player {
	// try to connect to an active game
	p := ghg.ConnectToActive()
	if p != nil {
		return p
	}

	log.Println("No game, making a new one")
	// if there are no valid active games, make a new one
	g := ghg.New()
	if g != nil {
		// connect to the new game
		p, err := g.Connect()
		if err == nil {
			return p
		}
	}

	return nil
}




