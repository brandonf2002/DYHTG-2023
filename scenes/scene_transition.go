package scenes

import (
	"time"
	"github.com/gopxl/pixel"
)

func SceneTransition(g *Game, sm *SceneManager) {
	g.CurScene = GetScene("transition", sm)
	win := g.Win
	sprite := pixel.NewSprite(g.CurScene.Background, g.CurScene.Background.Bounds())
	scaleX := win.Bounds().W() / g.CurScene.Background.Bounds().W()
	scaleY := win.Bounds().H() / g.CurScene.Background.Bounds().H()
	sprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(win.Bounds().Center()))

	for start := time.Now(); time.Since(start) < 5*time.Second; {
		win.Update()
	}
}