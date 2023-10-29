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
	startButton.Transform = NewCTransform(pixel.V(364, 290), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
	startButton.BoundingBox = NewCBoundingBox(pixel.V(700-364, 380-290))
	startButton.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("main_menu"), pixel.R(364, 290, 700, 380)))

	optionsButton := smm.AddEntity()
	optionsButton.Transform = NewCTransform(pixel.V(364, 290-125), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
	optionsButton.BoundingBox = NewCBoundingBox(pixel.V(700-364, 380-290))
	optionsButton.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("main_menu"), pixel.R(364, 290-125, 700, 380-125)))

	quitButton := smm.AddEntity()
	quitButton.Transform = NewCTransform(pixel.V(364, 290-250), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
	quitButton.BoundingBox = NewCBoundingBox(pixel.V(700-364, 380-290))
	quitButton.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("main_menu"), pixel.R(364, 290-250, 700, 380-250)))

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
		if (CTransform{}) != entity.Transform && (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(smm.game.Window, pixel.IM.Moved(Add(entity.Transform.Pos, entity.BoundingBox.Half())))
		}
	}
}

func (smm *SceneMainMenu) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {
		if Inside(action.Coords, smm.entityManager[0]) {
			smm.game.ChangeScene("OVERWORLD", NewSceneOverworld(smm.game))
			smm.game.ChangeScene("TRANSITION", NewSceneTransition(smm.game, "OVERWORLD"))
		}
		if Inside(action.Coords, smm.entityManager[1]) {
			smm.game.ChangeScene("OPTIONS", NewSceneOptions(smm.game))
			smm.game.ChangeScene("TRANSITION", NewSceneTransition(smm.game, "OPTIONS"))
		}
		if Inside(action.Coords, smm.entityManager[2]) {
			smm.game.Quit()
		}
	} else if action.Name == "ESC" {
		smm.game.Quit()
	}
}
