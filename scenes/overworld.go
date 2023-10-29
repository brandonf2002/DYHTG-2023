package scenes

import (
	"fmt"

	"github.com/gopxl/pixel"
)

type SceneOverworld struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
}

func NewSceneOverworld(game *Game) *SceneOverworld {
	sow := SceneOverworld{game: game, id: 0, background: game.Assets.GetPicture("overworld")}
	sow.entityManager = make([]ComponentVector, 64)

	yellowDoor := sow.AddEntity()
	yellowDoor.Transform = NewCTransform(pixel.V(22, 130), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
	yellowDoor.BoundingBox = NewCBoundingBox(pixel.V(135-22, 230))
	// yellowDoor.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("overworld"), pixel.R(364, 290, 700, 380)))

	greenDoor := sow.AddEntity()
	greenDoor.Transform = NewCTransform(pixel.V(215, 130), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
	greenDoor.BoundingBox = NewCBoundingBox(pixel.V(331-215, 230))
	// greenDoor.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("overworld"), pixel.R(364, 290-125, 700, 380-125)))

	redDoor := sow.AddEntity()
	redDoor.Transform = NewCTransform(pixel.V(696, 130), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
	redDoor.BoundingBox = NewCBoundingBox(pixel.V(819-696, 230))
	// redDoor.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("overworld"), pixel.R(364, 290-250, 700, 380-250)))

	blueDoor := sow.AddEntity()
	blueDoor.Transform = NewCTransform(pixel.V(898, 130), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
	blueDoor.BoundingBox = NewCBoundingBox(pixel.V(1012-898, 230))
	// blueDoor.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("overworld"), pixel.R(364, 290-250, 700, 380-250)))

	return &sow
}

func (sow *SceneOverworld) AddEntity() *ComponentVector {
	sow.entityManager[sow.id] = ComponentVector{}
	sow.id += 1
	return &sow.entityManager[sow.id-1]
}

func (sow *SceneOverworld) GetEntityManager() EntityManager {
	return sow.entityManager
}

func (sow *SceneOverworld) Update() {
	sow.Render()
}

func (sow *SceneOverworld) Render() {
	sprite := pixel.NewSprite(sow.background, sow.background.Bounds())

	scaleX := sow.game.Window.Bounds().W() / sow.background.Bounds().W()
	scaleY := sow.game.Window.Bounds().H() / sow.background.Bounds().H()
	sprite.Draw(sow.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(sow.game.Window.Bounds().Center()))

	for _, entity := range sow.entityManager {
		if (CTransform{}) != entity.Transform && (CSprite{}) != entity.Sprite {
			println("Hello")
			entity.Sprite.Sprite.Draw(sow.game.Window, pixel.IM.Moved(entity.Transform.Pos))
		}
	}
}

func (sow *SceneOverworld) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {
		mouseX, mouseY := sow.game.Window.MousePosition().XY()
		fmt.Printf("Mouse X: %v, Mouse Y: %v\n", mouseX, mouseY)
		if Inside(action.Coords, sow.entityManager[0]) {
			println("Yellow door")
			sow.game.ChangeScene("MENU", NewSceneMainMenu(sow.game))
			sow.game.ChangeScene("TRANSITION", NewSceneTransition(sow.game, "MENU"))
		}
		if Inside(action.Coords, sow.entityManager[1]) {
			println("green door")
			sow.game.ChangeScene("CODING", NewSceneCodingChallenge(sow.game))
			sow.game.ChangeScene("TRANSITION", NewSceneTransition(sow.game, "CODING"))
		}
		if Inside(action.Coords, sow.entityManager[2]) {
			println("red door")
			sow.game.ChangeScene("MENU", NewSceneMainMenu(sow.game))
			sow.game.ChangeScene("TRANSITION", NewSceneTransition(sow.game, "MENU"))
		}
		if Inside(action.Coords, sow.entityManager[3]) {
			println("blue door")
			sow.game.ChangeScene("MENU", NewSceneMainMenu(sow.game))
			sow.game.ChangeScene("TRANSITION", NewSceneTransition(sow.game, "MENU"))
		}
	}
	if action.Name == "ESC" {
		sow.game.ChangeScene("MENU", nil)
	}
}
