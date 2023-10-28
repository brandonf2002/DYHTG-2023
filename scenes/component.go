package scenes

import "github.com/gopxl/pixel"

type ComponentVector struct {
	BoundingBox CBoundingBox
	Sprite      CSprite
}

type CBoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewCBoundingBox(x float64, y float64, width float64, height float64) CBoundingBox {
	return CBoundingBox{X: x, Y: y, Width: width, Height: height}
}

func (bb CBoundingBox) Center() pixel.Vec {
	return pixel.V(bb.X+bb.Width/2, bb.Y+bb.Height/2)
}

func (bb CBoundingBox) Inside(v pixel.Vec) bool {
	return bb.X <= v.X && v.X <= bb.X+bb.Width && bb.Y <= v.Y && v.Y <= bb.Y+bb.Height
}

type CSprite struct {
	Sprite *pixel.Sprite
}

func NewCSprite(sprite *pixel.Sprite) CSprite {
	return CSprite{Sprite: sprite}
}
