package scenes

import "github.com/gopxl/pixel"

type ComponentVector struct {
	BoundingBox CBoundingBox
	Sprite      CSprite
}

type CBoundingBox struct {
	X      int
	Y      int
	Width  int
	Height int
}

func NewCBoundingBox(x int, y int, width int, height int) CBoundingBox {
	bb := CBoundingBox{X: x, Y: y, Width: width, Height: height}
	return bb
}

type CSprite struct {
	Sprite *pixel.Sprite
}

func NewCSprite(sprite *pixel.Sprite) CSprite {
	s := CSprite{Sprite: sprite}
	return s
}
