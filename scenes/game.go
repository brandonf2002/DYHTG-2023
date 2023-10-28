package scenes

import (
	"github.com/gopxl/pixel/pixelgl"
	"github.com/brandonf2002/DYHTG-2023/assets"
)

type SceneManager map[string]Scene

type Game struct {
	curScene string
	scenes SceneManager
	Assets *AssetManager
	Window pixelgl.Window
}

func NewGame(window pixelgl.Window) *Game {
	g := Game{
			curScene: nil,
			scenes: make(map[string]Scene)
			Assets: assets.LoadAssets()
			Window: window
		}
	g.ChangeScene("MENU", NewSceneMainMenu())
	return &g
}

func (g *Game) Run() {
	for !g.Window.Closed() {
		g.Update()
	}
}

func (g *Game) Update() {
	g.CurScene.Update()
	g.Window.Update()
}

func (g *Game) ChangeScene(name string, scene *Scene) {
	if scene != nil {
		scenes[name] = scene
	}
	else {
		_, ok = scenes[name]
		if ok {
			g.curScene = name
		}
	}
}

func (g *Game) GetCurrentScene() (Scene, bool) {
	return g.scenes[g.curScene]
}
	