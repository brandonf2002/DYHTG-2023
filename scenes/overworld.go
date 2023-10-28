package scenes

import (
	"fmt"

	"github.com/brandonf2002/DYHTG-2023/assets"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

func generateTestClicker(am *assets.AssetManager) *Entity {
	name := "testing"

	clikcer := NewEntityAction(func(g *Game, sm *SceneManager) {
		mouseX, mouseY := g.Win.MousePosition().XY()
		fmt.Printf("Mouse X: %v, Mouse Y: %v\n", mouseX, mouseY)
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := Entity{
		Name:   name,
		Sprite: nil,
		Rect:   pixel.R(0, 0, 1024, 1024),
		Actions: []EntityAction{
			clikcer,
		},
	}

	return &player
}

func generateGreenDoor(am *assets.AssetManager) *Entity {
	name := "green_door"

	clikcer := NewEntityAction(func(g *Game, sm *SceneManager) {
		fmt.Println("Clicked!")
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := Entity{
		Name:   name,
		Sprite: nil,
		Rect:   pixel.R(215, 134, 331, 358),
		Actions: []EntityAction{
			clikcer,
		},
	}

	return &player
}
func generateYellowDoor(am *assets.AssetManager) *Entity {
	name := "yellow_door"

	clikcer := NewEntityAction(func(g *Game, sm *SceneManager) {
		fmt.Println("Clicked the yellow door")
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := Entity{
		Name:   name,
		Sprite: nil,
		Rect:   pixel.R(22, 130, 135, 363),
		Actions: []EntityAction{
			clikcer,
		},
	}

	return &player
}
func generateRedDoor(am *assets.AssetManager) *Entity {
	name := "red_door"

	clikcer := NewEntityAction(func(g *Game, sm *SceneManager) {
		fmt.Println("Clicked! The red door")
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := Entity{
		Name:   name,
		Sprite: nil,
		Rect:   pixel.R(696, 131, 819, 363),
		Actions: []EntityAction{
			clikcer,
		},
	}

	return &player
}
func generateBlueDoor(am *assets.AssetManager) *Entity {
	name := "blue_door"

	clikcer := NewEntityAction(func(g *Game, sm *SceneManager) {
		fmt.Println("Clicked the blue door")
	}, MouseClick, pixelgl.MouseButtonLeft)

	player := Entity{
		Name:   name,
		Sprite: nil,
		Rect:   pixel.R(898, 132, 1012, 356),
		Actions: []EntityAction{
			clikcer,
		},
	}

	return &player
}

func GenerateOverworldScene(am *assets.AssetManager) *Scene {
	green_door := generateGreenDoor(am)
	yellow_door := generateYellowDoor(am)
	red_door := generateRedDoor(am)
	blue_door := generateBlueDoor(am)
	testing := generateTestClicker(am)
	return NewScene("overworld", assets.GetPicture("overworld", am), *green_door, *yellow_door, *red_door, *blue_door, *testing)
}
