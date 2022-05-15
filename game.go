package main

import (
	"time"

	"github.com/code-game-project/go-server/cg"
)

type game struct {
	cg *cg.Game
}

func newGame(cgGame *cg.Game) *game {
	game := &game{
		cg: cgGame,
	}
	cgGame.OnPlayerJoined = game.onPlayerJoined
	cgGame.OnPlayerLeft = game.onPlayerLeft
	return game
}

func (g *game) handleEvent(event cg.Event, player *cg.Player) {

}

func (g *game) update() {

}

func (g *game) onPlayerJoined(player *cg.Player) {

}

func (g *game) onPlayerLeft(player *cg.Player) {
	g.cg.Close()
}

func (g *game) sendPositions() {
	g.cg.Send("server", EventPositions, EventPositionsData{
		// TODO
	})
}

func (g *game) pollEvents() {
	for g.cg.Running() {
		select {
		case event := <-g.cg.Events:
			g.handleEvent(event.Event, event.Player)
		default:
			return
		}
	}
}

func (g *game) Run() {
	for g.cg.Running() {
		start := time.Now().UnixMilli()

		g.pollEvents()
		g.update()

		g.sendPositions()

		end := time.Now().UnixMilli()
		time.Sleep(time.Duration((1000/30)-(end-start)) * time.Millisecond)
	}
}
