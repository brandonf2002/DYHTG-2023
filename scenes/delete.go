package scenes

// import (
// 	"fmt"

// 	"github.com/brandonf2002/DYHTG-2023/assets"

// 	"github.com/gopxl/pixel"
// 	"github.com/gopxl/pixel/pixelgl"
// )

// type Game = struct {
// 	CurScene *Scene
// 	Score    int
// 	Name     string
// }

// func NewGame(name string, score int, curScene *Scene) *Game {
// 	g := Game{Name: name, Score: score, CurScene: curScene}
// 	return &g
// }

// type InteractionType int

// const (
// 	MouseClick InteractionType = iota
// 	KeyPress
// 	BoundingBox
// )

// type Scene struct {
// 	Name               string
// 	Background         pixel.Picture
// 	InteractiveSprites []InteractiveSprite
// }

// type InteractiveSprite struct {
// 	Sprite          *pixel.Sprite
// 	Rect            pixel.Rect
// 	Action          func(*Game, *SceneManager)
// 	InteractionType InteractionType
// 	Key             pixelgl.Button // for KeyPress InteractionType
// }

// type SceneManager struct {
// 	sceneMap map[string]*Scene
// }

// func NewScene(name string, background pixel.Picture, sprites ...InteractiveSprite) *Scene {
// 	s := Scene{Name: name, Background: background, InteractiveSprites: sprites}
// 	return &s
// }

// func LoadScenes() *SceneManager {
// 	sm := SceneManager{sceneMap: make(map[string]*Scene)}
// 	am := assets.LoadAssets()

// 	startButton := InteractiveSprite{
// 		Sprite:          nil,
// 		Rect:            pixel.R(364, 290, 700, 380),
// 		Action:          func(game *Game, sm *SceneManager) { game.CurScene = GetScene("overworld", sm); fmt.Println("Testing") },
// 		InteractionType: MouseClick,
// 	}

// 	playerSprite := InteractiveSprite{
// 		Sprite:          nil,
// 		Rect:            pixel.R(100, 100, 150, 150),
// 		Action:          func(*Game, *SceneManager) {},
// 		InteractionType: KeyPress,
// 		Key:             pixelgl.KeyW,
// 	}

// 	sm.sceneMap["main_menu"] = NewScene("main_menu", assets.GetPicture("main_menu", am), startButton, playerSprite)
// 	sm.sceneMap["overworld"] = NewScene("overworld", assets.GetPicture("overworld", am), startButton, playerSprite)

// 	return &sm
// }

// func GetScene(name string, sm *SceneManager) *Scene {
// 	return sm.sceneMap[name]
// }
