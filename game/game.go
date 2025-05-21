package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"	

	"github.com/BrianCarducci/go-game/game/characters"
	"github.com/BrianCarducci/go-game/handlers"
	"github.com/BrianCarducci/go-game/state"
)


type Game struct{
	State state.GameState
	Player characters.Player
	Font *text.GoTextFace
}

func (g *Game) Update() error {
	handlers.HandleKeyInput(&g.Player, &g.State)
	// handlers.HandleCollisions(&g.Player, &g.State)
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

	// testing drawing text
	pauseDrawOptions := &text.DrawOptions{}
	pauseDrawOptions.GeoM.Translate(float64(100), float64(100))
	pauseDrawOptions.ColorScale.SetB(100)
	text.Draw(g.Player.Img, "Test Text", g.Font, pauseDrawOptions)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}