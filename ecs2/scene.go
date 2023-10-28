package ecs2

import "github.com/gopxl/pixel/pixelgl"

type Scene interface {
	Enter()
	Exit()
	Update(dt float64, win *pixelgl.Window) bool // Returns true if scene should change
	Draw(win *pixelgl.Window)
}

type SceneManager struct {
	scenes       map[string]Scene
	currentScene string
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes: make(map[string]Scene),
	}
}

func (sm *SceneManager) RegisterScene(name string, scene Scene) {
	sm.scenes[name] = scene
}

func (sm *SceneManager) SetScene(name string) {
	if sm.currentScene != "" {
		sm.scenes[sm.currentScene].Exit()
	}
	sm.currentScene = name
	sm.scenes[name].Enter()
}

// func (sm *SceneManager) SetScene(name string) {
// 	if sm.currentScene != "" {
// 		sm.scenes[sm.currentScene].Exit()
// 	}
// 	sm.currentScene = name
// 	if scene, ok := sm.scenes[name].(*MenuScene); ok { // Assuming MenuScene as an example
// 		sm.em = scene.em
// 	}
// 	sm.scenes[name].Enter()
// }

func (sm *SceneManager) Update(dt float64, win *pixelgl.Window) {
	if sm.scenes[sm.currentScene].Update(dt, win) {
		// Change to another scene based on your game's logic
		// sm.SetScene("anotherSceneName")
	}
}

func (sm *SceneManager) Draw(win *pixelgl.Window) {
	sm.scenes[sm.currentScene].Draw(win)
}
