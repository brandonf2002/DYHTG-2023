package ecs2

import "github.com/gopxl/pixel"

type ButtonComponent struct {
	Label       string
	Rect        pixel.Rect
	TargetScene string // Name of the scene to transition to when this button is pressed
}

func (b ButtonComponent) Type() ComponentType {
	return NewComponentType()
}
