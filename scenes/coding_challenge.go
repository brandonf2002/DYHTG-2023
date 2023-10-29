package scenes

import (
	"fmt"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/text"
)

type SceneCodingChallenge struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
}

func NewSceneCodingChallenge(game *Game) *SceneCodingChallenge {
	sow := SceneCodingChallenge{game: game, id: 0}
	sow.entityManager = make([]ComponentVector, 64)

	display_text := text.New(pixel.V(100, 500), sow.game.Assets.GetFont("basic"))
	fmt.Fprintln(display_text, "Hello, text!")
	fmt.Fprintln(display_text, "I support multiple lines!")
	fmt.Fprintf(display_text, "And I'm an %s, yay!", "io.Writer")

	term := sow.AddEntity()
	term.Text = NewCText(display_text)

	return &sow
}

func (sow *SceneCodingChallenge) AddEntity() *ComponentVector {
	sow.entityManager[sow.id] = ComponentVector{}
	sow.id += 1
	return &sow.entityManager[sow.id-1]
}

func (sow *SceneCodingChallenge) GetEntityManager() EntityManager {
	return sow.entityManager
}

func (sow *SceneCodingChallenge) Update() {

	sow.Render()
}

func (sow *SceneCodingChallenge) Render() {

	sow.game.Window.Clear(pixel.RGB(0, 0, 0))

	// display_text := text.New(pixel.V(100, 500), sow.game.Assets.GetFont("basic"))

	// display_text.Draw(sow.game.Window, pixel.IM)

	// sprite := pixel.NewSprite(sow.background, sow.background.Bounds())

	// scaleX := sow.game.Window.Bounds().W() / sow.background.Bounds().W()
	// scaleY := sow.game.Window.Bounds().H() / sow.background.Bounds().H()

	// sprite.Draw(sow.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(sow.game.Window.Bounds().Center()))

	for _, entity := range sow.entityManager {
		if (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(sow.game.Window, pixel.IM.Moved(entity.BoundingBox.Center()))
		}

		if (CText{}) != entity.Text {
			entity.Text.Text.Draw(sow.game.Window, pixel.IM)
		}

	}
}

func (sow *SceneCodingChallenge) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {
	}

	if action.Name == "ESC" {
		sow.game.ChangeScene("MENU", nil)
	}
}
