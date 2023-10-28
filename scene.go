package main

import (
	"github.com/gopxl/pixel"
)

type Scene struct {
	name string
	background *pixel.Picture
}

func newScene(name string, background *pixel.Picture) *Scene {
	s := Scene{name: name, background: background}
	return &s
}