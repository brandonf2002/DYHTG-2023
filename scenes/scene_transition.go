package scenes

import (
	"time"

	"github.com/brandonf2002/DYHTG-2023/assets"

	"github.com/gopxl/pixel"
)

func SceneTransition(g *Game, sm *SceneManager, am *assets.AssetManager) {
	g.CurScene = GetScene("transition", sm)
	win := g.Win
	backgroundSprite := pixel.NewSprite(g.CurScene.Background, g.CurScene.Background.Bounds())
	scaleX := win.Bounds().W() / g.CurScene.Background.Bounds().W()
	scaleY := win.Bounds().H() / g.CurScene.Background.Bounds().H()
	backgroundSprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(win.Bounds().Center()))

	doorImg := assets.GetPicture("door_1", am)
	doorSprite := pixel.NewSprite(doorImg, doorImg.Bounds())
	doorX := 1.0
	doorY := 1.0

	doorSprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(doorX, doorY)).Moved(win.Bounds().Center()))
	win.Update()
	for start := time.Now(); time.Since(start) < 1*time.Second && !win.Closed() ; {}

	for start := time.Now(); time.Since(start) < 2*time.Second && !win.Closed() ; {
		doorSprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(doorX, doorY)).Moved(win.Bounds().Center()))

		doorX += 0.0003
		doorY += 0.0003

		win.Update()
	}

	for start := time.Now(); time.Since(start) < 1*time.Second && !win.Closed() ; {}

	doorPosition := win.Bounds().Center()
	positionChange := pixel.V(-0.01, 0)

	doorSound := assets.GetRandomDoorSound(am)
	doorSound.Play()
	for doorSound.IsPlaying() && !win.Closed() {
		backgroundSprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(win.Bounds().Center()))
		doorSprite.Draw(win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(doorX, doorY)).Moved(doorPosition))

		doorX -= 0.0002
		doorPosition = doorPosition.Add(positionChange)

		win.Update()
	}
}