package ecs2

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type MenuScene struct {
	em *EntityManager
}

func (m *MenuScene) Enter() {
	// Initialization when entering the scene
	entity := NewEntity()
	button := ButtonComponent{
		Label:       "Start Game",
		Rect:        pixel.R(100, 100, 300, 150), // Example rectangle for the button
		TargetScene: "gameScene",                 // The scene to transition to when this button is pressed
	}
	m.em.AddComponent(entity, button)
}

func (m *MenuScene) Exit() {
	// Cleanup when exiting the scene
}

func (m *MenuScene) Update(dt float64, win *pixelgl.Window) bool {
	// Handle other updates, if needed
	return false
}

func (m *MenuScene) Draw(win *pixelgl.Window) {
	// Drawing logic for the menu scene, including drawing the button
}
