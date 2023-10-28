package scenes

import (
	"fmt"

	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

func generateTestClicker(am *assets.AssetManager) *Entity {
	name := "testing"

	clikcer := NewEntityAction(func(g *Game, sm *SceneManager) {
		mouseX, mouseY := g.Win.MousePosition().XY()
		fmt.Printf("Mouse X: %v, Mouse Y: %v\n", mouseX, mouseY)
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := Entity{
		Name:   name,
		Sprite: nil,
		Rect:   pixel.R(0, 0, 1024, 1024),
		Actions: []EntityAction{
			clikcer,
		},
	}

	return &player
}

func generatePlayerEntity(am *assets.AssetManager) *Entity {
	name := "player10"

	clikcer := NewEntityAction(func(g *Game, sm *SceneManager) {
		fmt.Println("Clicked!")
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := Entity{
		Name:   name,
		Sprite: nil,
		Rect:   pixel.R(215, 134, 331, 358),
		Actions: []EntityAction{
			clikcer,
		},
	}

	return &player
}

func GenerateOverworldScene(am *assets.AssetManager) *Scene {
	playerSprite := generatePlayerEntity(am)
	testing := generateTestClicker(am)
	return NewScene("overworld", assets.GetPicture("overworld", am), *playerSprite, *testing)
}
