package main

import "math/rand"

type Difficulty string

const (
	EASY   Difficulty = "easy"
	MEDIUM Difficulty = "medium"
	HARD   Difficulty = "hard"
)

type Dimensions struct {
	nRows int
	nCols int
}
type Level struct {
	difficulty Difficulty
	dimensions Dimensions
}

// GetBoardDimensions returns the number of rows and columns per difficulty of the level
func getBoardDimensions(difficulty Difficulty) (int, int) {
	switch difficulty {
	case EASY:
		return 5, 4
	default:
		return 5, 4 //TODO: Add more levels
	}
}

func NewLevel(d Difficulty) *Level {
	r, c := getBoardDimensions(d)
	return &Level{
		difficulty: d,
		dimensions: Dimensions{
			nRows: r,
			nCols: c,
		},
	}
}

func (l *Level) GetBlackTilePos() int {
	return rand.Intn(l.dimensions.nCols - 1)
}
