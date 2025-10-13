package main

import (
	"fmt"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type Score struct {
	time time.Duration
}

type Game struct {
	game  *tl.Game
	board *Board
	score *Score
	stats *tl.Text
}

func NewGame() *Game {
	startLevel := NewLevel(EASY)
	startScore := &Score{
		time: 0,
	}
	game := &Game{
		game:  tl.NewGame(),
		board: NewBoard(*startLevel),
		score: startScore,
		stats: tl.NewText(25, 0, "", tl.ColorWhite, tl.ColorBlack),
	}
	game.board.SetGame(game)
	game.updateStatusText()
	return game
}

func (g *Game) updateStatusText() {
	statusText := fmt.Sprintf("Time Elapsed: %s", g.score.time)
	g.stats.SetText(statusText)
}

func (g *Game) addChrome() {
	screen := g.game.Screen()
	screen.AddEntity(tl.NewText(offSetX, 0, "Piano Tiles", tl.ColorBlack, tl.ColorBlack))
	x := 2*offSetX + (boardWidth * squareWidth) + (boardWidth * borderThickness) + borderThickness
	rules := tl.NewEntityFromCanvas(x, offSetX, tl.CanvasFromString("Piano Tile"))
	screen.AddEntity(rules)
	screen.AddEntity(g.stats)
}

func (g *Game) Run() {
	g.addChrome()
	g.buildLevel()
	g.game.Start()
}

func (g *Game) buildLevel() {
	level := tl.NewBaseLevel(tl.Cell{})
	width := boardWidth*squareWidth + (boardWidth+1)*1
	height := boardHeight*squareHeight + (boardHeight+1)*1
	level.AddEntity(tl.NewRectangle(1, 1, width, height, tl.ColorGreen))
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			x := offSetX + 1 + (j * squareWidth) + j*1
			y := offSetY + 1 + (i * squareHeight) + i*1
			level.AddEntity(tl.NewRectangle(x, y, squareWidth, squareHeight, tl.ColorBlue))
		}
	}
	g.board.populateBoard(level)
	g.game.Screen().SetLevel(level)
	g.updateStatusText()
}
