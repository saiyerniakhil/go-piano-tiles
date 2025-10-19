package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Board is a collection of squares
type Board struct {
	plan  *[][]Tile
	level tl.Level
	game  *Game
}

func (b *Board) SetGame(g *Game) {
	b.game = g
}

// Move moves the board
func (b *Board) Move(valid bool) {
	if valid {
		//TODO: From b.plan remove the row -> plan[boardHeigh - 1]
		for _, entity := range (*b.plan)[boardHeight-1] {
			b.level.RemoveEntity(&entity)
		}
		// newRow := b.NewRow()

	} else {
		//stop the timer - "Game Over"
	}
}

func (b *Board) NewRow() []Tile {
	row := make([]Tile, boardWidth)
	blackTile := GetBlackTilePos()
	for i := range boardWidth {
		bc := BoardCoords{x: i, y: 0}
		// place this row at the top
		newPosX, newPosY := getPosition(bc)
		var rc *Tile
		if blackTile == i {
			rc = NewTile(newPosX, newPosY, squareWidth, squareHeight, tl.ColorBlack, b.game, bc)
		} else {
			rc = NewTile(newPosX, newPosY, squareWidth, squareHeight, tl.ColorWhite, b.game, bc)
		}
		row = append(row, *rc)
	}
	return row
}

func (b *Board) IsValidMove(currKey int) bool {
	// x = board width, col
	// y = board height, row
	// get the last row
	if (*b.plan)[boardHeight-1][currKey].Color() == tl.ColorBlack {
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

func NewBoard(level tl.Level) *Board {
	rows, cols := boardHeight, boardWidth
	plan := make([][]Tile, rows)
	for i := 0; i < rows; i++ {
		plan[i] = make([]Tile, cols)
	}
	return &Board{
		plan:  &plan,
		level: level,
	}
}

func (b *Board) populateBoard(level tl.Level) {
	nRows, nCols := boardHeight, boardWidth
	// fmt.Printf("rows: %d cols: %d\n", nRows, nCols)
	for y := 0; y < nRows; y++ {
		blackTile := GetBlackTilePos()
		// fmt.Printf("[populateBoard] blackTilePos(%d, %d)\n", blackTile, y)
		for x := 0; x < nCols; x++ {
			var rc *Tile
			// fmt.Printf("[populateBoard] board(%d, %d)\n", x, y)
			boardCoords := BoardCoords{x: x, y: y}
			newPosX, newPosY := getPosition(boardCoords)
			if blackTile == x {
				rc = NewTile(newPosX, newPosY, squareWidth, squareHeight, tl.ColorBlack, b.game, boardCoords)
			} else {
				rc = NewTile(newPosX, newPosY, squareWidth, squareHeight, tl.ColorWhite, b.game, boardCoords)
			}
			(*b.plan)[y][x] = *rc
			level.AddEntity(rc)
		}
	}
}
