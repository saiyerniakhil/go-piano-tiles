package main

import (
	tl "github.com/JoelOtter/termloop"
)

type BoardCoords struct {
	x int
	y int
}

type Tile struct {
	x           int
	y           int
	boardCoords BoardCoords
	width       int
	height      int
	color       tl.Attr
	g           *Game
}

func NewTile(x, y, w, h int, color tl.Attr, g *Game, boardCoords BoardCoords) *Tile {
	t := Tile{x: x, y: y, height: h, width: w, color: color, g: g, boardCoords: boardCoords}
	return &t
}

func (t *Tile) Draw(s *tl.Screen) {
	for i := 0; i < t.width; i++ {
		for j := 0; j < t.height; j++ {
			s.RenderCell(t.x+i, t.y+j, &tl.Cell{Bg: t.color, Ch: ' '})
		}
	}
}

// TODO: Implement this
// 49 - Num1
// 50 - Num2
// 51 - Num3
// 52 - Num4
func (r *Tile) Tick(ev tl.Event) {
	if ev.Type == tl.EventKey {
		key := int(ev.Ch - 49)
		if key >= 0 && key < 4 {
			r.g.board.Move(r.g.board.IsValidMove(key))
		}
	}
}

// Size returns the width and height in characters of the Tile.
func (r *Tile) Size() (int, int) {
	return r.width, r.height
}

// Position returns the x and y coordinates of the Tile.
func (r *Tile) Position() (int, int) {
	return r.x, r.y
}

// SetPosition sets the coordinates of the Tile to be x and y.
func (r *Tile) SetPosition(x, y int) {
	r.x = x
	r.y = y
}

// SetSize sets the width and height of the Tile to be w and h.
func (r *Tile) SetSize(w, h int) {
	r.width = w
	r.height = h
}

// Color returns the color of the Tile.
func (r *Tile) Color() tl.Attr {
	return r.color
}

// SetColor sets the color of the Tile.
func (r *Tile) SetColor(color tl.Attr) {
	r.color = color
}
