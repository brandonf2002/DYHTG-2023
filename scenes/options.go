package scenes

import (
	"github.com/gopxl/pixel"
)

type SceneOptions struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
}

func NewSceneOptions(game *Game) *SceneOptions {
	sop := SceneOptions{game: game, id: 0, background: game.Assets.GetPicture("menu_background")}
	sop.entityManager = make([]ComponentVector, 64)

	return &sop
}

func (sop *SceneOptions) AddEntity() *ComponentVector {
	sop.entityManager[sop.id] = ComponentVector{}
	sop.id += 1
	return &sop.entityManager[sop.id-1]
}

func (sop *SceneOptions) GetEntityManager() EntityManager {
	return sop.entityManager
}

func (sop *SceneOptions) Update() {
	sop.Render()
}

func (sop *SceneOptions) Render() {
	sprite := pixel.NewSprite(sop.background, sop.background.Bounds())

	scaleX := sop.game.Window.Bounds().W() / sop.background.Bounds().W()
	scaleY := sop.game.Window.Bounds().H() / sop.background.Bounds().H()
	sprite.Draw(sop.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(sop.game.Window.Bounds().Center()))

	for _, entity := range sop.entityManager {
		if (CTransform{}) != entity.Transform && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(sop.game.Window, pixel.IM.Moved(entity.Transform.Pos))
		}
	}
}

func (sop *SceneOptions) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {

	}
	if action.Name == "ESC" {
		sop.game.ChangeScene("MENU", nil)
	}
}
