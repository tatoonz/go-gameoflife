package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Game Of Life",
		Bounds: pixel.R(0, 0, 1024, 768),
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
			pixel.V(10, 10),
		)
		imd.Rectangle(0)

		imd.Draw(win)

		win.Update()
	}
}
