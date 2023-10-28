package scenes

import (
	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel/pixelgl"
)

type SceneManager map[string]Scene

type Game struct {
	curScene string
	scenes   SceneManager
	Assets   *assets.AssetManager
	Window   *pixelgl.Window
}

func NewGame(window *pixelgl.Window) *Game {
	g := Game{
		curScene: "",
		scenes:   make(map[string]Scene),
		Assets:   assets.LoadAssets(),
		Window:   window,
	}
	g.ChangeScene("MENU", NewSceneMainMenu(&g))
	return &g
}

func (g *Game) Run() {
	for !g.Window.Closed() {
		g.Update()
	}
}

func (g *Game) Update() {
	scene, ok := g.GetCurrentScene()
	if ok {
		scene.Update()
	}
	g.Window.Update()
}

func (g *Game) ChangeScene(name string, scene Scene) {
	if scene != nil {
		g.scenes[name] = scene
	}
	if _, ok := g.scenes[name]; ok {
		g.curScene = name
	}
}

func (g *Game) GetCurrentScene() (Scene, bool) {
	val, ok := g.scenes[g.curScene]
	return val, ok
}
