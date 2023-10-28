package scenes

import (
	"fmt"

	"github.com/brandonf2002/DYHTG-2023/assets"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type Game = struct {
	CurScene *Scene
	Score    int
	Name     string
}

func NewGame(name string, score int, curScene *Scene) *Game {
	g := Game{Name: name, Score: score, CurScene: curScene}
	return &g
}

type InteractionType int

const (
	MouseClick InteractionType = iota
	KeyPress
	BoundingBox
)

type Scene struct {
	Name               string
	Background         pixel.Picture
	InteractiveSprites []Entity
}

type EntityAction struct {
	Action          func(*Game, *SceneManager)
	InteractionType InteractionType
	Key             pixelgl.Button
}

func NewEntityAction(action func(*Game, *SceneManager), interactionType InteractionType, key pixelgl.Button) EntityAction {
	ea := EntityAction{Action: action, InteractionType: interactionType, Key: key}
	return ea
}

type Entity struct {
	Name    string
	Sprite  *pixel.Sprite
	Rect    pixel.Rect
	Actions []EntityAction
	// Action          func(*Game, *SceneManager)
	// InteractionType InteractionType
	// Key             pixelgl.Button // for KeyPress InteractionType
}

func NewEntity(sprite *pixel.Sprite, rect pixel.Rect, actions ...EntityAction) Entity {
	e := Entity{Sprite: sprite, Rect: rect, Actions: actions}
	return e
}

type SceneManager struct {
	sceneMap map[string]*Scene
}

func NewScene(name string, background pixel.Picture, sprites ...Entity) *Scene {
	s := Scene{Name: name, Background: background, InteractiveSprites: sprites}
	return &s
}

func GetEntity(scene *Scene, name string) *Entity {
	for _, entity := range scene.InteractiveSprites {
		if entity.Name == name {
			return &entity
		}
	}
	return nil
}

func LoadScenes() *SceneManager {
	sm := SceneManager{sceneMap: make(map[string]*Scene)}
	am := assets.LoadAssets()

	start_action_func := func(game *Game, sm *SceneManager) { game.CurScene = GetScene("overworld", sm); fmt.Println("Testing") }
	start_action := NewEntityAction(start_action_func, MouseClick, pixelgl.MouseButtonLeft)

	startButton := Entity{
		Sprite:  nil,
		Rect:    pixel.R(364, 290, 700, 380),
		Actions: []EntityAction{start_action},
	}

	// playerSprite := Entity{
	// 	Sprite:          nil,
	// 	Rect:            pixel.R(100, 100, 150, 150),
	// 	Action:          func(*Game, *SceneManager) {},
	// 	InteractionType: KeyPress,
	// 	Key:             pixelgl.KeyW,
	// }

	sm.sceneMap["main_menu"] = NewScene("main_menu", assets.GetPicture("main_menu", am), startButton)
	sm.sceneMap["overworld"] = NewScene("overworld", assets.GetPicture("overworld", am), startButton)

	return &sm
}

func GetScene(name string, sm *SceneManager) *Scene {
	return sm.sceneMap[name]
}
