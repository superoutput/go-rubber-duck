package main

import "fmt"
// import "os"
// import "github.com/urfave/cli/v2"
import tl "github.com/JoelOtter/termloop"

type Avatar struct {
	*tl.Entity
}

// create canvas with duck avatar
func createCanvas(color tl.Attr) *tl.Canvas {
	canvas := tl.NewCanvas(500, 500)
	draw(canvas, 19, 10, "▀▄▄▂")
	draw(canvas, 26, 10, "▂▄█▀")
	draw(canvas, 21, 11, "▀█▄")
	draw(canvas, 25, 11, "▄██▀")
	draw(canvas, 16, 12, "▂▄▄▄██████████▄▄▄▂")
	draw(canvas, 13, 13, "▂▄████████████████████▄▂")
	draw(canvas, 12, 14, "▄████████████████████████▄")
	draw(canvas, 11, 15, "▄██████████████████████████▄")
	draw(canvas, 11, 16, "▀██████████████████████████▀")
	draw(canvas, 11, 17, "▄██████████████████████████▄")
	draw(canvas, 9, 18, "▂▄████████████████████████████▄▂")
	draw(canvas, 8, 19, "▄████████████████████████████████▄")
	draw(canvas, 8, 20, "▀████████████████████████████████▀")
	draw(canvas, 10, 21, "▀▀▀▀██████████████████████▀▀▀▀")
	draw(canvas, 16, 22, "▀▀▀▀██████████▀▀▀▀")
	draw(canvas, 16, 23, "▂▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▂")
	draw(canvas, 13, 24, "▄██████████████████████▄")
	draw(canvas, 11, 25, "▄██████████████████████████▄")
	draw(canvas, 11, 26, "▄██████████████████████████▄")
	draw(canvas, 11, 27, "▀██████████████████████████▀")
	draw(canvas, 12, 28, "▀████████████████████████▀")
	draw(canvas, 13, 29, "▀▀████████████████████▀▀")
	draw(canvas, 16, 30, "▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀")
	
	return &canvas
}

func draw(canvas tl.Canvas, x, y int, s string) {
	for i, c := range []rune(s) {
		cell := tl.Cell{
			// Bg: tl.ColorBlack,
			Bg: tl.RgbTo256Color(255, 195, 11),
			// Fg: tl.ColorYellow,
			Fg: tl.RgbTo256Color(252, 226, 5),
			Ch: c,
		}
		canvas[x+i][y].Ch = cell.Ch
		canvas[x+i][y].Bg = cell.Bg
		canvas[x+i][y].Fg = cell.Fg
	}
}

func main() {
	fmt.Println("Hello, Little Duck")
	// (&cli.App{}).Run(os.Args)
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorDefault,
		Fg: tl.ColorBlack,
		Ch: '.',
	})
	// level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))
	// player := Player{tl.NewEntity(1, 1, 1, 1)}
	// Set the character at position (0, 0) on the entity.
	// player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	avatar := Avatar{tl.NewEntityFromCanvas(0, 0, *createCanvas(tl.ColorRed))}
	// game.Screen().AddEntity(&avatar)
	level.AddEntity(&avatar)
	game.Screen().SetLevel(level)
	game.Start()
}