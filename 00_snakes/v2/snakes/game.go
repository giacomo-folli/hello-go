package main

import (
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen    tcell.Screen
	snakeBody SnakeBody
}

func drawParts(s tcell.Screen, parts []SnakePart, style tcell.Style) {
	for _, part := range parts {
		s.SetContent(part.X, part.Y, ' ', nil, style)
	}
}

func (g *Game) Run() {

	defStyle := tcell.StyleDefault.Background(tcell.ColorCoral).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		// clear the screen
		g.Screen.Clear()
		// update snakeBody location
		g.snakeBody.Update(width, height)
		// draw on the screen
		drawParts(g.Screen, g.snakeBody.Parts, snakeStyle)
		// time.Sleep
		time.Sleep(40 * time.Millisecond)
		// draw the new frame
		g.Screen.Show()
	}
}
