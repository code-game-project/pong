/*
An implementation of Pong for CodeGame.
*/
package pong

import "github.com/code-game-project/go-server/cg"

type Side string

const (
	SideLeft  Side = "left"
	SideRight Side = "right"
)

// Is sent once the game starts.
const EventStart cg.EventName = "start"

type EventStartData struct {
	// The side of the player.
	Side Side `json:"side"`
}

type Direction string

const (
	DirectionNone Direction = "none"
	DirectionUp   Direction = "up"
	DirectionDown Direction = "down"
)

// Clients send the move event to the server to signal a direction change of their paddle.
const EventMove cg.EventName = "move"

type EventMoveData struct {
	// The direction of the paddle movement.
	Direction Direction `json:"direction"`
}

// The positions event is sent to all players to tell them the positions of all objects in the game.
const EventPositions cg.EventName = "positions"

type EventPositionsData struct {
	// The position of the ball.
	Ball Rectangle `json:"ball"`
	// The position of left paddle.
	PaddleLeft Rectangle `json:"paddle_left"`
	// The position of right paddle.
	PaddleRight Rectangle `json:"paddle_right"`
}

// The score event is sent to all players when a player scores a point.
const EventScore cg.EventName = "score"

type EventScoreData struct {
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
