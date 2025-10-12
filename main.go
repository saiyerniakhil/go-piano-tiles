package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

// Need to build a screen with a canvas of 5*4 squares
// then call DrawBackground to color them

// Need to build a board ( collection of squares)

func main() {
	fmt.Println("Hello, World!")
	g := tl.NewGame()
	g.Start()
}
