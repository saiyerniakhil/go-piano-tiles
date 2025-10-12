package main

type Difficulty string

const (
	EASY   Difficulty = "easy"
	MEDIUM Difficulty = "medium"
	HARD   Difficulty = "hard"
)

type Level struct {
	difficulty Difficulty
}

// GetBoardDimensions returns the number of rows and columns per difficulty of the level
func (l *Level) GetBoardDimensions() (int, int) {
	switch l.difficulty {
	case EASY:
		return 5, 4
	default:
		return 5, 4 //TODO: Add more levels
	}
}
