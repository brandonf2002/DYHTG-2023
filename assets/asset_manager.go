package assets

import (
	"github.com/gopxl/pixel"
	"os"
	"image"
)

type AssetManager struct {
	pictureMap map[string]pixel.Picture
}

func loadAssets() *AssetManager {
	pictureMap := make(map[string]pixel.Picture)
	am := AssetManager{pictureMap: pictureMap}
	loadPicture("menu_background", "png/menu_backgound.png", &am)
	return &am
}

func loadPicture(name string, path string, am *AssetManager) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}
	am.pictureMap[name] = pixel.PictureDataFromImage(img)
}
