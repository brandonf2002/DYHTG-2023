package scenes

import (
	"github.com/gopxl/pixel"
)

type OnlyConnect struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
}

func NewSceneOnlyConnect(game *Game) *OnlyConnect {
	soc := OnlyConnect{game: game, id: 0}

	//display_text := text.New(pixel.V(100, 500), scc.game.Assets.GetFont("basic"))

	return &soc
}