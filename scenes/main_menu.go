package scenes

import (
	"github.com/brandonf2002/DYHTG-2023/assets"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

func generateStartButton(am *assets.AssetManager) *Entity {
	name := "start_button"

	clicker := NewEntityAction(func(g *Game, sm *SceneManager) {
		g.CurScene = GetScene("overworld", sm)
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := NewEntity(name, nil, pixel.R(364, 290, 700, 380), clicker)
	return &player
}

func GenerateMainMenuScene(am *assets.AssetManager) *Scene {
	testing := generateTestClicker(am)
	start_button := generateStartButton(am)
	return NewScene("main_menu", assets.GetPicture("main_menu", am), *testing, *start_button)
}
