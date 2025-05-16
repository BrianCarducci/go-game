package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/BrianCarducci/go-game/game/characters"
	"github.com/BrianCarducci/go-game/game/handlers"

	"log"

)


type Game struct{
	State string
	Player characters.Player
}

func (g *Game) Update() error {
	handlers.HandleKeyInput(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	img, _, err := ebitenutil.NewImageFromFile("doggo.png")
	if err != nil {
		log.Fatal(err)
	}

	screen.DrawImage(img, nil)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}