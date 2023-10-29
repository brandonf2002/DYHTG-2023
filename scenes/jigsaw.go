package scenes

import (
	"github.com/gopxl/pixel"
)

type SceneJigsaw struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
	dragged       *ComponentVector
	dragOffset    pixel.Vec
}

func NewSceneJigsaw(game *Game) *SceneJigsaw {
	sjs := SceneJigsaw{game: game, id: 0, background: game.Assets.GetPicture("menu_background")}
	sjs.entityManager = make([]ComponentVector, 64)

	skull_pic := game.Assets.GetPicture("spooky_man")

	skull := sjs.AddEntity()
	skull.Draggable = NewCDraggable()
	skull.Sprite = NewCSprite(pixel.NewSprite(skull_pic, skull_pic.Bounds()))
	skull.Transform = NewCTransform(game.Window.Bounds().Center(), pixel.ZV, pixel.V(0.5, 0.5), pixel.ZV, 0, 0)
	skull.BoundingBox = NewCBoundingBox(skull_pic.Bounds().Center())

	return &sjs
}

func (sjs *SceneJigsaw) AddEntity() *ComponentVector {
	sjs.entityManager[sjs.id] = ComponentVector{}
	sjs.id += 1
	return &sjs.entityManager[sjs.id-1]
}

func (sjs *SceneJigsaw) GetEntityManager() EntityManager {
	return sjs.entityManager
}

func (sjs *SceneJigsaw) Update() {
	sjs.sDrag()
	sjs.Render()
}

func (sjs *SceneJigsaw) sDrag() {
	if sjs.dragged != nil {
		sjs.dragged.Transform.Pos = Sub(sjs.game.Window.MousePosition(), sjs.dragOffset)
	}
}

func (sjs *SceneJigsaw) Render() {
	sprite := pixel.NewSprite(sjs.background, sjs.background.Bounds())

	scaleX := sjs.game.Window.Bounds().W() / sjs.background.Bounds().W()
	scaleY := sjs.game.Window.Bounds().H() / sjs.background.Bounds().H()
	sprite.Draw(sjs.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(sjs.game.Window.Bounds().Center()))

	for _, entity := range sjs.entityManager {
		if (CTransform{}) != entity.Transform && (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(sjs.game.Window, pixel.IM.ScaledXY(pixel.ZV, entity.Transform.Scale).Rotated(pixel.ZV, entity.Transform.Angle).Moved(Add(entity.Transform.Pos, entity.BoundingBox.Half())))
		}
	}
}

func (sjs *SceneJigsaw) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {
		for i, entity := range sjs.entityManager {
			if (CDraggable{}) != entity.Draggable {
				sjs.dragged = &sjs.entityManager[i]
				sjs.dragOffset = Sub(action.Coords, entity.Transform.Pos)
			}
		}
	}
	if action.Name == "LEFT_MOUSE_RELEASED" {
		sjs.dragged = nil
	}
	if action.Name == "ESC" {
		sjs.game.ChangeScene("OVERWORLD", nil)
	}
}
