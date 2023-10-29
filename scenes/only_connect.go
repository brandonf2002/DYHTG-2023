package scenes

import (
	"fmt"
	"image"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/text"
	"golang.org/x/image/colornames"
)

type SceneOnlyConnect struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
	pressedCount  int
}

const cellSpacing = 40 // Space between rectangles
const gridSize = 4
const cellWidth, cellHeight = 160, 100 // Adjust as needed

func getTextBounds(txt *text.Text) (float64, float64) {
	bounds := txt.Bounds()
	return bounds.W(), bounds.H()
}

func NewSceneOnlyConnect(game *Game) *SceneOnlyConnect {
	soc := SceneOnlyConnect{game: game, id: 0}
	soc.entityManager = make([]ComponentVector, 256)

	//display_text := text.New(pixel.V(100, 500), scc.game.Assets.GetFont("basic"))
	// Calculate starting position for the grid (top-left corner)
	startX := (soc.game.Window.Bounds().W() - gridSize*cellWidth) / 2
	startY := (soc.game.Window.Bounds().H() + gridSize*cellHeight) / 2

	answers := [][]string{
		{"Vorhees", "Beetlejuice", "Thriller", "Ghostbusters"},
		{"The flying dutchman", "IT", "Krueger", "King Boo"},
		{"Lecter", "Monster Mash", "Psycho", "Child's Play"},
		{"Spooky Scary Skeletons", "Casper", "Moaning Myrtle", "Myers"},
	}

	// Assume the correct answers are the first 4 for simplicity
	correctAnswers := map[string]bool{
		"Vorhees":      true,
		"Beetlejuice":  true,
		"Thriller":     true,
		"Ghostbusters": true,
	}

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			entity := soc.AddEntity()
			x := startX + float64(i)*(cellWidth+cellSpacing)
			y := startY - float64(j)*(cellHeight+cellSpacing)
			entity.Transform = CTransform{
				Pos: pixel.V(x, y),
			}
			// entity.Transform = CTransform{
			// 	Pos: pixel.V(startX+float64(i)*cellWidth, startY-float64(j)*cellHeight),
			// }
			entity.Sprite = CSprite{
				Sprite: pixel.NewSprite(nil, pixel.R(0, 0, cellWidth, cellHeight)), // Placeholder sprite, update this
			}
			entity.Text = CText{
				Text:        text.New(entity.Transform.Pos.Add(pixel.V(10, -20)), soc.game.Assets.GetFont("basic")), // Adjust positioning within the cell
				Str_of_Text: fmt.Sprintf("Cell %d,%d", i, j),
			}

			textWidth, _ := getTextBounds(entity.Text.Text)

			centeredTextX := x - (cellWidth+textWidth)/3
			centeredTextY := y
			entity.Text.Str_of_Text = answers[j][i]
			entity.Text.Text.WriteString(entity.Text.Str_of_Text)
			entity.ButtonState = CButtonState{
				IsPressed: false,
				Color:     pixel.RGB(1, 1, 1), // Default color (white)
				Correct:   correctAnswers[entity.Text.Str_of_Text],
			}

			entity.Text.Text = text.New(pixel.V(centeredTextX, centeredTextY), soc.game.Assets.GetFont("basic"))
			entity.Text.Text.WriteString(entity.Text.Str_of_Text)
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
	// sprite := pixel.NewSprite(soc.background, soc.background.Bounds())
	soc.game.Window.Clear(colornames.Midnightblue)
	// sprite := pixel.NewSprite(soc.background, soc.background.Bounds())

	// scaleX := soc.game.Window.Bounds().W() / soc.background.Bounds().W()
	// scaleY := soc.game.Window.Bounds().H() / soc.background.Bounds().H()
	// sprite.Draw(soc.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(soc.game.Window.Bounds().Center()))

	for _, entity := range soc.entityManager {

		if (CTransform{}) != entity.Transform && (CButtonState{}) != entity.ButtonState {
			img := image.NewRGBA(image.Rect(
				int(entity.Sprite.Sprite.Frame().W()),
				int(entity.Sprite.Sprite.Frame().H()), cellWidth, cellHeight))
			col := entity.ButtonState.Color
			for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
				for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
					img.Set(int(x), int(y), col)
				}
			}

			fmt.Printf("Frame: %v\n", entity.Sprite.Sprite.Frame())

			rect := pixel.PictureDataFromImage(img)
			sprite := pixel.NewSprite(rect, rect.Bounds())

			// bottomLeft := entity.Transform.Pos
			// topRight := entity.Transform.Pos.Add(pixel.V(cellWidth, -cellHeight))

			sprite.Draw(soc.game.Window, pixel.IM.Moved(entity.Transform.Pos))
		}

		if (CSprite{}) != entity.Sprite && (CButtonState{}) != entity.ButtonState {
			matrix := pixel.IM.Moved(entity.Transform.Pos)
			entity.Sprite.Sprite.Draw(soc.game.Window, matrix)
		}
		if (CTransform{}) != entity.Transform && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(soc.game.Window, pixel.IM.Moved(entity.Transform.Pos))
		}
		if (CText{}) != entity.Text {
			entity.Text.Text.Draw(soc.game.Window, pixel.IM)
		}
	}
}

func (soc *SceneOnlyConnect) DoAction(action Action) {

	if action.Name == "LEFT_MOUSE" {
		mouseX, mouseY := soc.game.Window.MousePosition().XY()
		for _, entity := range soc.entityManager {
			if entity.Transform.Pos.X <= mouseX && entity.Transform.Pos.X+cellWidth >= mouseX &&
				entity.Transform.Pos.Y-cellHeight <= mouseY && entity.Transform.Pos.Y >= mouseY {
				if !entity.ButtonState.IsPressed {
					entity.ButtonState.IsPressed = true
					if entity.ButtonState.Correct {
						soc.pressedCount++
						entity.ButtonState.Color = pixel.RGB(0, 1, 0) // Change to green if correct
					} else {
						entity.ButtonState.Color = pixel.RGB(1, 0, 0) // Change to red if incorrect
					}

					if soc.pressedCount == 4 {
						// Change color of all correct answers
						for _, e := range soc.entityManager {
							if e.ButtonState.Correct {
								e.ButtonState.Color = pixel.RGB(0, 0, 1) // Change to blue
							}
						}
					}
				}
			}
		}
	}
	// if action.Name == "LEFT_MOUSE" {
	// 	mouseX, mouseY := soc.game.Window.MousePosition().XY()
	// 	fmt.Printf("Mouse X: %v, Mouse Y: %v\n", mouseX, mouseY)
	// 	for _, entity := range soc.entityManager {
	// 		if entity.Transform.Pos.X <= mouseX && entity.Transform.Pos.X+cellWidth >= mouseX &&
	// 			entity.Transform.Pos.Y-cellHeight <= mouseY && entity.Transform.Pos.Y >= mouseY {
	// 			fmt.Printf("Clicked on: %s\n", entity.Text.Str_of_Text)
	// 		}
	// 	}
	// }
	if action.Name == "ESC" {
		soc.game.ChangeScene("MENU", nil)
	}
}
