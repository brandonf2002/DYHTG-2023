package main

import (
	"github.com/brandonf2002/DYHTG-2023/scenes"

	// "github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "JPMorgan is an evil bank",
		Bounds: pixel.R(0, 0, 1024, 1024),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	g := scenes.NewGame(win)
	g.Run()
}

func main() {
	pixelgl.Run(run)
}
