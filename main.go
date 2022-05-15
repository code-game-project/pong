package main

import "github.com/code-game-project/go-server/cg"

func main() {
	server := cg.NewServer(cg.ServerConfig{
		Port:              8080,
		CGEFilepath:       "pong.cge",
		MaxPlayersPerGame: 2,
	})

	server.Run(func(cgGame *cg.Game) {
		newGame(cgGame).Run()
	})
}
