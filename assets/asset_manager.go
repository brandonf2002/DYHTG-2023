package assets

import (
	"fmt"
	"github.com/gopxl/pixel"
	"os"
	"image/png"
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

func loadPicture(name string, path string, am *AssetManager) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to load %s: %s\n", path, err)
		return
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		fmt.Printf("failed to decode %s: %s\n", path, err)
		return
	}
	fmt.Printf("%s loaded successfully\n", path)
	am.pictureMap[name] = pixel.PictureDataFromImage(img)
}

func GetPicture(name string, am *AssetManager) pixel.Picture {
	return am.pictureMap[name]
}
