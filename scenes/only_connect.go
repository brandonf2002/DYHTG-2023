package scenes

import (
	"fmt"

	"github.com/gopxl/pixel"
)

type SceneOnlyConnect struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
}

func NewSceneOnlyConnect(game *Game) *SceneOnlyConnect {
	soc := SceneOnlyConnect{game: game, id: 0}

	//display_text := text.New(pixel.V(100, 500), scc.game.Assets.GetFont("basic"))

	return &soc
}
func (sow *SceneOnlyConnect) AddEntity() *ComponentVector {
	sow.entityManager[sow.id] = ComponentVector{}
	sow.id += 1
	return &sow.entityManager[sow.id-1]
}

func (sow *SceneOnlyConnect) GetEntityManager() EntityManager {
	return sow.entityManager
}

func (sow *SceneOnlyConnect) Update() {
	sow.Render()
}

func (sow *SceneOnlyConnect) Render() {
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

func (sow *SceneOnlyConnect) DoAction(action Action) {
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
			sow.game.ChangeScene("SPIDER", NewSceneSpiders(sow.game))
			sow.game.ChangeScene("TRANSITION", NewSceneTransition(sow.game, "SPIDER"))
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
