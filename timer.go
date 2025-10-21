package main

import tl "github.com/JoelOtter/termloop"

type TimerEntity struct {
	game *Game
}

func (T *TimerEntity) Draw(screen *tl.Screen) {

}

func (t *TimerEntity) Tick(event tl.Event) {
	if t.game.state == GameStatePlaying {
		t.game.updateStatusText()
	}
}
