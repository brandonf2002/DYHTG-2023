package scenes

import (
	"time"

	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type SceneManager map[string]Scene

type Game struct {
	curScene string
	scenes   SceneManager
	Assets   *assets.AssetManager
	Window   *pixelgl.Window
	running  bool
}

func NewGame(window *pixelgl.Window) *Game {
	g := Game{
		curScene: "",
		scenes:   make(map[string]Scene),
		Assets:   assets.LoadAssets(),
		Window:   window,
		running:  true,
	}
	g.ChangeScene("MENU", NewSceneMainMenu(&g))
	return &g
}

func (g *Game) Run() {
	for !g.Window.Closed() && g.running {
		start := time.Now()
		g.Update()
		time.Sleep(time.Second/60 - time.Since(start))
	}
}

func (g *Game) Quit() {
	g.running = false
}

func (g *Game) Update() {
	scene, ok := g.GetCurrentScene()
	if ok {
		g.sUserInput()
		scene.Update()
	}
	g.Window.Update()
}

func (g *Game) sUserInput() {
	scene, _ := g.GetCurrentScene()
	if g.Window.JustPressed(pixelgl.MouseButtonLeft) {
		scene.DoAction(NewAction("LEFT_MOUSE", g.Window.MousePosition()))
	}
	if g.Window.JustPressed(pixelgl.KeyEscape) {
		scene.DoAction(NewAction("ESC", pixel.ZV))
	}
	if g.Window.JustPressed(pixelgl.KeyX) {
		scene.DoAction(NewAction("X", pixel.ZV))
	}
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
