package scenes

import (
	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
)

type SceneMainMenu struct {
	game          *Game
	entityManager EntityManager
	background    pixel.Picture
}

func NewSceneMainMenu(game *Game) *SceneMainMenu {
	smm := SceneMainMenu{game: game, background: assets.GetPicture("main_menu", game.Assets)}
	return &smm
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
}

func (smm *SceneMainMenu) DoAction(action *Action) {

}
