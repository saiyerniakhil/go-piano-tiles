package main

import (
	"fmt"
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type Score struct {
	time  time.Duration
	moves int
}

const (
	AssetDir = "assets"
)

type GameState int

const (
	GameStatePlaying GameState = iota
	GameStateGameOver
)

type Game struct {
	game      *tl.Game
	board     *Board
	score     *Score
	stats     *tl.Text
	state     GameState
	startTime time.Time
	maxMoves  int
	won       bool
}

func GetBlackTilePos() int {
	return rand.Intn(boardWidth)
}

// TODO: EASY -15, MEDIUM: 20, HARD: 30
func GetMaxMoves() int {
	return totalMoves //TODO: Change later
}

func NewGame() *Game {
	startLevel := tl.NewBaseLevel(tl.Cell{})
	startScore := &Score{
		time: 0,
	}
	game := &Game{
		game:      tl.NewGame(),
		board:     NewBoard(startLevel),
		score:     startScore,
		stats:     tl.NewText(25, 0, "", tl.ColorWhite, tl.ColorBlack),
		startTime: time.Now(),
		maxMoves:  GetMaxMoves(),
	}
	game.board.SetGame(game)
	game.updateStatusText()
	return game
}

func (g *Game) updateStatusText() {
	elapsed := time.Since(g.startTime).Round(time.Millisecond)
	if g.state == GameStateGameOver {
		if g.won {
			statusText := fmt.Sprintf("Time Elapsed: %v\n Moves: %d/%d\n\n YOU WIN!", elapsed, g.score.moves, g.maxMoves)
			g.stats.SetText(statusText)
		} else {
			statusText := fmt.Sprintf("Time Elapsed:%v\n Moves: %d/%d\n\n YOU LOSE :()", elapsed, g.score.moves, g.maxMoves)
			g.stats.SetText(statusText)
		}
	} else {
		statusText := fmt.Sprintf("Time Elapsed: %v\nMoves: %d/%d", elapsed, g.score.moves, g.maxMoves)
		g.stats.SetText(statusText)
	}
}

func (g *Game) IncrementScore() {
	g.score.moves++
	if g.score.moves >= g.maxMoves {
		g.GameOver(true)
	}
	g.updateStatusText()
}

func (g *Game) GameOver(won bool) {
	g.state = GameStateGameOver
	g.won = won
	g.updateStatusText()
}

func (g *Game) RefreshScreen() {
	level := tl.NewBaseLevel(tl.Cell{})
	width := boardWidth*squareWidth + (boardWidth+2)*borderThickness
	height := boardHeight*squareHeight + (boardHeight+2)*borderThickness

	level.AddEntity(tl.NewRectangle(0, 0, width, height, tl.ColorGreen))

	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			level.AddEntity(&(*g.board.plan)[y][x])
		}
	}

	timerEntity := &TimerEntity{game: g}
	level.AddEntity(timerEntity)
	inputHandler := &InputHandler{game: g}
	level.AddEntity(inputHandler)

	g.board.level = level
	g.game.Screen().SetLevel(level)
}

func (g *Game) addChrome() {
	gapX := 4
	screen := g.game.Screen()
	screen.AddEntity(tl.NewText(offSetX+gapX, 0, "Piano Tiles", tl.ColorWhite, tl.ColorBlack))
	x := 2*offSetX + (boardWidth * squareWidth) + (boardWidth * borderThickness) + borderThickness
	instructions := fmt.Sprintf("%s \n\n\n\nHOW TO PLAY: \n\nPress 1, 2, 3, 4\nto match the black\ntile in the bottom\nrow.\n\nBoard: 4x5 grid\nTotal: %dmoves\n\n", loadAsset("header"), g.maxMoves)
	instructionsText := tl.NewEntityFromCanvas(x+gapX, offSetY+2, tl.CanvasFromString(instructions))
	screen.AddEntity(instructionsText)
	g.stats.SetPosition(x+gapX, offSetY+16)

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
	// g.game.DebugOn()
	g.board.populateBoard(level)

	timerEntity := &TimerEntity{game: g}
	level.AddEntity(timerEntity)
	inputHandler := &InputHandler{game: g}
	level.AddEntity(inputHandler)
	g.game.Screen().SetLevel(level)
	g.updateStatusText()
}
