package scenes

import (
	"math/rand"

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

	//skull_pic := game.Assets.GetPicture("spooky_man")
	jigsaw_pic := game.Assets.GetPicture("jigsaw")

	// skull := sjs.AddEntity()
	// skull.Draggable = NewCDraggable()
	// skull.Sprite = NewCSprite(pixel.NewSprite(skull_pic, skull_pic.Bounds()))
	// skull.Transform = NewCTransform(game.Window.Bounds().Center(), pixel.ZV, pixel.V(0.5, 0.5), pixel.ZV, 0, 0)
	// skull.BoundingBox = NewCBoundingBox(skull_pic.Bounds().Center())

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			jigsawPieceWidth := jigsaw_pic.Bounds().W() / 3
			jigsawPieceHeight := jigsaw_pic.Bounds().H() / 5

			jigsawPiece := sjs.AddEntity()
			jigsawPiece.Draggable = NewCDraggable()
			jigsawPiece.Sprite = NewCSprite(pixel.NewSprite(jigsaw_pic, pixel.R(float64(i)*jigsawPieceWidth, float64(j)*jigsawPieceHeight, float64(i+1)*jigsawPieceWidth, float64(j+1)*jigsawPieceHeight)))
			jigsawPiece.Transform = NewCTransform(pixel.V(rand.Float64()*(game.Window.Bounds().W()-jigsawPieceWidth), rand.Float64()*(game.Window.Bounds().H()-jigsawPieceHeight)), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
			jigsawPiece.BoundingBox = NewCBoundingBox(pixel.V(jigsawPieceWidth, jigsawPieceHeight))
		}
	}

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
		for i := 0; i < len(sjs.entityManager); i++ {
			if (CDraggable{}) != sjs.entityManager[i].Draggable && Inside(action.Coords, sjs.entityManager[i]) {
				sjs.dragged = &sjs.entityManager[i]
				sjs.dragOffset = Sub(action.Coords, sjs.entityManager[i].Transform.Pos)
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
