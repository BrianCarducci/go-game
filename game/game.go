package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/BrianCarducci/go-game/game/characters"
	"github.com/BrianCarducci/go-game/handlers"
)


type Game struct{
	State string
	Player characters.Player
}

func (g *Game) Update() error {
	handlers.HandleKeyInput(&g.Player)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw player
	playerDrawOptions := &ebiten.DrawImageOptions{}
	playerDrawOptions.GeoM.Translate(float64(g.Player.Xposition), float64(g.Player.Yposition))
	screen.DrawImage(
		g.Player.Img, 
		playerDrawOptions,
	)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}