package handlers

import (
	"github.com/BrianCarducci/go-game/game/characters"
	"github.com/BrianCarducci/go-game/state"
	"github.com/hajimehoshi/ebiten/v2"
)

func HandleKeyInput(p *characters.Player, gameState *state.GameState) {
	// Pause
	// if ebiten.IsKeyPressed(ebiten.KeyEscape) {
	// 	if *gameState == state.StatePlaying {
	// 		*gameState = state.StatePaused
	// 	}
	// 	if *gameState == state.StatePaused {
	// 		*gameState = state.StatePlaying
	// 	}
	// }

	// Player movement
	if (*gameState == state.StatePlaying) {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			p.Yposition -= 1	
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			p.Yposition += 1	
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			p.Xposition -= 1	
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			p.Xposition += 1	
		}
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.Yposition -= 1	
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.Yposition += 1	
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.Xposition -= 1	
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.Xposition += 1	
		}
	}
}