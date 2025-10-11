package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	fmt.Println("Hello, World!")
	g := tl.NewGame()
	g.Start()
}
