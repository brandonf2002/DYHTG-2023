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
	scc := SceneCodingChallenge{game: game, id: 0}
	scc.entityManager = make([]ComponentVector, 64)

	display_text := text.New(pixel.V(100, 500), scc.game.Assets.GetFont("basic"))
	fmt.Fprintln(display_text, "Hello, text!")
	fmt.Fprintln(display_text, "I support multiple lines!")
	fmt.Fprintf(display_text, "And I'm an %s, yay!", "io.Writer")

	term := scc.AddEntity()
	term.Text = NewCText(display_text)

	return &scc
}

func (scc *SceneCodingChallenge) AddEntity() *ComponentVector {
	scc.entityManager[scc.id] = ComponentVector{}
	scc.id += 1
	return &scc.entityManager[scc.id-1]
}

func (scc *SceneCodingChallenge) GetEntityManager() EntityManager {
	return scc.entityManager
}

func (scc *SceneCodingChallenge) Update() {

	scc.Render()
}

func (scc *SceneCodingChallenge) Render() {

	scc.game.Window.Clear(pixel.RGB(0, 0, 0))

	// display_text := text.New(pixel.V(100, 500), scc.game.Assets.GetFont("basic"))

	// display_text.Draw(scc.game.Window, pixel.IM)

	// sprite := pixel.NewSprite(scc.background, scc.background.Bounds())

	// scaleX := scc.game.Window.Bounds().W() / scc.background.Bounds().W()
	// scaleY := scc.game.Window.Bounds().H() / scc.background.Bounds().H()

	// sprite.Draw(scc.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(scc.game.Window.Bounds().Center()))

	scc.entityManager[0].Text.Text.WriteString(scc.game.Window.Typed())

	for _, entity := range scc.entityManager {
		if (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(scc.game.Window, pixel.IM.Moved(entity.BoundingBox.Half()))
		}

		if (CText{}) != entity.Text {
			entity.Text.Text.Draw(scc.game.Window, pixel.IM)
		}

	}
}

func (scc *SceneCodingChallenge) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {
	}

	if action.Name == "ESC" {
		scc.game.ChangeScene("MENU", nil)
	}
}
