package main

import "fmt"
import "math/rand"
import "os"
import "time"
import "github.com/urfave/cli/v2"
import tl "github.com/JoelOtter/termloop"

type Avatar struct {
	*tl.Entity
}

type Mouth struct {
	*tl.Entity
}

type Eye struct {
	*tl.Entity
}

// create duck avatar body
func createBody(color tl.Attr) *tl.Canvas {
	canvas := tl.NewCanvas(500, 500)
	cell := tl.Cell{
		Bg: tl.ColorDefault,
		Fg: tl.RgbTo256Color(252, 226, 5),
	}

	draw(canvas, cell, 19, 10, "▀▄▄▂")
	draw(canvas, cell, 26, 10, "▂▄█▀")
	draw(canvas, cell, 21, 11, "▀█▄")
	draw(canvas, cell, 25, 11, "▄██▀")
	draw(canvas, cell, 16, 12, "▂▄▄▄██████████▄▄▄▂")
	draw(canvas, cell, 13, 13, "▂▄████████████████████▄▂")
	draw(canvas, cell, 12, 14, "▄████████████████████████▄")
	draw(canvas, cell, 12, 15, "██████████████████████████")
	draw(canvas, cell, 12, 16, "██████████████████████████")
	draw(canvas, cell, 11, 17, "▄██████████████████████████▄")
	draw(canvas, cell, 9, 18, "▄██████████████████████████████▄")
	draw(canvas, cell, 8, 19, "▄████████████████████████████████▄")
	draw(canvas, cell, 8, 20, "▀████████████████████████████████▀")
	draw(canvas, cell, 10, 21, "▀▀▀▀██████████████████████▀▀▀▀")
	draw(canvas, cell, 16, 22, "▀▀▀▀██████████▀▀▀▀")
	draw(canvas, cell, 16, 23, "▂▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▂")
	draw(canvas, cell, 13, 24, "▄██████████████████████▄")
	draw(canvas, cell, 11, 25, "▄██████████████████████████▄")
	draw(canvas, cell, 11, 26, "▄██████████████████████████▄")
	draw(canvas, cell, 11, 27, "▀██████████████████████████▀")
	draw(canvas, cell, 12, 28, "▀████████████████████████▀")
	draw(canvas, cell, 13, 29, "▀▀████████████████████▀▀")
	draw(canvas, cell, 16, 30, "▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀")

	return &canvas
}

// create duck avatar's mouth
func createMouth(color tl.Attr) *tl.Canvas {
	canvas := tl.NewCanvas(500, 500)
	cell := tl.Cell{
		// Bg: tl.RgbTo256Color(255, 250, 80),
		Bg: tl.RgbTo256Color(255, 210, 15),
		Fg: tl.RgbTo256Color(255, 131, 0),
	}

	draw(canvas, cell, 20, 18, "▂▄▄████▄▄▂")
	draw(canvas, cell, 13, 19, "▂▄▂▂▄██████████████▄▂▂▄▂")
	draw(canvas, cell, 13, 20, "▀█▀▀▀███▄▄▂▂▂▂▄▄███▀▀▀█▀")
	draw(canvas, cell, 19, 21, "▀▀▀██████▀▀▀")

	return &canvas
}

// create duck avatar's eyes
func createEyes(color tl.Attr) *tl.Canvas {
	canvas := tl.NewCanvas(500, 500)
	cell := tl.Cell{
		Bg: tl.RgbTo256Color(255, 210, 15),
		Fg: tl.ColorBlack,
	}

	draw(canvas, cell, 17, 15, "▂▄▀▀▀▄▂")
	draw(canvas, cell, 27, 15, "▂▄▀▀▀▄▂")
	draw(canvas, cell, 16, 16, "▄▀▄██▄ █")
	draw(canvas, cell, 27, 16, "█ ▄██▄▀▄")
	draw(canvas, cell, 16, 17, "█▄███▂▄▀")
	draw(canvas, cell, 27, 17, "▀▄▂███▄█")

	return &canvas
}

func draw(canvas tl.Canvas, cell tl.Cell, x, y int, s string) {
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

	avatar := Avatar{tl.NewEntityFromCanvas(0, 0, *createBody(tl.ColorRed))}
	mouth := Mouth{tl.NewEntityFromCanvas(0, 0, *createMouth(tl.ColorRed))}
	eye := Eye{tl.NewEntityFromCanvas(0, 0, *createEyes(tl.ColorRed))}
	level.AddEntity(&avatar)
	level.AddEntity(&mouth)
	level.AddEntity(&eye)
	game.Screen().SetLevel(level)
	game.Start()
}