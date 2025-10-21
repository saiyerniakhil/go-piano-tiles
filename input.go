package main

import tl "github.com/JoelOtter/termloop"

type InputHandler struct {
	game *Game
}

func (i *InputHandler) Draw(screen *tl.Screen) {

}

func (i *InputHandler) Tick(event tl.Event) {
	if event.Type == tl.EventKey && i.game.state == GameStatePlaying {
		key := int(event.Ch - 49)
		if key >= 0 && key < 4 {
			i.game.board.Move(i.game.board.IsValidMove(key))
		}
	}
}
