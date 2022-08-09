package pong

import "github.com/code-game-project/go-server/cg"

type Direction string

const (
	DirectionNone Direction = "none"
	DirectionUp   Direction = "up"
	DirectionDown Direction = "down"
)

// Clients send the move command to the server to signal a direction change of their paddle.
const MoveCmd cg.CommandName = "move"

type MoveCmdData struct {
	// The direction of the paddle movement.
	Direction Direction `json:"direction"`
}

type Side string

const (
	SideLeft  Side = "left"
	SideRight Side = "right"
)

// Is sent once the game starts.
const StartEvent cg.EventName = "start"

type StartEventData struct {
	// The side of the player.
	Side Side `json:"side"`
}

// The positions event is sent to all players to tell them the positions of all objects in the game.
const PositionsEvent cg.EventName = "positions"

type PositionsEventData struct {
	// The position of the ball.
	Ball Rectangle `json:"ball"`
	// The position of left paddle.
	PaddleLeft Rectangle `json:"paddle_left"`
	// The position of right paddle.
	PaddleRight Rectangle `json:"paddle_right"`
}

// The score event is sent to all players when a player scores a point.
const ScoreEvent cg.EventName = "score"

type ScoreEventData struct {
	// The score of the left player.
	PlayerLeft int `json:"player_left"`
	// The score of the right player.
	PlayerRight int `json:"player_right"`
}

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type GameConfig struct {
}
