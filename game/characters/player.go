package characters

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	Img ebiten.Image
	Xposition int
	Yposition int
}