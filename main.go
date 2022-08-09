package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/code-game-project/go-server/cg"
	"github.com/spf13/pflag"

	"github.com/code-game-project/pong/pong"
)

func main() {
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
		port = 8080
	}

	server := cg.NewServer("pong", cg.ServerConfig{
		DisplayName:             "Pong",
		MaxPlayersPerGame:       2,
		Version:                 "0.3",
		Description:             "An implementation of Pong for CodeGame.",
		RepositoryURL:           "https://github.com/code-game-project/pong",
		Port:                    port,
		CGEFilepath:             "events.cge",
		DeleteInactiveGameDelay: 5 * time.Minute,
		KickInactivePlayerDelay: 15 * time.Minute,
	})

	server.Run(func(cgGame *cg.Game, config json.RawMessage) {
		var gameConfig pong.GameConfig
		err := json.Unmarshal(config, &gameConfig)
		if err == nil {
			cgGame.SetConfig(gameConfig)
		} else {
			cgGame.Log.Error("Failed to unmarshal game config: %s", err)
		}

		pong.NewGame(cgGame, gameConfig).Run()
	})
}
