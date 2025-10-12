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
}

// Board is a collection of squares
type Board struct {
	plan  [][]Square
	nRows int
	nCols int
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
func NewSquare(boardX, boardY, value int) *Square {
	sq := &Square{
		value: value,
		content: tl.Cell{
			Bg: getTileColor(value),
		},
		entity: tl.NewEntity(1, 1, 20, 20),
	}
	//TODO: Set entity position
	//sq.entity.SetPosition()
	return sq
}

const (
	offSetX = 2
	offSetY
	squareOffsetX   = 1
	squareOffsetY   = 1
	borderThickness = 1
	squareWidth     = 2
	squareHeight    = 2
)

// getPosition returns the position of a given square
func (sq *Square) getPosition() (int, int) {
	x := offSetX + borderThickness + (sq.boardX * squareWidth) + (sq.boardX * squareOffsetX)
	y := offSetY + borderThickness + (sq.boardY * squareHeight) + (sq.boardY * squareOffsetY)
	return x, y
}

func (sq *Square) Tick(event tl.Event) {}

// Draw draws the square on a given screen
func (sq *Square) Draw(screen *tl.Screen) {
	x, y := sq.getPosition()
	for i := 0; i < 5; i++ {
		screen.RenderCell(x+i, y, &sq.content)
	}
}

func NewBoard(level Level) *Board {
	rows, cols := level.GetBoardDimensions()
	plan := make([][]Square, rows)
	for i := 0; i < rows; i++ {
		plan[i] = make([]Square, cols)
	}
	return &Board{
		plan:  plan,
		nRows: rows,
		nCols: cols,
	}
}

func (b *Board) populateBoard(level tl.Level) {
	for y := 0; y < b.nRows; y++ {
		for x := 0; x < b.nCols; x++ {
			//TODO: Get random number between 1...4
			rc := NewSquare(x, y, 0)
			b.plan[y][x] = *rc
			level.AddEntity(rc)
		}
	}
}
