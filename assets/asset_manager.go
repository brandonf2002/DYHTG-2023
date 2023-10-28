package assets

import (
	"fmt"
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/gopxl/pixel"
	"os"
	"image"
)

type AssetManager struct {
	pictureMap map[string]pixel.Picture
}

func LoadAssets() *AssetManager {
	pictureMap := make(map[string]pixel.Picture)
	am := AssetManager{pictureMap: pictureMap}
	loadPicture("menu_background", "./assets/png/menu_background.png", &am)
	return &am
}

func LoadPicture(name string, path string, am *AssetManager) {
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
