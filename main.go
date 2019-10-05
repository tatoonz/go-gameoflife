package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	cellSize = 10
)

func main() {
	pixelgl.Run(run)
}

func run() {
	uni := newUniverse(50, 100)
	uni.randCells()

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Game Of Life",
		Bounds: pixel.R(0, 0, float64(uni.cols*cellSize), float64(uni.rows*cellSize)),
		VSync:  true,
	})

	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Whitesmoke)

	imd := imdraw.New(nil)
	tick := time.Tick(100 * time.Millisecond)
	for !win.Closed() {
		select {
		case <-tick:
			imd.Clear()

			for x := 0; x < uni.cols; x++ {
				for y := 0; y < uni.rows; y++ {
					imd.Color = colornames.Whitesmoke

					if uni.cellAlive(x, y) {
						imd.Color = colornames.Black
					}

					imd.Push(
						pixel.V(float64(x*cellSize), float64((uni.rows-y-1)*cellSize)),
						pixel.V(float64((x+1)*cellSize), float64((uni.rows-y)*cellSize)),
					)
					imd.Rectangle(0)
				}
			}

			imd.Draw(win)
			uni = uni.next()
		}

		win.Update()
	}
}
