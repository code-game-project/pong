package main

import (
	"math/rand"
	"time"

	"github.com/code-game-project/go-server/cg"
	"github.com/code-game-project/pong/pong"
	"github.com/spf13/pflag"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var port int
	pflag.IntVarP(&port, "port", "p", 80, "The network port of the game server.")
	pflag.Parse()

	server := cg.NewServer("pong", cg.ServerConfig{
		Port:              port,
		CGEFilepath:       "pong.cge",
		MaxPlayersPerGame: 2,
		DisplayName:       "Pong",
		Description:       "An implementation of Pong for CodeGame.",
		RepositoryURL:     "https://github.com/code-game-project/pong",
	})

	server.Run(func(cgGame *cg.Game) {
		pong.NewGame(cgGame).Run()
	})
}
