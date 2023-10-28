package scenes

import "github.com/gopxl/pixel"

type ComponentVector struct {
	BoundingBox CBoundingBox
	Sprite      CSprite
	Animation   []CAnimation
	LifeSpan    CLifeSpan
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

type CAnimation struct {
	DeltaScaleX  float64
	DeltaScaleY  float64
	DeltaX       float64
	DeltaY       float64
	DeltaRot     float64
	CurrentFrame int
	NumOfFrames  int
}

func NewCAnimation(deltaScaleX float64, deltaScaleY float64, deltaX float64, deltaY float64, deltaRot float64, currentFrame int, numOfFrames int) CAnimation {
	return CAnimation{DeltaScaleX: deltaScaleX, DeltaScaleY: deltaScaleY, DeltaX: deltaX, DeltaY: deltaY, DeltaRot: deltaRot, CurrentFrame: currentFrame, NumOfFrames: numOfFrames}
}

type CLifeSpan struct {
	NumOfFrames int
}

func NewCLifeSpan(numOfFrames int) CLifeSpan {
	return CLifeSpan{NumOfFrames: numOfFrames}
}
