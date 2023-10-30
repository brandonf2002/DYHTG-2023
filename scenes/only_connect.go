package scenes

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/text"
	"golang.org/x/image/colornames"
)

const (
	GroupA int = iota
	GroupB
	GroupC
	GroupD
)

type SceneOnlyConnect struct {
	game           *Game
	entityManager  EntityManager
	id             int
	background     pixel.Picture
	pressedCount   int
	correctAnswers [][]int
	pressed        [][]bool
}

const cellSpacing = 40 // Space between rectangles
const gridSize = 4
const cellWidth, cellHeight = 160, 100 // Adjust as needed

func getTextBounds(txt *text.Text) (float64, float64) {
	bounds := txt.Bounds()
	return bounds.W(), bounds.H()
}

func NewSceneOnlyConnect(game *Game) *SceneOnlyConnect {
	soc := SceneOnlyConnect{game: game, id: 0, pressedCount: 0}
	soc.entityManager = make([]ComponentVector, 256)

	// Calculate starting position for the grid (top-left corner)
	startX := (soc.game.Window.Bounds().W() - gridSize*cellWidth) / 2
	startY := (soc.game.Window.Bounds().H() + gridSize*cellHeight) / 2

	answers := [][]string{
		{"Vorhees", "Beetlejuice", "Thriller", "Ghostbusters"},
		{"The flying dutchman", "IT", "Krueger", "King Boo"},
		{"Lecter", "Monster Mash", "Psycho", "Child's Play"},
		{"Spooky Scary Skeletons", "Casper", "Moaning Myrtle", "Myers"},
	}

	soc.correctAnswers = [][]int{
		{GroupA, GroupC, GroupA, GroupD},
		{GroupC, GroupB, GroupD, GroupB},
		{GroupA, GroupC, GroupD, GroupD},
		{GroupC, GroupB, GroupB, GroupA},
	}

	soc.pressed = make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		soc.pressed[i] = make([]bool, gridSize)
	}

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			entity := soc.AddEntity()
			x := startX + float64(i)*(cellWidth+cellSpacing)
			y := startY - float64(j)*(cellHeight+cellSpacing)
			entity.Transform = NewCTransform(pixel.V(x, y), pixel.ZV, pixel.V(1, 1), pixel.ZV, 0, 0)
			entity.BoundingBox = NewCBoundingBox(pixel.V(cellWidth, cellHeight))
			entity.Sprite = NewCSprite(pixel.NewSprite(game.Assets.GetPicture("door"), pixel.R(0, 0, cellWidth, cellHeight)))

			centeredTextX := x - cellWidth/3
			centeredTextY := y
			entity.Text = NewCText(text.New(pixel.V(centeredTextX, centeredTextY), soc.game.Assets.GetFont("basic")))
			entity.Text.Text.WriteString(answers[j][i])
		}
	}

	return &soc
}
func (soc *SceneOnlyConnect) AddEntity() *ComponentVector {
	soc.entityManager[soc.id] = ComponentVector{}
	soc.id += 1
	return &soc.entityManager[soc.id-1]
}

func (soc *SceneOnlyConnect) GetEntityManager() EntityManager {
	return soc.entityManager
}

func (soc *SceneOnlyConnect) Update() {
	soc.Render()
}

func (soc *SceneOnlyConnect) Render() {
	soc.game.Window.Clear(colornames.Midnightblue)
	imd := imdraw.New(nil)

	for i, entity := range soc.entityManager {
		if i < 16 && soc.pressed[i/4][i%4] {
			imd.Color = pixel.RGB(1, 0, 0)
			imd.Push(entity.Transform.Pos, Add(entity.Transform.Pos, entity.BoundingBox.size))
		}
		imd.Rectangle(5)
		imd.Draw(soc.game.Window)
		if (CTransform{}) != entity.Transform && (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(soc.game.Window, pixel.IM.Moved(Add(entity.Transform.Pos, entity.BoundingBox.Half())))
		}
		if (CText{}) != entity.Text {
			entity.Text.Text.Draw(soc.game.Window, pixel.IM)
		}
	}
}

func (soc *SceneOnlyConnect) DoAction(action Action) {

	if action.Name == "LEFT_MOUSE" {
		for i, entity := range soc.entityManager {
			if (CTransform{}) != entity.Transform && (CBoundingBox{}) != entity.BoundingBox && Inside(action.Coords, entity) {
				if soc.pressedCount < 4 && !soc.pressed[i/4][i%4] {
					soc.pressedCount += 1
					soc.pressed[i/4][i%4] = true
					if soc.pressedCount == 4 && soc.fourCorrect() {
						entity := soc.AddEntity()
						entity.Text = NewCText(text.New(pixel.V(300, 400), soc.game.Assets.GetFont("basic")))
						entity.Text.Text.WriteString("Correct")
						entity.LifeSpan = NewCLifeSpan(180)
					}
				} else if soc.pressed[i/4][i%4] {
					soc.pressedCount -= 1
					soc.pressed[i/4][i%4] = false
				}
			}
		}
	}
	if action.Name == "ESC" {
		soc.game.ChangeScene("MENU", nil)
	}
}

func (soc *SceneOnlyConnect) fourCorrect() bool {
	group := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if soc.pressed[i][j] {
				if group != -1 && soc.correctAnswers[i][j] != group {
					return false
				}
				group = soc.correctAnswers[i][j]
			}
		}
	}
	return true
}
