package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "JPMorgan is an evil bank",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	//am = assets.loadAssets()

	//s1 := newScene{name: "Menu", background: am.pictureMap["menu_background"]}

	//g := Game{name: "player", score: 0, curScene: s1};
	pixelgl.Run(run)
}
