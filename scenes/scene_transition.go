package scenes

func SceneTransition(g *Game, sm *SceneManager) {
	g.CurScene = GetScene("transition", sm)
}