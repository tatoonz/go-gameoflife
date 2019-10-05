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

		for x := 0; x < cells; x++ {
			for y := 0; y < rows; y++ {
				imd.Color = colornames.Black
				if x%2 == 0 {
					imd.Color = colornames.Red
				}

				imd.Push(
					pixel.V(float64(x*cellSize), float64((rows-y-1)*cellSize)),
					pixel.V(float64((x+1)*cellSize), float64((rows-y)*cellSize)),
				)
				imd.Rectangle(0)
			}
		}

		imd.Draw(win)
		win.Update()
	}
}
