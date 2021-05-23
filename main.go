package main

import "fmt"
import "math/rand"
import "os"
import "time"
import "github.com/urfave/cli/v2"
import tl "github.com/JoelOtter/termloop"

type BaseGraphic struct {
	*tl.BaseLevel
}

type Avatar struct {
	*tl.Entity
	m *Mouth
	e *Eye
}

type Mouth struct {
	*tl.Entity
}

type Eye struct {
	*tl.Entity
}

func (a *Avatar) Draw(s *tl.Screen) {
	a.Entity.Draw(s)
}

func (a *Avatar) Tick(ev tl.Event) {
	if ev.Type == tl.EventKey {
		x1, y1 := a.Entity.Position()
		x2, y2 := a.m.Entity.Position()
		x3, y3 := a.e.Entity.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			x1 += 1
			x2 += 1
			x3 += 1
		case tl.KeyArrowLeft:
			x1 -= 1
			x2 -= 1
			x3 -= 1
		case tl.KeyArrowUp:
			y1 -= 1
			y2 -= 1
			y3 -= 1
		case tl.KeyArrowDown:
			y1 += 1
			y2 += 1
			y3 += 1
		}
		a.Entity.SetPosition(x1, y1)
		a.m.Entity.SetPosition(x2, y2)
		a.e.Entity.SetPosition(x3, y3)
	}
}

// create duck avatar body
func CreateBody(color tl.Attr) *tl.Canvas {
	canvas := tl.NewCanvas(500, 500)
	cell := tl.Cell{
		Bg: tl.ColorDefault,
		Fg: tl.RgbTo256Color(252, 226, 5),
	}

	Draw(canvas, cell, 19, 10, "▀▄▄▂")
	Draw(canvas, cell, 26, 10, "▂▄█▀")
	Draw(canvas, cell, 21, 11, "▀█▄")
	Draw(canvas, cell, 25, 11, "▄██▀")
	Draw(canvas, cell, 16, 12, "▂▄▄▄██████████▄▄▄▂")
	Draw(canvas, cell, 13, 13, "▂▄████████████████████▄▂")
	Draw(canvas, cell, 12, 14, "▄████████████████████████▄")
	Draw(canvas, cell, 12, 15, "██████████████████████████")
	Draw(canvas, cell, 12, 16, "██████████████████████████")
	Draw(canvas, cell, 11, 17, "▄██████████████████████████▄")
	Draw(canvas, cell, 9, 18, "▄██████████████████████████████▄")
	Draw(canvas, cell, 8, 19, "▄████████████████████████████████▄")
	Draw(canvas, cell, 8, 20, "▀████████████████████████████████▀")
	Draw(canvas, cell, 10, 21, "▀▀▀▀██████████████████████▀▀▀▀")
	Draw(canvas, cell, 16, 22, "▀▀▀▀██████████▀▀▀▀")
	Draw(canvas, cell, 16, 23, "▂▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▂")
	Draw(canvas, cell, 13, 24, "▄██████████████████████▄")
	Draw(canvas, cell, 11, 25, "▄██████████████████████████▄")
	Draw(canvas, cell, 11, 26, "▄██████████████████████████▄")
	Draw(canvas, cell, 11, 27, "▀██████████████████████████▀")
	Draw(canvas, cell, 12, 28, "▀████████████████████████▀")
	Draw(canvas, cell, 13, 29, "▀▀████████████████████▀▀")
	Draw(canvas, cell, 16, 30, "▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀")

	return &canvas
}

// create duck avatar's mouth
func CreateMouth(color tl.Attr) *tl.Canvas {
	canvas := tl.NewCanvas(500, 500)
	cell := tl.Cell{
		// Bg: tl.RgbTo256Color(255, 250, 80),
		Bg: tl.RgbTo256Color(255, 210, 15),
		Fg: tl.RgbTo256Color(255, 131, 0),
	}

	Draw(canvas, cell, 20, 18, "▂▄▄████▄▄▂")
	Draw(canvas, cell, 13, 19, "▂▄▂▂▄██████████████▄▂▂▄▂")
	Draw(canvas, cell, 13, 20, "▀█▀▀▀███▄▄▂▂▂▂▄▄███▀▀▀█▀")
	Draw(canvas, cell, 19, 21, "▀▀▀██████▀▀▀")

	return &canvas
}

// create duck avatar's eyes
func CreateEyes(color tl.Attr) *tl.Canvas {
	canvas := tl.NewCanvas(500, 500)
	cell := tl.Cell{
		Bg: tl.RgbTo256Color(255, 210, 15),
		Fg: tl.ColorBlack,
	}

	Draw(canvas, cell, 15, 15, "▂▄▀▀▀▄▂")
	Draw(canvas, cell, 28, 15, "▂▄▀▀▀▄▂")
	Draw(canvas, cell, 14, 16, "▄▀▄██▄ █")
	Draw(canvas, cell, 28, 16, "█ ▄██▄▀▄")
	Draw(canvas, cell, 14, 17, "█▄███▂▄▀")
	Draw(canvas, cell, 28, 17, "▀▄▂███▄█")

	return &canvas
}

func Draw(canvas tl.Canvas, cell tl.Cell, x, y int, s string) {
	for i, c := range []rune(s) {
		canvas[x+i][y].Ch = c
		canvas[x+i][y].Bg = cell.Bg
		canvas[x+i][y].Fg = cell.Fg
	}
}

func main() {
	fmt.Println("Hello, Little Duck")
	rand.Seed(time.Now().UTC().UnixNano())
	(&cli.App{}).Run(os.Args)

	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorDefault,
		Fg: tl.ColorBlack,
		Ch: '.',
	})

	mouth := Mouth{tl.NewEntityFromCanvas(0, 0, *CreateMouth(tl.ColorRed))}
	eye := Eye{tl.NewEntityFromCanvas(0, 0, *CreateEyes(tl.ColorRed))}
	avatar := Avatar{
		tl.NewEntityFromCanvas(0, 0, *CreateBody(tl.ColorRed)),
		&mouth,
		&eye,
	}

	level.AddEntity(&avatar)
	level.AddEntity(&mouth)
	level.AddEntity(&eye)
	game.Screen().SetLevel(level)
	game.Start()
}