package main

import "sync"

type GameHost struct {
	Discord *Discord
	Clients GameHostClients
	Games GameHostGames
}

func (gh *GameHost) Init() {
	gh.Discord = &Discord{}

	gh.Discord.Connect()

	gh.Clients = gh.Clients.New()
}

// Data store for clients
type GameHostClients struct {
	sync.RWMutex
	data map[string]*GameClient
}

func (ghc GameHostClients) New() GameHostClients{
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

func (ghg GameHostGames) New() GameHostGames{
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