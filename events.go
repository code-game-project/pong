/*
Pong v0.1
*/
package main

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
	Ball        Point `json:"ball"`
	PaddleLeft  Point `json:"paddle_left"`
	PaddleRight Point `json:"paddle_right"`
}

const EventScore cg.EventName = "score"

type EventScoreData struct {
	PlayerA int `json:"player_a"`
	PlayerB int `json:"player_b"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
