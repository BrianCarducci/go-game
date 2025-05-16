package handlers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/BrianCarducci/go-game/game"
)

func HandleKeyInput(g game.Game) {
if ebiten.IsKeyPressed(ebiten.KeyW) {
		println("yesss")
	}
}