package scenes

import (
	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
)

func generatePlayerEntity(am *assets.AssetManager) *Entity {
	name := "player"

	// aMoveLeft := NewEntityAction(func(g *Game, sm *SceneManager) {
	// 	entity := GetEntity(g.CurScene, name)
	// 	entity.Rect = entity.Rect.Moved(pixel.V(-10, 0))
	// }, KeyPress, pixelgl.KeyLeft)

	// aMoveRight := NewEntityAction(func(g *Game, sm *SceneManager) {
	// 	entity := GetEntity(g.CurScene, name)
	// 	entity.Rect = entity.Rect.Moved(pixel.V(10, 0))
	// }, KeyPress, pixelgl.KeyRight)

	pic := assets.GetPicture("player", am)
	sprite := pixel.NewSprite(pic, pic.Bounds())
	player := Entity{
		Name:    name,
		Sprite:  sprite,
		Rect:    pixel.R(100, 100, 150, 150),
		Actions: []EntityAction{
			// aMoveLeft,
			// aMoveRight,
		},
	}

	return &player
}

func GenerateOverworldScene(am *assets.AssetManager) *Scene {
	playerSprite := Entity{
		Sprite: nil,
		Rect:   pixel.R(100, 100, 150, 150),
		// Action:          func(*Game, *SceneManager) {},
		// InteractionType: KeyPress,
		// Key:             pixelgl.KeyW,
	}

	return NewScene("overworld", assets.GetPicture("overworld", am), playerSprite)
}
