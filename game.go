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
	width := boardWidth*squareWidth + (boardWidth+2)*borderThickness
	height := boardHeight*squareHeight + (boardHeight+2)*borderThickness
	// background
	level.AddEntity(tl.NewRectangle(0, 0, width, height, tl.ColorGreen))
	fmt.Printf("[buildLevel] offset values - X (%d) Y (%d)\n", offSetX, offSetY)
	fmt.Printf("[buildLevel] board values - width[j] (%d) height[i] (%d)\n", boardWidth, boardHeight)
	// board rendering
	g.board.populateBoard(level)
	g.game.Screen().SetLevel(level)
	g.updateStatusText()
}
