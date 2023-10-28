package scenes

type Component interface {
	Has() bool
}

type CBoundingBox struct {
	X int
	Y int
	Width int
	Height int
}

func NewCBoundingBox(x int, y int, width int, height int) *CBoundingBox {
	bb := CBoundingBox{X:x, Y:y, Width:width, Height:height}
	return &bb
}

func (bb *CBoundingBox) Has() bool {
	return true
}

type CChangeScene struct {
	Name string
}

func NewCChangeScene(name string) *CChangeScene {
	cs := CChangeScene{Name:name}
	return &cs
}

func (cs *CChangeScene) Has() bool {
	return true
}
