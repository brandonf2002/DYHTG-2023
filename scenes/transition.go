package scenes

import (
	"github.com/gopxl/pixel"
	"golang.org/x/image/colornames"
)

type SceneTransition struct {
	game          *Game
	entityManager EntityManager
	id            int
	phase         int
}

func NewSceneTransition(game *Game) *SceneTransition {
	str := SceneTransition{game: game, id: 0}
	str.entityManager = make([]ComponentVector, 64)

	door := str.AddEntity()
	door.BoundingBox = NewCBoundingBox((str.game.Window.GetPos().X/2)-100, (str.game.Window.GetPos().Y/2)-100, 200, 200)
	door.Sprite = NewCSprite(pixel.NewSprite(str.game.Assets.GetPicture("door"), str.game.Assets.GetPicture("door").Bounds()))
	door.Transform = NewCTransform(game.Window.Bounds().Center(), pixel.V(-1, 0), pixel.V(1, 1), pixel.V(0.0003, 0.0003), 0, 0)
	door.Animation = []CAnimation{
		NewCAnimation(0, 0, 0, 0, 0, 0, 60),
		NewCAnimation(30, 30, 0, 0, 0, 0, 120),
		NewCAnimation(0, 0, 0, 0, 0, 0, 60),
		NewCAnimation(-0.0002, 0, -0.01, 0, 0, 0, 60),
	}

	str.phase = 0

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
	str.sMovement()
	str.sAmimation()
	str.Render()
}

func (str *SceneTransition) sMovement() {
	for _, entity := range str.entityManager {
		if (CTransform{}) != entity.Transform {
			entity.Transform.PrevPos = entity.Transform.Pos
			entity.Transform.Pos.X += entity.Transform.Velocity.X
			entity.Transform.Pos.Y += entity.Transform.Velocity.Y
			entity.Transform.Scale.X += entity.Transform.DeltaScale.X
			entity.Transform.Scale.Y += entity.Transform.DeltaScale.Y
			entity.Transform.Angle += entity.Transform.DeltaAngle
		}
	}
}

func (str *SceneTransition) sAmimation() {
	if str.phase > 3 {
		return
	}

	str.entityManager[0].Animation[str.phase].CurrentFrame++
	// str.entityManager[0].Sprite.Sprite.Set(str.entityManager[0].Sprite.Sprite.Picture(), pixel.R(
	// 	str.entityManager[0].Sprite.Sprite.Picture().Bounds().Min.X,
	// 	str.entityManager[0].Sprite.Sprite.Picture().Bounds().Min.Y,
	// 	str.entityManager[0].Sprite.Sprite.Picture().Bounds().Max.X+str.entityManager[0].Animation[str.phase].DeltaScaleX,
	// 	str.entityManager[0].Sprite.Sprite.Picture().Bounds().Max.Y+str.entityManager[0].Animation[str.phase].DeltaScaleY,
	// ))

	if str.entityManager[0].Animation[str.phase].CurrentFrame >= str.entityManager[0].Animation[str.phase].NumOfFrames {
		str.phase++
	}
}

func (str *SceneTransition) Render() {
	// scaleX := str.game.Window.Bounds().W() / str.background.Bounds().W()
	// scaleY := str.game.Window.Bounds().H() / str.background.Bounds().H()

	// sprite.Draw(str.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(str.game.Window.Bounds().Center()))

	str.game.Window.Clear(colornames.Olive)

	for _, entity := range str.entityManager {
		if (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(str.game.Window, pixel.IM.Moved(entity.Transform.Pos).ScaledXY(pixel.ZV, entity.Transform.Scale).Rotated(pixel.ZV, entity.Transform.Angle))
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
