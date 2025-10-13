package main

import tl "github.com/JoelOtter/termloop"

/***
 A 5 * 4 board looks like this
	[0] [1] [0] [0]
	[0] [1] [0] [0]
	[0] [0] [1] [0]
	[0] [0] [1] [0]
	[0] [1] [0] [0]
*/

// Square denotes a tile, it is either black or white (denoted by 0 / 1)
type Square struct {
	boardX  int
	boardY  int
	value   int
	content tl.Cell
	entity  *tl.Entity
	game    *Game
}

// Board is a collection of squares
type Board struct {
	plan  [][]Square
	level Level
	game  *Game
}

// getTileColor gives tile color
func getTileColor(value int) tl.Attr {
	switch value {
	case 0:
		return tl.ColorWhite
	default:
		return tl.ColorCyan
	}
}

// NewSquare returns a square at a given position
func NewSquare(boardX, boardY, value int, game *Game) *Square {
	sq := &Square{
		value: value,
		content: tl.Cell{
			Bg: getTileColor(value),
		},
		entity: tl.NewEntity(1, 1, 20, 20),
		game:   game,
	}
	//TODO: Set entity position
	sq.entity.SetPosition(sq.getPosition())
	return sq
}

func (b *Board) SetGame(g *Game) {
	b.game = g
}

func (sq *Square) SetGame(g *Game) {
	sq.game = g
}

const (
	offSetX = 2
	offSetY
	squareOffsetX   = 1
	squareOffsetY   = 1
	borderThickness = 1
	squareWidth     = 2
	squareHeight    = 2
	boardWidth      = 8
	boardHeight     = 8
)

// getPosition returns the position of a given square
func (sq *Square) getPosition() (int, int) {
	x := offSetX + borderThickness + (sq.boardX * squareWidth) + (sq.boardX * squareOffsetX)
	y := offSetY + borderThickness + (sq.boardY * squareHeight) + (sq.boardY * squareOffsetY)
	return x, y
}

func (sq *Square) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.MouseLeft:
			if sq.value == 1 {
				//TODO: success
			}
		default:
		}
	}
}

// Draw draws the square on a given screen
func (sq *Square) Draw(screen *tl.Screen) {
	x, y := sq.getPosition()
	for i := 0; i < 5; i++ {
		screen.RenderCell(x+i, y, &sq.content)
	}
}

func NewBoard(level Level) *Board {
	rows, cols := level.dimensions.nRows, level.dimensions.nCols
	plan := make([][]Square, rows)
	for i := 0; i < rows; i++ {
		plan[i] = make([]Square, cols)
	}
	return &Board{
		plan:  plan,
		level: level,
	}
}

func (b *Board) populateBoard(level tl.Level) {
	nRows, nCols := b.level.dimensions.nRows, b.level.dimensions.nCols
	for y := 0; y < nRows; y++ {
		for x := 0; x < nCols; x++ {
			var rc *Square
			if b.level.GetBlackTile() == x {
				rc = NewSquare(x, y, 1, b.game)
			} else {
				rc = NewSquare(x, y, 0, b.game)
			}
			b.plan[y][x] = *rc
			level.AddEntity(rc)
		}
	}
}
