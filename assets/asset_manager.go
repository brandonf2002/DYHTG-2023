package assets

import (
	"fmt"
	"os"
	"time"
	"image/png"
	"github.com/gopxl/pixel"
	"github.com/hajimehoshi/oto/v2"
	"github.com/hajimehoshi/go-mp3"
)

type AssetManager struct {
	pictureMap map[string]pixel.Picture
	soundMap map[string]oto.Player
}

func LoadAssets() *AssetManager {
	pictureMap := make(map[string]pixel.Picture)
	soundMap := make(map[string]oto.Player)
	am := AssetManager{pictureMap: pictureMap, soundMap: soundMap}
	loadPicture("menu_background", "./assets/png/menu_background.png", &am)
	loadSound("door_squeak_1", "./assets/audio/door_squeak_1.mp3", &am)
	loadSound("door_squeak_2", "./assets/audio/door_squeak_2.mp3", &am)
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

func loadSound(name string, path string, am *AssetManager) {
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
	am.soundMap[name] = player
	fmt.Printf("%s loaded successfully\n", path)
}

func PlaySound(name string, am *AssetManager) {
	player := am.soundMap[name]
	player.Play()
	for player.IsPlaying() {
        time.Sleep(time.Millisecond)
    }
}