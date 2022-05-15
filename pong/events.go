/*
Pong v0.1
*/
package pong

import "github.com/code-game-project/go-server/cg"

type Direction string

const (
	DirectionNone Direction = "none"
	DirectionUp   Direction = "up"
	DirectionDown Direction = "down"
)

const EventMove cg.EventName = "move"

type EventMoveData struct {
	Direction Direction `json:"direction"`
}

const EventPositions cg.EventName = "positions"

type EventPositionsData struct {
	Ball        Rectangle `json:"ball"`
	PaddleLeft  Rectangle `json:"paddle_left"`
	PaddleRight Rectangle `json:"paddle_right"`
}

const EventScore cg.EventName = "score"

type EventScoreData struct {
	PlayerLeft  int `json:"player_left"`
	PlayerRight int `json:"player_right"`
}

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
