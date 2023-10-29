package scenes

import "github.com/gopxl/pixel"

type Action struct {
	Name   string
	Coords pixel.Vec
}

func NewAction(name string, coords pixel.Vec) Action {
	return Action{Name: name, Coords: coords}
}
