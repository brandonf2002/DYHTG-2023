package main

import (
	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game = struct {
	curScene *Scene
	score    int
	name     string
}

func newGame(name string, score int, curScene *Scene) *Game {
	g := Game{name: name, score: score, curScene: curScene}
	return &g
}

type Scene struct {
	name       string
	background pixel.Picture
}

func newScene(name string, background pixel.Picture) *Scene {
	s := Scene{name: name, background: background}
	return &s
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "JPMorgan is an evil bank",
		Bounds: pixel.R(0, 0, 1024, 1024),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	am := assets.LoadAssets()
	s1 := newScene("Menu", assets.GetPicture("main_menu", am))
	g := newGame("player", 0, s1)
	sprite := pixel.NewSprite(g.curScene.background, g.curScene.background.Bounds())

	win.Clear(colornames.Greenyellow)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Clear(colornames.Greenyellow)

		// Adjust the sprite matrix to scale according to the window size
		scaleX := win.Bounds().W() / g.curScene.background.Bounds().W()
		scaleY := win.Bounds().H() / g.curScene.background.Bounds().H()
		sprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(win.Bounds().Center()))

		win.Update()
	}
	// for !win.Closed() {
	// 	win.Update()
	// }
}

func main() {
	//pixelgl.Run(run)
	am := assets.LoadAssets()
	assets.PlayRandomDoorSound(am)
}
