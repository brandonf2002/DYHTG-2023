package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/brandonf2002/DYHTG-2023/assets"
)

type Game = struct {
	curScene *Scene
	score int
	name string
}

func newGame(name string, score int, curScene *Scene) *Game {
	g := Game{name: name, score: score, curScene: curScene}
	return &g
}

type Scene struct {
	name string
	background pixel.Picture
}

func newScene(name string, background pixel.Picture) *Scene {
	s := Scene{name: name, background: background}
	return &s
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "JPMorgan is an evil bank",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	am := assets.LoadAssets()
	s1 := newScene("Menu", assets.GetPicture("menu_background", am))
	g := newGame("player", 0, s1)
	sprite := pixel.NewSprite(g.curScene.background, pixel.R(0, 0, 200, 200))

	//win.Clear(colornames.Skyblue)
	sprite.Draw(win, pixel.IM)

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
