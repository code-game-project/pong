package main

import (
	"math/rand"
	"time"

	"github.com/code-game-project/go-server/cg"
	"github.com/code-game-project/pong/pong"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	server := cg.NewServer(cg.ServerConfig{
		Port:              8080,
		CGEFilepath:       "pong.cge",
		MaxPlayersPerGame: 2,
	})

	server.Run(func(cgGame *cg.Game) {
		pong.NewGame(cgGame).Run()
	})
}
