package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/BrianCarducci/go-game/game"
	"github.com/BrianCarducci/go-game/setup"
)

var img *ebiten.Image

func init() {}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	Player := setup.Setup()

	game := game.Game{ 
		State: "game", 
		Player: Player,
	}
	
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}