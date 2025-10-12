package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

/***
 A 5 * 4 board looks like this
	[0] [1] [0] [0] -- 1st row
	[0] [1] [0] [0] -- 2nd row
	[0] [0] [1] [0] -- 3rd
	[0] [0] [1] [0] -- 4th
	[0] [1] [0] [0] -- 5th
*/

const (
	SPEED = 2 * time.Second
)

type Board struct {
	plan           [][]int
	nRows          int
	nCols          int
	totalLevelRows int
}

// NewBoard return a new board with a plan
func NewBoard() *Board {
	startBoard := Board{
		nRows:          5,
		nCols:          4,
		plan:           make([][]int, 5),
		totalLevelRows: 20,
	}
	startBoard.levelGenerator()
	startBoard.levelGenerator()
	startBoard.levelGenerator()
	startBoard.levelGenerator()
	startBoard.levelGenerator()
	return &startBoard
}

// levelGenerator generates level on every successful input
// which means on every successful input removes row1, pushed rows 2 - 5 forward
// and appends a new row at 5
func (b *Board) levelGenerator() {
	if b.totalLevelRows > 0 {
		b.plan = b.plan[1:]
		b.plan = append(b.plan, getNewRandomRow(b.nCols))
		b.totalLevelRows--
	}
}

// getNewRow generates a new random number between 1..4
func getNewRandomRow(size int) []int {
	row := make([]int, size)
	blackTile := rand.Intn(size - 1)
	row[blackTile] = 1
	return row
}

func (b *Board) paintBoard(speed time.Duration) {
	for i := len(b.plan) - 1; i >= 0; i-- {
		fmt.Println(b.plan[i])
	}
	time.Sleep(speed)
	if b.totalLevelRows > 0 {
		b.moveNext()
	}
}

func (b *Board) moveNext() {
	if err := clearScreen(); err != nil {
		fmt.Printf("paint board: %s", err)
	}
	b.levelGenerator()

	b.paintBoard(SPEED)
}

func (b *Board) tick() (int, error) {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		return input, fmt.Errorf("read input: %w", err)
	}
	return -1, nil
}

func (b *Board) ValidateInput(input int) (bool, error) {
	// validate
	// value should be between 1...4
	// value shouldn't be > 4 and < 1
	if input >= 1 && input <= 4 {
		if b.plan[0][input] == 1 { // check the start/latest row input is equal to 1 (black)
			return true, nil
		} else {
			return false, fmt.Errorf("wrong tile")
		}
	} else {
		return false, fmt.Errorf("validate input: invalid input")
	}
}

func clearScreen() error {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("clear terminal: %w \n", err)
		}
		return nil
	} else {
		// ANSI escape code for clearing the screen and moving cursor to home position
		fmt.Print("\033[H\033[2J")
	}
	return nil
}

func main() {
	b := NewBoard()
	b.paintBoard(SPEED)
}
