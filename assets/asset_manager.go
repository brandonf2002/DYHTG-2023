package assets

import (
	"fmt"
	"image/png"
	"os"

	"github.com/gopxl/pixel"
	"github.com/hajimehoshi/oto/v2"
	"github.com/hajimehoshi/go-mp3"
)

type AssetManager struct {
	pictureMap map[string]pixel.Picture
	//soundMap map[string]
}

func LoadAssets() *AssetManager {
	pictureMap := make(map[string]pixel.Picture)
	am := AssetManager{pictureMap: pictureMap}
	loadPicture("menu_background", "./assets/png/menu_background.png", &am)
	loadPicture("main_menu", "./assets/png/main_menu.png", &am)
	loadPicture("overworld", "./assets/png/overworld.png", &am)
	return &am
}

func loadPicture(name string, path string, am *AssetManager) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return
	}
	fmt.Printf("%s loaded successfully\n", path)
	am.pictureMap[name] = pixel.PictureDataFromImage(img)
}

func GetPicture(name string, am *AssetManager) pixel.Picture {
	return am.pictureMap[name]
}

func LoadSound(name string, path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	sound, err := mp3.NewDecoder(file)
	if err != nil {
		return
	}
	context, ready, err := oto.NewContext(sound.SampleRate(), 2, 2)
	if err != nil {
		return
	}
	<-ready

	player := context.NewPlayer(sound)
	player.Play()
}