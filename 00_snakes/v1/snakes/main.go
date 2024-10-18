package snakes

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	snakeBody := SnakeBody{
		X:      0,
		Y:      0,
		Xspeed: 1,
		Yspeed: 0,
	}

	game := Game{
		Screen:    screen,
		snakeBody: snakeBody,
	}

	go game.Run()
	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlC || event.Key() == tcell.KeyEscape {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.snakeBody.ChangeDir(0, -1)
			} else if event.Key() == tcell.KeyDown {
				game.snakeBody.ChangeDir(0, 1)
			} else if event.Key() == tcell.KeyLeft {
				game.snakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.KeyRight {
				game.snakeBody.ChangeDir(1, 0)
			}
		}
	}
}
