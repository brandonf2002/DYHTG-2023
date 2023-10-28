package scenes

import (
	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
)

func new_generateTestClicker(am *assets.AssetManager) *Entity {
	name := "testing"

	// clicker := NewEntityAction(func(g *Game, sm *SceneManager) {
	// 	mouseX, mouseY := g.Win.MousePosition().XY()
	// 	fmt.Printf("Mouse X: %v, Mouse Y: %v\n", mouseX, mouseY)

	// 	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	// 	basicTxt := text.New(pixel.V(100, 500), basicAtlas)

	// 	fmt.Fprintln(basicTxt, "Hello, text!")
	// 	fmt.Fprintln(basicTxt, "I support multiple lines!")
	// 	fmt.Fprintf(basicTxt, "And I'm an %s, yay!", "io.Writer")

	// 	g.Win.Clear(colornames.Black)
	// 	basicTxt.Draw(g.Win, pixel.IM)
	// 	g.Win.Update()
	// }, MouseClick, pixelgl.MouseButtonLeft)

	player := NewEntity(name, nil, pixel.R(0, 0, 1024, 1024))
	return &player
}

// func generateStartButton(am *assets.AssetManager) *Entity {
// 	name := "start_button"

// 	clicker := NewEntityAction(func(g *Game, sm *SceneManager) {
// 		g.CurScene = GetScene("overworld", sm)
// 	}, MouseClick, pixelgl.MouseButtonLeft)

// 	player := NewEntity(name, nil, pixel.R(364, 290, 700, 380), clicker)
// 	return &player
// }

func GenerateCodingChallengeScene(am *assets.AssetManager) *Scene {
	testing := new_generateTestClicker(am)
	return NewScene("coding_challenge", assets.GetPicture("black", am), *testing)
}
