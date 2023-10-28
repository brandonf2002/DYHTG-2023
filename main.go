package main

import (
	"fmt"

	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/brandonf2002/DYHTG-2023/ecs2"
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

	all_assets := assets.LoadAssets()
	all_scenes := scenes.LoadScenes(all_assets)
	// am := assets.LoadAssets()
	g := scenes.NewGame("player", 0, scenes.GetScene("main_menu", all_scenes), win)

	sprite := pixel.NewSprite(g.CurScene.Background, g.CurScene.Background.Bounds())

	win.Clear(colornames.Greenyellow)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	oldScene := g.CurScene
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

			sprite.Sprite.Draw(win, pixel.IM.Moved(sprite.Rect.Center()))
		}

		if g.CurScene != oldScene {
			currentScene := g.CurScene
			fmt.Println("Scene changed.")
			scenes.SceneTransition(g, all_scenes, all_assets)
			g.CurScene = currentScene
		}
		oldScene = g.CurScene

		win.Update()
	}
}

// func main() {
// 	// pixelgl.Run(run)
// 	// am := assets.LoadAssets()
// 	// assets.PlayRandomDoorSound(am)

func run2() {
	// ... [Window initialization and other setup here] ...
	cfg := pixelgl.WindowConfig{
		Title:  "JPMorgan is an evil bank",
		Bounds: pixel.R(0, 0, 1024, 1024),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		fmt.Println("Error %v", err)
		panic(err)
	}

	sceneManager := ecs2.NewSceneManager()
	sceneManager.RegisterScene("menu", &ecs2.MenuScene{})

	sceneManager.SetScene("menu")

	// inputSystem := &ecs2.InputSystem{Win: win}
	inputSystem := ecs2.NewInputSystem(win, sceneManager)

	all_assets := assets.LoadAssets()
	all_scenes := scenes.LoadScenes(all_assets)
	// am := assets.LoadAssets()
	g := scenes.NewGame("player", 0, scenes.GetScene("main_menu", all_scenes), win)

	sprite := pixel.NewSprite(g.CurScene.Background, pixel.R(100, 100, 200, 200))

	for !win.Closed() {
		win.Clear(colornames.White) // Clearing with a white background for contrast

		// Draw your sprite
		sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		// Draw the black rectangle (bounding box) of the sprite
		rect := sprite.Frame()
		blackMask := pixel.MakePictureData(pixel.R(0, 0, 1, 1))
		for i := range blackMask.Pix {
			blackMask.Pix[i] = colornames.Black
		}
		blackSprite := pixel.NewSprite(pixel.Picture(blackMask), blackMask.Bounds())

		matrix := pixel.IM
		matrix = matrix.ScaledXY(pixel.ZV, pixel.V(rect.W(), rect.H())) // scale to the sprite's size
		matrix = matrix.Moved(win.Bounds().Center())                    // move to the sprite's position

		blackSprite.DrawColorMask(win, matrix, nil)

		dt := 0.2

		inputSystem.Update(dt, sceneManager.em)
		sceneManager.Update(dt, win)
		sceneManager.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run2)
}

// }
