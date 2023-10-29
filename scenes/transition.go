package scenes

import (
	"github.com/gopxl/pixel"
	"golang.org/x/image/colornames"
)

type SceneTransition struct {
	game          *Game
	entityManager EntityManager
	id            int
	frameCounter  int
	nextScene     string
}

func NewSceneTransition(game *Game, nextScene string) *SceneTransition {
	str := SceneTransition{game: game, id: 0, nextScene: nextScene}
	str.entityManager = make([]ComponentVector, 64)

	door := str.AddEntity()
	door.Sprite = NewCSprite(pixel.NewSprite(str.game.Assets.GetPicture("door"), str.game.Assets.GetPicture("door").Bounds()))
	door.Transform = NewCTransform(game.Window.Bounds().Center(), pixel.V(0, 0), pixel.V(1, 1), pixel.V(0, 0), 0, 0)
	door.LifeSpan = NewCLifeSpan(300)

	return &str
}

func (str *SceneTransition) AddEntity() *ComponentVector {
	str.entityManager[str.id] = ComponentVector{}
	str.id += 1
	return &str.entityManager[str.id-1]
}

func (str *SceneTransition) GetEntityManager() EntityManager {
	return str.entityManager
}

func (str *SceneTransition) Update() {
	str.sLifespan()
	str.sMovement()
	str.Render()
	str.frameCounter += 1
}

func (str *SceneTransition) sMovement() {
	for i, entity := range str.entityManager {
		if (CTransform{}) != entity.Transform {
			if str.frameCounter == 60 {
				str.entityManager[i].Transform.DeltaScale.X = 0.003
				str.entityManager[i].Transform.DeltaScale.Y = 0.003
			} else if str.frameCounter == 180 {
				str.entityManager[i].Transform.DeltaScale.X = 0
				str.entityManager[i].Transform.DeltaScale.Y = 0
			} else if str.frameCounter == 240 {
				str.entityManager[i].Transform.DeltaScale.X = -0.002
				str.entityManager[i].Transform.Velocity.X = -0.1
			}
			str.entityManager[i].Transform.PrevPos = str.entityManager[0].Transform.Pos
			str.entityManager[i].Transform.Pos.X += str.entityManager[0].Transform.Velocity.X
			str.entityManager[i].Transform.Pos.Y += str.entityManager[0].Transform.Velocity.Y
			str.entityManager[i].Transform.Scale.X += str.entityManager[0].Transform.DeltaScale.X
			str.entityManager[i].Transform.Scale.Y += str.entityManager[0].Transform.DeltaScale.Y
			str.entityManager[i].Transform.Angle += str.entityManager[0].Transform.DeltaAngle
		}
	}
}

func (str *SceneTransition) sLifespan() {
	for i, entity := range str.entityManager {
		if (CLifeSpan{}) != entity.LifeSpan {
			if entity.LifeSpan.FrameCounter >= entity.LifeSpan.NumOfFrames {
				str.game.ChangeScene(str.nextScene, nil)
			}
			str.entityManager[i].LifeSpan.FrameCounter += 1
		}
	}
}

func (str *SceneTransition) Render() {
	// scaleX := str.game.Window.Bounds().W() / str.background.Bounds().W()
	// scaleY := str.game.Window.Bounds().H() / str.background.Bounds().H()

	// sprite.Draw(str.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(str.game.Window.Bounds().Center()))

	str.game.Window.Clear(colornames.Black)

	for _, entity := range str.entityManager {
		if (CTransform{}) != entity.Transform && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(str.game.Window, pixel.IM.ScaledXY(pixel.ZV, entity.Transform.Scale).Rotated(pixel.ZV, entity.Transform.Angle).Moved(entity.Transform.Pos))
		}
	}
}

func (str *SceneTransition) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {

	}
	if action.Name == "ESC" {
		str.game.ChangeScene("MENU", nil)
	}
}
