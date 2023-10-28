package ecs2

import "github.com/gopxl/pixel/pixelgl"

type InputSystem struct {
	win          *pixelgl.Window
	sceneManager *SceneManager
}

func NewInputSystem(win *pixelgl.Window, sm *SceneManager) *InputSystem {
	return &InputSystem{
		win:          win,
		sceneManager: sm,
	}
}

func (is *InputSystem) Update(dt float64, em *EntityManager) {
	if is.win.JustPressed(pixelgl.MouseButtonLeft) {
		mousePos := is.win.MousePosition()

		for e, mask := range em.entities {
			if mask&(1<<ButtonComponent{}.Type()) != 0 {
				button := em.GetComponent(e, ButtonComponent{}.Type()).(ButtonComponent)
				if button.Rect.Contains(mousePos) {
					// Transition to the target scene
					is.sceneManager.SetScene(button.TargetScene)
				}
			}
		}
	}
}
