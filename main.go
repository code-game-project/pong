package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Bananenpro/log"
	"github.com/code-game-project/go-server/cg"
	"github.com/code-game-project/pong/pong"
	"github.com/spf13/pflag"
)

func main() {
	log.SetSeverity(log.TRACE)
	rand.Seed(time.Now().UnixNano())

	var port int
	pflag.IntVarP(&port, "port", "p", 0, "The network port of the game server.")
	pflag.Parse()

	if port == 0 {
		portStr, ok := os.LookupEnv("CG_PORT")
		if ok {
			port, _ = strconv.Atoi(portStr)
		}
	}

	if port == 0 {
		port = 80
	}

	server := cg.NewServer("pong", cg.ServerConfig{
		Port:                    port,
		CGEFilepath:             "pong.cge",
		MaxPlayersPerGame:       2,
		DeleteInactiveGameDelay: 15 * time.Minute,
		KickInactivePlayerDelay: 15 * time.Minute,
		DisplayName:             "Pong",
		Description:             "An implementation of Pong for CodeGame.",
		Version:                 "0.2",
		RepositoryURL:           "https://github.com/code-game-project/pong",
	})

	server.Run(func(cgGame *cg.Game) {
		pong.NewGame(cgGame).Run()
	})
}
