package main

import (
	"github.com/brandonf2002/DYHTG-2023/scenes"
	// "github.com/brandonf2002/DYHTG-2023/assets"
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
	g := scenes.NewGame("player", 0, scenes.GetScene("main_menu", all_scenes), win)

	sprite := pixel.NewSprite(g.CurScene.Background, g.CurScene.Background.Bounds())

	win.Clear(colornames.Greenyellow)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		sprite := pixel.NewSprite(g.CurScene.Background, g.CurScene.Background.Bounds())

		// Adjust the sprite matrix to scale according to the window size (Might delete)
		scaleX := win.Bounds().W() / g.CurScene.Background.Bounds().W()
		scaleY := win.Bounds().H() / g.CurScene.Background.Bounds().H()
		sprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(win.Bounds().Center()))

		// mouseX, mouseY := win.MousePosition().XY()
		// fmt.Printf("Mouse X: %v, Mouse Y: %v\n", mouseX, mouseY)

		for _, sprite := range g.CurScene.Entities {
			for _, action := range sprite.Actions {
				switch action.InteractionType {
				case scenes.MouseClick:
					if win.JustPressed(pixelgl.MouseButtonLeft) && sprite.Rect.Contains(win.MousePosition()) {
						action.Action(g, all_scenes)
					}
				case scenes.KeyPress:
					if win.JustPressed(action.Key) {
						action.Action(g, all_scenes)
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

			// sprite.Sprite.Draw(win, pixel.IM.Moved(sprite.Rect.Center()))
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
	// am := assets.LoadAssets()
	// assets.PlayRandomDoorSound(am)
}
