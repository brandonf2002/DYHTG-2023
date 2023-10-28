package main

import (
	"fmt"

	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/brandonf2002/DYHTG-2023/scenes"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game = struct {
	curScene *scenes.Scene
	score    int
	name     string
}

func newGame(name string, score int, curScene *scenes.Scene) *Game {
	g := Game{name: name, score: score, curScene: curScene}
	return &g
}

// type Scene struct {
// 	name       string
// 	background pixel.Picture
// }

// Notes:
// The start button is from (360, 640) and (700, 740)

// func newScene(name string, background pixel.Picture) *scenes.Scene {
// 	s := scenes.Scene{name: name, background: background}
// 	return &s
// }

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "JPMorgan is an evil bank",
		Bounds: pixel.R(0, 0, 1024, 1024),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	all_scenes := scenes.LoadScenes()
	// am := assets.LoadAssets()
	g := newGame("player", 0, scenes.GetScene("main_menu", all_scenes))

	sprite := pixel.NewSprite(g.curScene.Background, g.curScene.Background.Bounds())

	win.Clear(colornames.Greenyellow)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		// mouseX, mouseY := win.MousePosition().XY()
		// fmt.Printf("Mouse X: %v, Mouse Y: %v\n", mouseX, mouseY)
		win.Clear(colornames.Greenyellow)

		// Adjust the sprite matrix to scale according to the window size (Might delete)
		scaleX := win.Bounds().W() / g.curScene.Background.Bounds().W()
		scaleY := win.Bounds().H() / g.curScene.Background.Bounds().H()
		sprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(win.Bounds().Center()))

		for _, sprite := range g.curScene.InteractiveSprites {
			switch sprite.InteractionType {
			case scenes.MouseClick:
				// println("Mouse Clicked %b, ", win.JustPressed(pixelgl.MouseButtonLeft))
				// println("Mouse in rect %b, ", sprite.Rect.Contains(win.MousePosition()))
				if win.JustPressed(pixelgl.MouseButtonLeft) && sprite.Rect.Contains(win.MousePosition()) {
					sprite.Action()
				}
			case scenes.KeyPress:
				if win.JustPressed(sprite.Key) {
					sprite.Action()
				}
				// case scenes.BoundingBox:
				// 	// Check if any other sprite's bounding box intersects with this sprite's bounding box
				// 	for _, otherSprite := range currentScene.InteractiveSprites {
				// 		if sprite != otherSprite && sprite.Rect.Intersect(otherSprite.Rect) != pixel.ZR {
				// 			sprite.Action()
				// 		}
				// 	}
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
	fmt.Println("Hello, World!")

	am := assets.LoadAssets()
	assets.PlaySound("door_squeak_1", am)
}
