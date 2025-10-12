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

// NewSquare returns a square at a given position
func NewSquare(boardX, boardY int) Square {
	sq := Square{
		value:   0,
		content: tl.Cell{},
		entity:  tl.NewEntity(1, 1, 20, 20),
	}
	return sq
}
