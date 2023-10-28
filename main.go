package main

import (
	"github.com/brandonf2002/DYHTG-2023/scenes"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

	all_scenes := scenes.LoadScenes()
	// am := assets.LoadAssets()
	g := scenes.NewGame("player", 0, scenes.GetScene("main_menu", all_scenes))

	sprite := pixel.NewSprite(g.CurScene.Background, g.CurScene.Background.Bounds())

	win.Clear(colornames.Greenyellow)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		sprite := pixel.NewSprite(g.CurScene.Background, g.CurScene.Background.Bounds())

		win.Clear(colornames.Greenyellow)

		// Adjust the sprite matrix to scale according to the window size (Might delete)
		scaleX := win.Bounds().W() / g.CurScene.Background.Bounds().W()
		scaleY := win.Bounds().H() / g.CurScene.Background.Bounds().H()
		sprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(win.Bounds().Center()))

		for _, sprite := range g.CurScene.InteractiveSprites {
			switch sprite.InteractionType {
			case scenes.MouseClick:
				if win.JustPressed(pixelgl.MouseButtonLeft) && sprite.Rect.Contains(win.MousePosition()) {
					sprite.Action(g, all_scenes)
				}
			case scenes.KeyPress:
				if win.JustPressed(sprite.Key) {
					sprite.Action(g, all_scenes)
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
}
