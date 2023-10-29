package scenes

import (
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
			entity.Sprite.Sprite.Draw(sow.game.Window, pixel.IM.Moved(entity.Transform.Pos))
		}
	}
}

func (sow *SceneOverworld) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {

	}
	if action.Name == "ESC" {
		sow.game.ChangeScene("MENU", nil)
	}
}
