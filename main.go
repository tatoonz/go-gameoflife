package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	rows     = 5
	cells    = 5
	cellSize = 20
)

func main() {
	pixelgl.Run(run)
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Game Of Life",
		Bounds: pixel.R(0, 0, float64(cells*cellSize), float64(rows*cellSize)),
		VSync:  true,
	})

	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Whitesmoke)

	imd := imdraw.New(nil)
	for !win.Closed() {
		imd.Clear()

		imd.Color = colornames.Black
		imd.Push(
			pixel.V(0, 0),
			pixel.V(float64(cellSize), float64(cellSize)),
		)
		imd.Rectangle(0)

		imd.Draw(win)

		win.Update()
	}
}
