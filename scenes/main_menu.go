package scenes

import (
	"github.com/gopxl/pixel"
)

type SceneMainMenu struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
}

func NewSceneMainMenu(game *Game) *SceneMainMenu {
	smm := SceneMainMenu{game: game, id: 0, background: game.Assets.GetPicture("main_menu")}
	smm.entityManager = make([]ComponentVector, 64)

	startButton := smm.AddEntity()
	startButton.BoundingBox = NewCBoundingBox(0, 0, 200, 50)
	startButton.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("main_menu"), pixel.R(300, 300, 500, 350)))

	optionsButton := smm.AddEntity()
	optionsButton.BoundingBox = NewCBoundingBox(0, 50, 200, 50)
	optionsButton.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("main_menu"), pixel.R(300, 200, 500, 250)))

	quitButton := smm.AddEntity()
	quitButton.BoundingBox = NewCBoundingBox(0, 100, 200, 50)
	quitButton.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("main_menu"), pixel.R(300, 250, 500, 300)))

	return &smm
}

func (smm *SceneMainMenu) AddEntity() *ComponentVector {
	smm.entityManager[smm.id] = ComponentVector{}
	smm.id += 1
	return &smm.entityManager[smm.id-1]
}

func (smm *SceneMainMenu) GetEntityManager() EntityManager {
	return smm.entityManager
}

func (smm *SceneMainMenu) Update() {
	smm.Render()
}

func (smm *SceneMainMenu) Render() {
	sprite := pixel.NewSprite(smm.background, smm.background.Bounds())

	scaleX := smm.game.Window.Bounds().W() / smm.background.Bounds().W()
	scaleY := smm.game.Window.Bounds().H() / smm.background.Bounds().H()
	sprite.Draw(smm.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(smm.game.Window.Bounds().Center()))

	for _, entity := range smm.entityManager {
		if (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(smm.game.Window, pixel.IM.Moved(entity.BoundingBox.Center()))
		}
	}
}

func (smm *SceneMainMenu) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {
		if smm.entityManager[0].BoundingBox.Inside(action.Coords) {
			smm.game.ChangeScene("OVERWORLD", nil)
		}
		if smm.entityManager[1].BoundingBox.Inside(action.Coords) {
			smm.game.ChangeScene("OPTIONS", nil)
		}
		if smm.entityManager[2].BoundingBox.Inside(action.Coords) {
			smm.game.Quit()
		}
	}
}
