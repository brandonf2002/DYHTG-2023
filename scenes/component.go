package scenes

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/text"
)

type ComponentVector struct {
	Text        CText
	Transform   CTransform
	BoundingBox CBoundingBox
	Sprite      CSprite
	Animation   []CAnimation
	LifeSpan    CLifeSpan
}

type CTransform struct {
	Pos        pixel.Vec
	PrevPos    pixel.Vec
	Velocity   pixel.Vec
	Scale      pixel.Vec
	DeltaScale pixel.Vec
	Angle      float64
	DeltaAngle float64
}

func NewCTransform(pos pixel.Vec, velocity pixel.Vec, scale pixel.Vec, deltaScale pixel.Vec, angle float64, deltaAngle float64) CTransform {
	return CTransform{Pos: pos, PrevPos: pos, Velocity: velocity, Scale: scale, DeltaScale: deltaScale, Angle: angle, DeltaAngle: deltaAngle}
}

type CBoundingBox struct {
	size pixel.Vec
}

func NewCBoundingBox(v pixel.Vec) CBoundingBox {
	return CBoundingBox{size: v}
}

func (bb CBoundingBox) Half() pixel.Vec {
	return pixel.V(bb.size.X/2, bb.size.Y/2)
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
	FrameCounter int
	NumOfFrames  int
}

func NewCLifeSpan(numOfFrames int) CLifeSpan {
	return CLifeSpan{FrameCounter: 0, NumOfFrames: numOfFrames}
}

func Inside(v pixel.Vec, cv ComponentVector) bool {
	t := cv.Transform
	bb := cv.BoundingBox
	return t.Pos.X <= v.X && v.X <= t.Pos.X+bb.size.X && t.Pos.Y <= v.Y && v.Y <= t.Pos.Y+bb.size.Y
}

func Add(v pixel.Vec, u pixel.Vec) pixel.Vec {
	return pixel.V(v.X+u.X, v.Y+u.Y)
}

func Sub(v pixel.Vec, u pixel.Vec) pixel.Vec {
	return pixel.V(v.X-u.X, v.Y-u.Y)
}

type CText struct {
	Text        *text.Text
	Str_of_Text string
}

func NewCText(text *text.Text) CText {
	return CText{Text: text, Str_of_Text: ""}
}
