package scenes

import (
	"fmt"

	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type InteractionType int

const (
	MouseClick InteractionType = iota
	KeyPress
	BoundingBox
)

type Scene struct {
	name               string
	Background         pixel.Picture
	InteractiveSprites []InteractiveSprite
}

type InteractiveSprite struct {
	Sprite          *pixel.Sprite
	Rect            pixel.Rect
	Action          func()
	InteractionType InteractionType
	Key             pixelgl.Button // for KeyPress InteractionType
}

type SceneManager struct {
	sceneMap map[string]*Scene
}

func NewScene(name string, background pixel.Picture, sprites ...InteractiveSprite) *Scene {
	s := Scene{name: name, Background: background, InteractiveSprites: sprites}
	return &s
}

func LoadScenes() *SceneManager {
	sm := SceneManager{sceneMap: make(map[string]*Scene)}
	am := assets.LoadAssets()

	startButton := InteractiveSprite{
		Sprite:          nil,
		Rect:            pixel.R(364, 290, 700, 380),
		Action:          func() { fmt.Print("Hello from the button") },
		InteractionType: MouseClick,
	}

	playerSprite := InteractiveSprite{
		Sprite:          nil,
		Rect:            pixel.R(100, 100, 150, 150),
		Action:          func() {},
		InteractionType: KeyPress,
		Key:             pixelgl.KeyW,
	}

	sm.sceneMap["main_menu"] = NewScene("main_menu", assets.GetPicture("main_menu", am), startButton, playerSprite)
	sm.sceneMap["overworld"] = NewScene("overworld", assets.GetPicture("overworld", am), startButton, playerSprite)

	return &sm
}

func GetScene(name string, sm *SceneManager) *Scene {
	return sm.sceneMap[name]
}
