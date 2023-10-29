package scenes

import (
	"strconv"

	"github.com/gopxl/pixel"
	"golang.org/x/image/colornames"
)

type SceneSpiders struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
}

func NewSceneSpiders(game *Game) *SceneSpiders {
	ssp := SceneSpiders{game: game, id: 0, background: game.Assets.GetPicture("menu_background")}
	ssp.entityManager = make([]ComponentVector, 64)

	spiderSprites := make([]*pixel.Sprite, 5)
	candySprites := make([]*pixel.Sprite, 5)

	for i := 0; i < 5; i++ {
		spider_pic := game.Assets.GetPicture("spider" + strconv.Itoa(i+1))
		spiderSprites[i] = pixel.NewSprite(spider_pic, spider_pic.Bounds())
		candy_pic := game.Assets.GetPicture("candy" + strconv.Itoa(i+1))
		candySprites[i] = pixel.NewSprite(candy_pic, candy_pic.Bounds())

		spider := ssp.AddEntity()
		spider.Transform = NewCTransform(pixel.V(200, float64(100+i*64)), pixel.ZV, pixel.V(0.5, 0.5), pixel.ZV, 0, 0)
		spider.BoundingBox = NewCBoundingBox(pixel.V(64, 64))
		spider.Sprite = NewCSprite(spiderSprites[i])

		candy := ssp.AddEntity()
		candy.Transform = NewCTransform(pixel.V(600, float64(100+i*64)), pixel.ZV, pixel.V(0.5, 0.5), pixel.ZV, 0, 0)
		candy.BoundingBox = NewCBoundingBox(pixel.V(64, 64))
		candy.Sprite = NewCSprite(candySprites[i])
	}

	return &ssp
}

func (ssp *SceneSpiders) AddEntity() *ComponentVector {
	ssp.entityManager[ssp.id] = ComponentVector{}
	ssp.id += 1
	return &ssp.entityManager[ssp.id-1]
}

func (ssp *SceneSpiders) GetEntityManager() EntityManager {
	return ssp.entityManager
}

func (ssp *SceneSpiders) Update() {
	ssp.Render()
}

func (ssp *SceneSpiders) Render() {
	//sprite := pixel.NewSprite(ssp.background, ssp.background.Bounds())

	//scaleX := ssp.game.Window.Bounds().W() / ssp.background.Bounds().W()
	//scaleY := ssp.game.Window.Bounds().H() / ssp.background.Bounds().H()
	//sprite.Draw(ssp.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(ssp.game.Window.Bounds().Center()))

	ssp.game.Window.Clear(colornames.Antiquewhite)

	for _, entity := range ssp.entityManager {
		if (CTransform{}) != entity.Transform && (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(ssp.game.Window, pixel.IM.ScaledXY(pixel.ZV, entity.Transform.Scale).Rotated(pixel.ZV, entity.Transform.Angle).Moved(Add(entity.Transform.Pos, entity.BoundingBox.Half())))
		}
	}
}

func (ssp *SceneSpiders) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {

	}
	if action.Name == "ESC" {
		ssp.game.ChangeScene("MENU", nil)
	}
}
