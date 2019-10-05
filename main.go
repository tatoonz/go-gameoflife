package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

	for !win.Closed() {
		win.Update()
	}
}
