package scenes

import (
	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

func generateStartButton(am *assets.AssetManager) *Entity {
	name := "start_button"

	clikcer := NewEntityAction(func(g *Game, sm *SceneManager) {
		g.CurScene = GetScene("overworld", sm)
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := Entity{
		Name:   name,
		Sprite: nil,
		Rect:   pixel.R(364, 290, 700, 380),
		Actions: []EntityAction{
			clikcer,
		},
	}

	return &player
}

func GenerateMainMenuScene(am *assets.AssetManager) *Scene {
	testing := generateTestClicker(am)
	start_button := generateStartButton(am)
	return NewScene("main_menu", assets.GetPicture("main_menu", am), *testing, *start_button)
}
