package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/BrianCarducci/go-game/game"
	"github.com/BrianCarducci/go-game/game/characters"
)

var img *ebiten.Image

func init() {}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	game := game.Game{ 
		State: "game", 
		Player: characters.Player{ Img:ajdlkfjl  Xposition: 0, Yposition: 0 },
	}
	
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}