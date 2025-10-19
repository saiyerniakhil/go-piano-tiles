package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

// Board is a collection of squares
type Board struct {
	plan  [][]Tile
	level Level
	game  *Game
}

func (b *Board) SetGame(g *Game) {
	b.game = g
}

// Move moves the board
func (b *Board) Move(valid bool) {
	if valid {
		//TODO: From b.plan remove the row -> plan[boardHeigh - 1]
		//add new entities
	} else {
		//stop the timer - "Game Over"
	}
}

func (b *Board) NewRow() []Tile {
	row := make([]Tile, boardWidth)
	for i := range boardWidth {
		bc := BoardCoords{x: i, y: boardHeight - 1}
		// place this row at the top
		newPosX, newPosY := getPosition(bc)
		row = append(row, *NewTile(newPosX, newPosY, squareWidth, squareHeight, tl.ColorWhite, b.game, bc))
	}
	return row
}

func (b *Board) IsValidMove(currKey int) bool {
	// x = board width, col
	// y = board height, row
	// get the last row
	if b.plan[boardHeight-1][currKey].Color() == tl.ColorBlack {
		return true
	} else {
		return false
	}
}

const (
	offSetX         = 1
	offSetY         = 1
	squareOffsetX   = 1
	squareOffsetY   = 2
	borderThickness = 1
	squareWidth     = 3 //DONE
	squareHeight    = 6 //DONE
	boardWidth      = 4 //Relative to Square width
	boardHeight     = 5 //Relative
)

// getPosition returns the position when the boardPos of square is given
// for a square at (1, 2) position, 2nd row and 1st col returns the absolute pos
func getPosition(bc BoardCoords) (int, int) {
	x := offSetX + borderThickness + (bc.x * squareWidth) + (bc.x * borderThickness)
	y := offSetY + borderThickness + (bc.y * squareHeight) + (bc.y * borderThickness)
	return x, y
}

func NewBoard(level Level) *Board {
	rows, cols := boardHeight, boardWidth
	plan := make([][]Tile, rows)
	for i := 0; i < rows; i++ {
		plan[i] = make([]Tile, cols)
	}
	return &Board{
		plan:  plan,
		level: level,
	}
}

func (b *Board) populateBoard(level tl.Level) {
	nRows, nCols := boardHeight, boardWidth
	fmt.Printf("rows: %d cols: %d\n", nRows, nCols)
	for y := 0; y < nRows; y++ {
		blackTile := b.level.GetBlackTilePos()
		fmt.Printf("[populateBoard] blackTilePos(%d, %d)\n", blackTile, y)
		for x := 0; x < nCols; x++ {
			var rc *Tile
			fmt.Printf("[populateBoard] board(%d, %d)\n", x, y)
			boardCoords := BoardCoords{x: x, y: y}
			newPosX, newPosY := getPosition(boardCoords)
			if blackTile == x {
				rc = NewTile(newPosX, newPosY, squareWidth, squareHeight, tl.ColorBlack, b.game, boardCoords)
			} else {
				rc = NewTile(newPosX, newPosY, squareWidth, squareHeight, tl.ColorWhite, b.game, boardCoords)
			}
			b.plan[y][x] = *rc
			level.AddEntity(rc)
		}
	}
}
