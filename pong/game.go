package pong

import (
	"math"
	"math/rand"
	"time"

	"github.com/code-game-project/go-server/cg"
)

type player struct {
	cg    *cg.Player
	score int
	rect  Rectangle
	vY    float64
}

type Game struct {
	cg     *cg.Game
	config GameConfig

	width  float64
	height float64

	playerLeft  *player
	playerRight *player

	ball     Rectangle
	ballVelX float64
	ballVelY float64
}

func NewGame(cgGame *cg.Game, config GameConfig) *Game {
	game := &Game{
		cg:     cgGame,
		config: config,
		width:  650,
		height: 480,
		ball: Rectangle{
			Width:  15,
			Height: 15,
		},
	}
	cgGame.OnPlayerJoined = game.onPlayerJoined
	cgGame.OnPlayerLeft = game.onPlayerLeft
	return game
}

func (g *Game) handleCommand(player *player, cmd cg.Command) {
	if cmd.Name == MoveCmd {
		var data MoveCmdData
		err := cmd.UnmarshalData(&data)
		if err != nil {
			player.cg.Log.ErrorData(cmd, "failed to decode command data: %s", err)
			return
		}
		switch data.Direction {
		case DirectionNone:
			player.vY = 0
		case DirectionUp:
			player.vY = -8
		case DirectionDown:
			player.vY = 8
		}
	}
}

func (g *Game) update() {
	g.movePaddle(g.playerLeft)
	g.movePaddle(g.playerRight)

	g.moveBall()

	g.checkCollsions()
}

func (g *Game) movePaddle(player *player) {
	if player.rect.Y+player.rect.Height+player.vY > g.height {
		player.rect.Y = g.height - player.rect.Height
	} else if player.rect.Y+player.vY < 0 {
		player.rect.Y = 0
	} else {
		player.rect.Y += player.vY
	}
}

func (g *Game) moveBall() {
	g.ball.X += g.ballVelX

	if g.ball.Y+g.ballVelY+g.ball.Height > g.height {
		g.ball.Y = g.height - g.ball.Height
		g.ballVelY = -g.ballVelY
	} else if g.ball.Y+g.ballVelY < 0 {
		g.ball.Y = 0
		g.ballVelY = -g.ballVelY
	} else {
		g.ball.Y += g.ballVelY
	}
}

func (g *Game) checkCollsions() {
	if g.ball.X < -20 {
		g.goal(g.playerRight)
	} else if g.ball.X > g.width-g.ball.Width+20 {
		g.goal(g.playerLeft)
	}

	if g.ballVelX < 0 && g.ball.checkCollision(g.playerLeft.rect) {
		angle := ((g.ball.Y - g.playerLeft.rect.Y - g.ball.Height) / (g.playerLeft.rect.Height + g.ball.Height) * 0.75 * math.Pi) - 0.1875*math.Pi
		g.ballVelX = math.Cos(angle) * 14
		g.ballVelY = math.Sin(angle) * 14
	}
	if g.ballVelX > 0 && g.ball.checkCollision(g.playerRight.rect) {
		angle := ((g.ball.Y - g.playerRight.rect.Y - g.ball.Height) / (g.playerRight.rect.Height + g.ball.Height) * 0.75 * math.Pi) - 0.1875*math.Pi
		g.ballVelX = -math.Cos(angle) * 14
		g.ballVelY = math.Sin(angle) * 14
	}
}

func (g *Game) goal(player *player) {
	player.score++

	g.cg.Send(ScoreEvent, ScoreEventData{
		PlayerLeft:  g.playerLeft.score,
		PlayerRight: g.playerRight.score,
	})

	g.newBall()
	g.sendPositions()

	time.Sleep(time.Second)
}

func (g *Game) newBall() {
	g.playerLeft.rect.Y = g.height/2 - g.playerLeft.rect.Height/2
	g.playerRight.rect.Y = g.height/2 - g.playerRight.rect.Height/2

	g.ball.X = g.width/2 - g.ball.Width/2
	g.ball.Y = g.height/2 - g.ball.Height/2
	angle := rand.Float64()*0.5*math.Pi - 0.25*math.Pi
	g.ballVelX = math.Cos(angle) * 7
	g.ballVelY = math.Sin(angle) * 7

	if (g.playerLeft.score+g.playerRight.score)%2 == 0 {
		g.ballVelX = -g.ballVelX
	}
}

func (g *Game) onPlayerJoined(cgPlayer *cg.Player) {
	if g.playerLeft == nil {
		g.playerLeft = &player{
			rect: Rectangle{
				X:      15,
				Width:  15,
				Height: 75,
			},
			cg: cgPlayer,
		}
	} else {
		g.playerRight = &player{
			rect: Rectangle{
				X:      g.width - 15 - 15,
				Width:  15,
				Height: 75,
			},
			cg: cgPlayer,
		}
		g.start()
	}
}

func (g *Game) onPlayerLeft(player *cg.Player) {
	g.cg.Close()
}

func (g *Game) sendPositions() {
	g.cg.Send(PositionsEvent, PositionsEventData{
		Ball:        g.ball,
		PaddleLeft:  g.playerLeft.rect,
		PaddleRight: g.playerRight.rect,
	})
}

func (g *Game) start() {
	g.playerLeft.cg.Send(StartEvent, StartEventData{
		Side: SideLeft,
	})
	g.playerRight.cg.Send(StartEvent, StartEventData{
		Side: SideRight,
	})
	g.newBall()
}

func (g *Game) pollCommands() {
	for g.cg.Running() {
		cmd, ok := g.cg.NextCommand()
		if !ok {
			return
		}
		player := g.playerLeft
		if cmd.Origin == g.playerRight.cg {
			player = g.playerRight
		}
		g.handleCommand(player, cmd.Cmd)
	}
}

func (g *Game) Run() {
	for g.cg.Running() {
		start := time.Now().UnixMilli()

		g.pollCommands()

		if g.playerLeft != nil && g.playerRight != nil {
			g.update()
			g.sendPositions()
		}

		end := time.Now().UnixMilli()
		time.Sleep(time.Duration((1000/30)-(end-start)) * time.Millisecond)
	}
}

func (r Rectangle) checkCollision(other Rectangle) bool {
	return r.X+r.Width >= other.X && r.X <= other.X+other.Width && r.Y+r.Height >= other.Y && r.Y <= other.Y+other.Height
}
