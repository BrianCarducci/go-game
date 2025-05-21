package main

import (
	"log"
	"embed"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/BrianCarducci/go-game/game"
	"github.com/BrianCarducci/go-game/setup"
	"github.com/BrianCarducci/go-game/state"
)

//go:embed assets/fonts
var fontFiles embed.FS

var app game.Game


func init() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	Player, defaultFontFace := setup.Setup(fontFiles)

	app = game.Game{ 
		State: state.StatePlaying, 
		Player: Player,
		Font: defaultFontFace,
	}
}

func main() {
	if err := ebiten.RunGame(&app); err != nil {
		log.Fatal(err)
	}
}