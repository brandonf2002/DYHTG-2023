package assets

import (
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/gopxl/pixel"
)

type AssetManager struct {
	pictureMap map[string]pixel.Picture
}

func LoadAssets() *AssetManager {
	pictureMap := make(map[string]pixel.Picture)
	am := AssetManager{pictureMap: pictureMap}
	LoadPicture("menu_background", "./assets/png/menu_background.png", &am)
	return &am
}

func LoadPicture(name string, path string, am *AssetManager) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	am.pictureMap[name] = pixel.PictureDataFromImage(img)
}

func GetPicture(name string, am *AssetManager) pixel.Picture {
	return am.pictureMap[name]
}
